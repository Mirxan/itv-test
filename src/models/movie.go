package models

import (
	"time"

	"gorm.io/gorm"
)

type Movie struct {
	gorm.Model
	Title       string    `json:"title" gorm:"size:255;not null"`
	Director    string    `json:"director" gorm:"size:255;not null"`
	ReleaseYear int       `json:"release_year" gorm:"not null"`
	Genre       string    `json:"genre" gorm:"size:100;not null"`
	Plot        string    `json:"plot" gorm:"type:text;not null"`
	Rating      float64   `json:"rating" gorm:"type:decimal(3,1);default:0.0"`
	Duration    int       `json:"duration" gorm:"not null"` // in minutes
	Language    string    `json:"language" gorm:"size:100;not null"`
	Country     string    `json:"country" gorm:"size:100;not null"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

func (Movie) TableName() string {
	return "movies"
}
