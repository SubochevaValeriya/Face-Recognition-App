package repository

import (
	"github.com/jmoiron/sqlx"
)

type Staff interface {
}

type Repository struct {
	Staff
}

func NewRepository(db *sqlx.DB, dbTables DbTables) *Repository {
	return &Repository{NewApiPostgres(db, dbTables)}
}
