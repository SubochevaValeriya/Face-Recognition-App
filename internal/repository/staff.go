package repository

import (
	"fmt"
	"github.com/SubochevaValeriya/face-recognition-app/internal/models"
	"gorm.io/datatypes"
)

const staffTableName = "Staff"

func (a ApiPostgres) AddStaff(staff models.Staff) (models.Staff, error) {
	result := a.db.Table(staffTableName).Create(&staff)
	return staff, result.Error
}

func (a ApiPostgres) UpdateStaff(updatedStaff models.Staff) (models.Staff, error) {
	var staff models.Staff

	if result := a.db.Table(staffTableName).First(&staff, updatedStaff.ID); result.Error != nil {
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

	if result := a.db.Table(staffTableName).First(&staff, id); result.Error != nil {
		fmt.Println(result.Error)
	}

	result := a.db.Table(staffTableName).Delete(&staff)
	return result.Error
}

func (a ApiPostgres) GetStaff(id int) (models.Staff, error) {
	var staff models.Staff

	result := a.db.Table(staffTableName).First(&staff, id)
	return staff, result.Error
}

func (a ApiPostgres) GetAllStaff() ([]models.Staff, error) {
	var staff []models.Staff

	result := a.db.Table(staffTableName).Find(&staff)
	return staff, result.Error
}

func (a ApiPostgres) FindStaff(meta datatypes.JSONMap) ([]models.Staff, error) {
	var staff []models.Staff

	result := a.db.Table(staffTableName).Where("meta = ?", meta).Find(&staff)
	return staff, result.Error
}

func (a ApiPostgres) RecognizeStaff(imageId int) (models.Staff, error) {
	var staff models.Staff

	result := a.db.Table(staffTableName).Where("photo_id = ?", imageId).Find(&staff)
	return staff, result.Error
}
