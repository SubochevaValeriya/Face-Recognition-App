package service

import (
	"github.com/SubochevaValeriya/face-recognition-app/internal/repository"
)

type Staff interface {
	//
}

type Service struct {
	Staff
}

func NewService(repos *repository.Repository) *Service {
	return &Service{newApiService(repos.Staff)}
}
