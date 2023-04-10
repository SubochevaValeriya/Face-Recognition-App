package service

import (
	"github.com/SubochevaValeriya/face-recognition-app/internal/repository"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"os"
)

var DB *sqlx.DB

func Connect() {
	db, err := repository.NewPostgresDB(repository.Config{
		Host:     os.Getenv("host"),
		Port:     viper.GetString("db.port"),
		Username: viper.GetString("db.username"),
		Password: os.Getenv("DB_PASSWORD"),
		DBName:   viper.GetString("db.dbname"),
		SSLMode:  viper.GetString("db.sslmode"),
	})

	if err != nil {
		logrus.Fatalf("failed to inititalize db: %s", err.Error())
	}

	DB = db
}
