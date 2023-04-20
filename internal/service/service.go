package service

import (
	"context"
	"io"
	"mime/multipart"
	"time"

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
	RecognizeStaff(imageId int) (models.Staff, error)
}

type User interface {
	GetUserByID(uid uint) (models.User, error)
	LoginCheck(username string, password string) (string, error)
	SaveUser(user *models.User) (*models.User, error)
}

type Image interface {
	GetImage(id string) (models.Image, error)
	GetImageAsFile(id string) (string, io.Reader, error)
	SaveImage(file io.Reader, header *multipart.FileHeader) (models.Image, error)
	RecognizeImage(file io.Reader, header *multipart.FileHeader) (models.Image, error)
	UploadImageWithFace(file io.Reader, header *multipart.FileHeader) (models.Image, error)
}

type TimeRecordDb interface {
	Add(timeRecord models.AddTimeRecord) (models.TimeRecord, error)
	Update(timeRecord models.UpdateTimeRecord) (models.TimeRecord, error)
	Delete(id int) error
	Get(id int) (models.TimeRecord, error)
	All() ([]models.TimeRecord, error)
	ByEmployeeId(id int) ([]models.TimeRecord, error)
	ByDate(start time.Time, end time.Time, employeeId int) ([]models.TimeRecord, error)
	LastByEmployeeId(id int) (models.TimeRecord, error)
	Stream(ctx context.Context) (chan models.StreamModel[models.TimeRecord], error)
}
type Service struct {
	Staff
	User
	Image
	TimeRecordDb
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		newStaffApiService(repos.Staff),
		newUserApiService(repos.User),
		newImageApiService(repos.Image),
	}
}
