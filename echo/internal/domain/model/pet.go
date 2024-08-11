package model

import (
	"database/sql"

	"gorm.io/gorm"
)

type Pet struct {
	gorm.Model
	Name        string
	Tag         sql.NullString
	DateOfBirth sql.NullTime
}
