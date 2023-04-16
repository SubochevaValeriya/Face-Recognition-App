package repository

import (
	"io"
	"mime/multipart"
	"os"

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
	RecognizeStaff(file os.File) (models.Staff, error)
}

type User interface {
	CreateUser(user *models.User) (*models.User, error)
	GetUser(uid uint) (*models.User, error)
	GetUserByName(username string) (*models.User, error)
}

type Image interface {
	GetImage(id string) (models.Image, error)
	SaveImageToFS(file io.Reader, header *multipart.FileHeader) (string, error)
	CreateImage(image models.Image) (models.Image, error)
}

type Repository struct {
	Staff
	User
	Image
}

func NewRepository(db *gorm.DB, dbTables DbTables) *Repository {
	apiPostgres := NewApiPostgres(db, dbTables)
	return &Repository{
		apiPostgres,
		apiPostgres,
		apiPostgres,
	}
}
