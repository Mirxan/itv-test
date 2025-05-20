package repositories

import (
	"github.com/Mirxan/itv-test/models"

	"gorm.io/gorm"
)

type MovieRepository struct {
	db *gorm.DB
}

func NewMovieRepository(db *gorm.DB) *MovieRepository {
	return &MovieRepository{db: db}
}

func (r *MovieRepository) Create(movie *models.Movie) error {
	return r.WithTransaction(func(tx *gorm.DB) error {
		if err := tx.Create(movie).Error; err != nil {
			return err
		}
		return nil
	})
}

func (r *MovieRepository) FindAll() ([]models.Movie, error) {
	var movies []models.Movie
	err := r.db.Find(&movies).Error
	return movies, err
}

func (r *MovieRepository) FindByID(id uint) (*models.Movie, error) {
	var movie models.Movie
	err := r.db.First(&movie, id).Error
	return &movie, err
}

func (r *MovieRepository) Update(movie *models.Movie) error {
	return r.WithTransaction(func(tx *gorm.DB) error {
		if err := r.db.Save(movie).Error; err != nil {
			return err
		}
		return nil
	})
}

func (r *MovieRepository) Delete(id uint) error {
	return r.db.Delete(&models.Movie{}, id).Error
}

func (r *MovieRepository) WithTransaction(txFunc func(*gorm.DB) error) error {
	tx := r.db.Begin()
	if tx.Error != nil {
		return tx.Error
	}

	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
			panic(r)
		}
	}()

	if err := txFunc(tx); err != nil {
		tx.Rollback()
		return err
	}

	return tx.Commit().Error
}
