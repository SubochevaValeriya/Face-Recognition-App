package service

import (
	"github.com/SubochevaValeriya/face-recognition-app/internal/repository"
)

type ApiService struct {
	repo repository.Staff
}

func newApiService(repo repository.Staff) *ApiService {
	return &ApiService{repo: repo}
}
