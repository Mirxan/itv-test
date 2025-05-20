package utils

import (
	"fmt"

	"github.com/Mirxan/itv-test/config"
	"github.com/Mirxan/itv-test/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func NewDBConnection(config *config.Config) (*gorm.DB, error) {
	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=UTC",
		config.DBHost,
		config.DBUser,
		config.DBPassword,
		config.DBName,
		config.DBPort,
	)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	// AutoMigrate models
	err = db.AutoMigrate(&models.User{}, &models.Movie{})
	if err != nil {
		return nil, err
	}

	return db, nil
}
