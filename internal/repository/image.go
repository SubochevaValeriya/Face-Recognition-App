package repository

import (
	"io"
	"io/fs"
	"io/ioutil"
	"mime/multipart"
	"os"
	"path/filepath"

	"github.com/SubochevaValeriya/face-recognition-app/internal/models"
	"github.com/google/uuid"
)

const imageTableName = "Image"

func (a ApiPostgres) GetImage(id string) (models.Image, error) {
	var image models.Image
	result := a.db.Table(imageTableName).First(&image, id)
	return image, result.Error
}

func (a ApiPostgres) GetImageByPath(path string) (models.Image, error) {
	var image models.Image
	result := a.db.Table(imageTableName).Where("path = ?", path).First(&image)
	return image, result.Error
}

func (a ApiPostgres) CreateImage(image models.Image) (models.Image, error) {
	result := a.db.Table(imageTableName).Create(&image)
	return image, result.Error
}

func (a ApiPostgres) UpdateImage(image models.Image) (models.Image, error) {
	result := a.db.Table(imageTableName).Save(&image)
	return image, result.Error
}

func (a ApiPostgres) DeleteImage(image models.Image) (models.Image, error) {
	result := a.db.Table(imageTableName).Delete(&image)
	return image, result.Error
}

func (a ApiPostgres) SaveImageToFS(file io.Reader, header *multipart.FileHeader) (string, error) {
	ext := filepath.Ext(header.Filename)
	filename := uuid.New().String() + ext
	path := "images/" + filename
	out, err := os.Create(path)
	if err != nil {
		return "", err
	}
	defer out.Close()
	_, err = io.Copy(out, file)
	if err != nil {
		return "", err
	}
	return filename, nil
}

func (a ApiPostgres) DeleteImageFromFS(filename string) error {
	path := "images/" + filename
	err := os.Remove(path)
	if err != nil {
		return err
	}
	return nil
}

func (a ApiPostgres) GetImageFromFS(filename string) (*os.File, error) {
	file, err := os.Open("images/" + filename)

	if err != nil {
		return &os.File{}, err
	}
	defer file.Close()

	return file, nil
}

func (a ApiPostgres) GetFiles() ([]fs.FileInfo, error) {
	files, err := ioutil.ReadDir("images/")

	if err != nil {
		return []fs.FileInfo{}, err
	}

	return files, err
}
