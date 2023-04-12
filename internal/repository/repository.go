package repository

import (
	"github.com/SubochevaValeriya/face-recognition-app/internal/models"
	"gorm.io/datatypes"
	"gorm.io/gorm"
	"os"
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

type Repository struct {
	Staff
}

func NewRepository(db *gorm.DB, dbTables DbTables) *Repository {
	return &Repository{NewApiPostgres(db, dbTables)}
}
