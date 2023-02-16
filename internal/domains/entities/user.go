package entities

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Name     string
	LastName string
	Email    string `gorm:"uniqueIndex"`
	IsActive bool
}
