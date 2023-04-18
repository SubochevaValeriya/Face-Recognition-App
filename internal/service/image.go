package service

import (
	"io"
	"mime/multipart"
	"os"
	"reflect"

	"github.com/SubochevaValeriya/face-recognition-app/internal/models"
	"github.com/SubochevaValeriya/face-recognition-app/internal/repository"
	face "github.com/modeckrus/go-face"
)

type ImageApiService struct {
	repo repository.Image
}

func newImageApiService(repo repository.Image) *ImageApiService {
	return &ImageApiService{repo: repo}
}

func (s *ImageApiService) GetImage(id string) (models.Image, error) {
	return s.repo.GetImage(id)
}

func (s *ImageApiService) GetImageAsFile(id string) (string, io.Reader, error) {
	image, err := s.repo.GetImage(id)

	if err != nil {
		return "", &os.File{}, err
	}

	imgFile, err := s.repo.GetImageFromFS(image.Path)

	if err != nil {
		return "", &os.File{}, err
	}

	return image.Path, imgFile, err
}

func (s *ImageApiService) SaveImage(file io.Reader, header *multipart.FileHeader) (models.Image, error) {
	path, err := s.repo.SaveImageToFS(file, header)

	if err != nil {
		return models.Image{}, err
	}

	image := models.Image{Path: path}

	image, err = s.repo.CreateImage(image)

	return image, err
}

func (s *ImageApiService) UploadImageWithFace(file io.Reader, header *multipart.FileHeader) (models.Image, error) {
	path, err := s.repo.SaveImageToFS(file, header)

	if err != nil {
		return models.Image{}, err
	}

	faces, err := s.RecognizeFaceOnImage(path)

	if err != nil || len(faces) == 0 {
		return models.Image{}, err
	}

	image := models.Image{Path: path}

	image, err = s.repo.CreateImage(image)

	return image, err
}

func (s *ImageApiService) RecognizeFaceOnImage(path string) ([]face.Face, error) {
	rec, err := face.NewRecognizer("models")

	if err != nil {
		return []face.Face{}, err
	}
	defer rec.Close()

	imgFile, err := s.repo.GetImageFromFS(path)

	if err != nil {
		return []face.Face{}, err
	}

	data, err := io.ReadAll(imgFile)

	if err != nil {
		return []face.Face{}, err
	}

	faces, err := rec.Recognize(data)

	if err != nil || len(faces) == 0 {
		return []face.Face{}, err
	}

	return faces, err
}

func (s *ImageApiService) RecognizeImage(file io.Reader, header *multipart.FileHeader) (models.Image, error) {
	temporaryImage, err := s.repo.SaveImageToFS(file, header)

	if err != nil {
		return models.Image{}, err
	}
	defer s.repo.DeleteImageFromFS(temporaryImage)

	facesToCompare, err := s.RecognizeFaceOnImage(temporaryImage)

	if err != nil || len(facesToCompare) == 0 {
		return models.Image{}, err
	}

	files, err := s.repo.GetFiles()

	if err != nil {
		return models.Image{}, err
	}

	var result models.Image
	for _, file := range files {
		if !file.IsDir() {
			faces, err := s.RecognizeFaceOnImage(file.Name())

			if err != nil || len(faces) == 0 {
				return models.Image{}, err
			}

			if !testPartialEqual(facesToCompare, faces) {
				continue
			}

			result, err = s.repo.GetImageByPath(file.Name())
			if err != nil {
				return models.Image{}, err
			}
			return result, nil
		}
	}

	return models.Image{}, nil
}

func testPartialEqual(a []face.Face, b []face.Face) bool {
	if len(a) == len(b) {
		return true
	}
	for i := range a {
		if reflect.DeepEqual(a[i], b[i]) {
			return true
		}
	}
	return true
}
