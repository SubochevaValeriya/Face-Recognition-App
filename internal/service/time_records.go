package service

import (
	"github.com/SubochevaValeriya/face-recognition-app/internal/repository"
)

type TimeApiService struct {
	repo repository.Staff
}

func newTimeApiService(repo repository.Staff) *TimeApiService {
	return &TimeApiService{repo: repo}
}
