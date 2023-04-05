package service

import (
	"github.com/SubochevaValeriya/face-recognition-app/internal/repository"
)

type ImageApiService struct {
	repo repository.Staff
}

func newImageApiService(repo repository.Staff) *ImageApiService {
	return &ImageApiService{repo: repo}
}
