package models

import "gorm.io/datatypes"

type Staff struct {
	ID      int               `json:"id" db:"id" gorm:"primaryKey"`
	Name    string            `json:"name" db:"name"`
	PhotoId int               `json:"photo_id" db:"photo_id"`
	Meta    datatypes.JSONMap `json:"meta" db:"meta"`
}
