package services

import (
	"github.com/Mirxan/itv-test/exceptions"
	"github.com/Mirxan/itv-test/models"
	"github.com/Mirxan/itv-test/repositories"
)

type MovieService struct {
	repo *repositories.MovieRepository
}

func NewMovieService(repo *repositories.MovieRepository) *MovieService {
	return &MovieService{repo: repo}
}

func (s *MovieService) CreateMovie(movie *models.Movie) error {
	if movie.Title == "" || movie.Director == "" || movie.Plot == "" {
		return exceptions.NewBadRequestError("Title, director, and plot are required")
	}

	if movie.ReleaseYear < 1888 { // First movie ever made was in 1888
		return exceptions.NewBadRequestError("Invalid release year")
	}

	return s.repo.Create(movie)
}

func (s *MovieService) GetAllMovies() ([]models.Movie, error) {
	return s.repo.FindAll()
}

func (s *MovieService) GetMovieByID(id uint) (*models.Movie, error) {
	movie, err := s.repo.FindByID(id)
	if err != nil {
		return nil, exceptions.NewNotFoundError("Movie not found")
	}
	return movie, nil
}

func (s *MovieService) UpdateMovie(id uint, updatedMovie *models.Movie) (*models.Movie, error) {
	existingMovie, err := s.repo.FindByID(id)
	if err != nil {
		return nil, exceptions.NewNotFoundError("Movie not found")
	}

	// Update fields
	existingMovie.Title = updatedMovie.Title
	existingMovie.Director = updatedMovie.Director
	existingMovie.ReleaseYear = updatedMovie.ReleaseYear
	existingMovie.Genre = updatedMovie.Genre
	existingMovie.Plot = updatedMovie.Plot
	existingMovie.Rating = updatedMovie.Rating
	existingMovie.Duration = updatedMovie.Duration
	existingMovie.Language = updatedMovie.Language
	existingMovie.Country = updatedMovie.Country

	err = s.repo.Update(existingMovie)
	if err != nil {
		return nil, err
	}

	return existingMovie, nil
}

func (s *MovieService) DeleteMovie(id uint) error {
	_, err := s.repo.FindByID(id)
	if err != nil {
		return exceptions.NewNotFoundError("Movie not found")
	}

	return s.repo.Delete(id)
}
