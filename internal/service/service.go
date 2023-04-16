package service

import (
	"io"
	"mime/multipart"
	"os"

	"github.com/SubochevaValeriya/face-recognition-app/internal/models"
	"github.com/SubochevaValeriya/face-recognition-app/internal/repository"
	"gorm.io/datatypes"
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

type User interface {
	GetUserByID(uid uint) (models.User, error)
	LoginCheck(username string, password string) (string, error)
	SaveUser(user *models.User) (*models.User, error)
}

type Image interface {
	GetImage(id string) (models.Image, error)
	SaveImage(file io.Reader, header *multipart.FileHeader) (models.Image, error)
}

type Service struct {
	Staff
	User
	Image
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		newStaffApiService(repos.Staff),
		newUserApiService(repos.User),
		newImageApiService(repos.Image),
	}
}
