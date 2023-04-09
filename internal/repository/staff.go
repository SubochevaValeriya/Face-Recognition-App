package repository

import (
	"fmt"
	"github.com/SubochevaValeriya/face-recognition-app/internal/models"
	"gorm.io/datatypes"
	"os"
)

const tableName = "Staff"

func (a ApiPostgres) AddStaff(staff models.Staff) (models.Staff, error) {
	result := a.db.Table(tableName).Create(&staff)
	return staff, result.Error
}

func (a ApiPostgres) UpdateStaff(updatedStaff models.Staff) (models.Staff, error) {
	var staff models.Staff

	if result := a.db.Table(tableName).First(&staff, updatedStaff.ID); result.Error != nil {
		return staff, result.Error
	}

	staff.Name = updatedStaff.Name
	staff.Meta = updatedStaff.Meta
	staff.PhotoId = updatedStaff.PhotoId

	result := a.db.Save(&staff)

	return staff, result.Error
}

func (a ApiPostgres) DeleteStaff(id int) error {
	var staff models.Staff

	if result := a.db.Table(tableName).First(&staff, id); result.Error != nil {
		fmt.Println(result.Error)
	}

	result := a.db.Table(tableName).Delete(&staff)
	return result.Error
}

func (a ApiPostgres) GetStaff(id int) (models.Staff, error) {
	var staff models.Staff

	result := a.db.Table(tableName).First(&staff, id)
	return staff, result.Error
}

func (a ApiPostgres) GetAllStaff() ([]models.Staff, error) {
	var staff []models.Staff

	result := a.db.Table(tableName).Find(&staff)
	return staff, result.Error
}

func (a ApiPostgres) FindStaff(meta datatypes.JSONMap) ([]models.Staff, error) {
	var staff []models.Staff

	result := a.db.Table(tableName).Where("meta = ?", meta).Find(&staff)
	return staff, result.Error
}

func (a ApiPostgres) RecognizeStaff(file os.File) (models.Staff, error) {
	//TODO implement me
	panic("implement me")
}
