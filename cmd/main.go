package main

import (
	"context"
	"os"
	"os/signal"
	"syscall"

	faceRecognition "github.com/SubochevaValeriya/face-recognition-app"
	"github.com/SubochevaValeriya/face-recognition-app/internal/handler"
	"github.com/SubochevaValeriya/face-recognition-app/internal/repository"
	"github.com/SubochevaValeriya/face-recognition-app/internal/service"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

// @title Face Recognition API
// @version 1.0
// @description API Server for Face Recognition Application

// @host localhost:8000
// @BasePath /

func main() {
	logrus.SetFormatter(new(logrus.JSONFormatter))

	if err := initConfig(); err != nil {
		logrus.Fatalf("error initializing congigs: %s", err.Error())
	}

	if err := godotenv.Load(); err != nil {
		logrus.Fatalf("error loading env variables: %s", err.Error())
	}

	//connect w/t docker-compose:
	//sudo docker run --name=faceRecognition -e POSTGRES_PASSWORD='qwerty' -p 5432:5432 -d --rm postgres
	// migrate -path ./schema -database 'postgres://postgres:qwerty@localhost:5432/postgres?sslmode=disable' up

	db, err := service.ConnectToDB()
	if err != nil {
		logrus.Fatalf("failed to inititalize db: %s", err.Error())
	}

	dbTables := repository.DbTables{
		User:        viper.GetString("dbTables.user"),
		Staff:       viper.GetString("dbTables.staff"),
		Images:      viper.GetString("dbTables.images"),
		TimeRecords: viper.GetString("dbTables.time_records"),
	}

	// dependency injection
	repos := repository.NewRepository(db, dbTables)
	services := service.NewService(repos)
	handlers := handler.NewHandler(services)
	srv := new(faceRecognition.Server)

	go func() {
		if err := srv.Run(viper.GetString("port"), handlers.InitRoutes()); err != nil {
			logrus.Fatalf("error occured while running http server: %s", err.Error())
		}
	}()

	logrus.Print("App Started")

	// graceful shutdown
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)
	<-quit

	logrus.Print("App Shutting Down")

	if err := srv.Shutdown(context.Background()); err != nil {
		logrus.Errorf("error occured on server shutting down: %s", err.Error())
	}

	sqlDB, err := db.DB()
	if err := sqlDB.Close(); err != nil {
		logrus.Errorf("error occured on db connection close: %s", err.Error())
	}
}

func initConfig() error {
	viper.AddConfigPath("configs")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}
