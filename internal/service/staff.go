package service

import (
	"github.com/SubochevaValeriya/face-recognition-app/internal/repository"
)

type StaffApiService struct {
	repo repository.Staff
}

func newStaffApiService(repo repository.Staff) *StaffApiService {
	return &StaffApiService{repo: repo}
}
