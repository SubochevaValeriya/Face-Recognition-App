package service

import (
	"github.com/SubochevaValeriya/face-recognition-app/internal/models"
	"github.com/SubochevaValeriya/face-recognition-app/internal/repository"
	"gorm.io/datatypes"
	"os"
)

type Staff interface {
	AddStaff(staff models.Staff) (models.Staff, error)
	UpdateStaff(staff models.Staff) (models.Staff, error)
	DeleteStaff(id int) error
	GetStaff(id int) (models.Staff, error)
	GetAllStaff() ([]models.Staff, error)
	FindStaff(meta datatypes.JSONMap) ([]models.Staff, error)
	RecognizeStaff(file os.File) (models.Staff, error)
}

type Service struct {
	Staff
}

func NewService(repos *repository.Repository) *Service {
	return &Service{newStaffApiService(repos.Staff)}
}
