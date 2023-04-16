package service

import (
	"io"
	"mime/multipart"

	"github.com/SubochevaValeriya/face-recognition-app/internal/models"
	"github.com/SubochevaValeriya/face-recognition-app/internal/repository"
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

func (s *ImageApiService) SaveImage(file io.Reader, header *multipart.FileHeader) (models.Image, error) {
	path, err := s.repo.SaveImageToFS(file, header)

	if err != nil {
		return models.Image{}, err
	}

	image := models.Image{Path: path}

	image, err = s.repo.CreateImage(image)

	return image, err
}
