package service

import (
	"github.com/SubochevaValeriya/face-recognition-app/internal/models"
	"github.com/SubochevaValeriya/face-recognition-app/internal/repository"
	"gorm.io/datatypes"
	"path/filepath"
)

type StaffApiService struct {
	repo repository.Staff
}

func newStaffApiService(repo repository.Staff) *StaffApiService {
	return &StaffApiService{repo: repo}
}

func (s *StaffApiService) AddStaff(staff models.Staff) (models.Staff, error) {
	return s.repo.AddStaff(staff)
}

func (s *StaffApiService) UpdateStaff(staff models.Staff) (models.Staff, error) {
	return s.repo.UpdateStaff(staff)
}

func (s *StaffApiService) DeleteStaff(id int) error {
	return s.repo.DeleteStaff(id)
}

func (s *StaffApiService) GetStaff(id int) (models.Staff, error) {
	return s.repo.GetStaff(id)
}

func (s *StaffApiService) GetAllStaff() ([]models.Staff, error) {
	return s.repo.GetAllStaff()
}
func (s *StaffApiService) FindStaff(meta datatypes.JSONMap) ([]models.Staff, error) {
	return s.repo.FindStaff(meta)
}

const dataDir = "testdata"

var (
	modelsDir = filepath.Join(dataDir, "models")
	imagesDir = filepath.Join(dataDir, "images")
)

func (s *StaffApiService) RecognizeStaff(imageId int) (models.Staff, error) {
	return s.repo.RecognizeStaff(imageId)
}
