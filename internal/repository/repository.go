package repository

import (
	"context"
	"io"
	"io/fs"
	"mime/multipart"
	"os"
	"time"

	"github.com/SubochevaValeriya/face-recognition-app/internal/models"
	"gorm.io/datatypes"
	"gorm.io/gorm"
)

type Staff interface {
	AddStaff(staff models.Staff) (models.Staff, error)
	UpdateStaff(updatedStaff models.Staff) (models.Staff, error)
	DeleteStaff(id int) error
	GetStaff(id int) (models.Staff, error)
	GetAllStaff() ([]models.Staff, error)
	FindStaff(meta datatypes.JSONMap) ([]models.Staff, error)
	RecognizeStaff(imageId int) (models.Staff, error)
}

type User interface {
	CreateUser(user *models.User) (*models.User, error)
	GetUser(uid uint) (*models.User, error)
	GetUserByName(username string) (*models.User, error)
}

type Image interface {
	GetImage(id string) (models.Image, error)
	GetImageByPath(path string) (models.Image, error)
	CreateImage(image models.Image) (models.Image, error)
	DeleteImageFromFS(filename string) error
	SaveImageToFS(file io.Reader, header *multipart.FileHeader) (string, error)
	GetImageFromFS(path string) (*os.File, error)
	GetFiles() ([]fs.FileInfo, error)
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

type Repository struct {
	Staff
	User
	Image
	TimeRecordDb
}

func NewRepository(db *gorm.DB, dbTables DbTables) *Repository {
	apiPostgres := NewApiPostgres(db, dbTables)
	return &Repository{
		apiPostgres,
		apiPostgres,
		apiPostgres,
	}
}
