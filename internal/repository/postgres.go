package repository

import (
	"fmt"
	"github.com/SubochevaValeriya/face-recognition-app/internal/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Config struct {
	Host     string
	Port     string
	Username string
	Password string
	DBName   string
	SSLMode  string
}

func NewPostgresDB(cfg Config) (*gorm.DB, error) {
	//dbURL := fmt.Sprintf("postgres://%s:pass@%s:%s/", cfg.Username, cfg.Host, cfg.Port)
	config := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=%s TimeZone=Europe/Moscow", cfg.Host, cfg.Username, cfg.Password, cfg.DBName, cfg.Port, cfg.SSLMode)
	//	db, err := gorm.Open(postgres.Open(fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=%s", cfg.Host, cfg.Port, cfg.Username, cfg.DBName, cfg.Password, cfg.SSLMode)), &gorm.Config{})
	db, err := gorm.Open(postgres.Open(config), &gorm.Config{})
	if err != nil {
		return nil, fmt.Errorf("connect error: %s", err)
	}
	db.Table("Staff").AutoMigrate(&models.Staff{})
	return db, nil
}

type ApiPostgres struct {
	db       *gorm.DB
	dbTables DbTables
}

type DbTables struct {
	User        string
	Staff       string
	Images      string
	TimeRecords string
}

func NewApiPostgres(db *gorm.DB, dbTables DbTables) *ApiPostgres {
	return &ApiPostgres{db: db,
		dbTables: dbTables}
}
