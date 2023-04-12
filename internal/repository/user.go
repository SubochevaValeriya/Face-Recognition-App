package repository

import (
	"github.com/SubochevaValeriya/face-recognition-app/internal/models"
)

func (a ApiPostgres) CreateUser(user *models.User) (*models.User, error) {
	err := a.db.Create(&user).Error
	if err != nil {
		return &models.User{}, err
	}
	return user, nil
}

func (a ApiPostgres) GetUser(uid uint) (*models.User, error) {
	var user models.User
	err := a.db.First(&user, uid).Error
	if err != nil {
		return &models.User{}, err
	}
	return &user, nil
}

func (a ApiPostgres) GetUserByName(username string) (*models.User, error) {
	user := models.User{}
	err := a.db.Model(models.User{}).Where("username = ?", username).Take(&user).Error
	if err != nil {
		return &models.User{}, err
	}
	return &user, nil
}
