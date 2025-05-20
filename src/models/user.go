package models

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Username string `json:"username" gorm:"unique;not null"`
	Email    string `json:"email" gorm:"unique;not null"`
	Password string `json:"-" gorm:"not null"`
	IsAdmin  bool   `json:"is_admin" gorm:"default:false"`
}

func (User) TableName() string {
	return "users"
}
