package service

import (
	"os"

	"github.com/SubochevaValeriya/face-recognition-app/internal/repository"
	"github.com/spf13/viper"
	"gorm.io/gorm"
)

func ConnectToDB() (*gorm.DB, error) {
	db, err := repository.NewPostgresDB(repository.Config{
		Host:     os.Getenv("host"),
		Port:     viper.GetString("db.port"),
		Username: viper.GetString("db.username"),
		Password: os.Getenv("DB_PASSWORD"),
		DBName:   viper.GetString("db.dbname"),
		SSLMode:  viper.GetString("db.sslmode"),
	})

	return db, err
}
