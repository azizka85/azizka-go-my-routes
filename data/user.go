package data

import (
	"gorm.io/gorm"
)

type User struct {
	Email    string
	Password string
	FullName string
	Photo    string
	gorm.Model
}
