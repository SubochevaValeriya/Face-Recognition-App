package repository

import (
	"github.com/SubochevaValeriya/face-recognition-app/internal/models"
)

const userTableName = "User"

func (a ApiPostgres) CreateUser(user *models.User) (*models.User, error) {
	err := a.db.Table(userTableName).Create(&user).Error
	if err != nil {
		return &models.User{}, err
	}
	return user, nil
}

func (a ApiPostgres) GetUser(uid uint) (*models.User, error) {
	var user models.User
	err := a.db.Table(userTableName).First(&user, uid).Error
	if err != nil {
		return &models.User{}, err
	}
	return &user, nil
}

func (a ApiPostgres) GetUserByName(username string) (*models.User, error) {
	user := models.User{}
	err := a.db.Table(userTableName).Model(models.User{}).Where("username = ?", username).Take(&user).Error
	if err != nil {
		return &models.User{}, err
	}
	return &user, nil
}
