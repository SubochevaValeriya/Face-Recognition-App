package models

type Image struct {
	ID   int    `json:"id" db:"id" gorm:"primaryKey"`
	Path string `json:"path" db:"path"`
}
