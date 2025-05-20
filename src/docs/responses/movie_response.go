package responses

import (
	"time"

	"github.com/Mirxan/itv-test/models"
)

type MovieResponse struct {
	ID          uint    `json:"id" example:"1"`
	Title       string  `json:"title" example:"Title of the Movie"`
	Director    string  `json:"director" example:"Director Name"`
	ReleaseYear int     `json:"release_year" example:"2010"`
	Genre       string  `json:"genre" example:"Sci-Fi"`
	Plot        string  `json:"plot" example:"Plot description goes here."`
	Rating      float64 `json:"rating" example:"8.8"`
	Duration    int     `json:"duration" example:"148"`
	Language    string  `json:"language" example:"English"`
	Country     string  `json:"country" example:"USA"`
	CreatedAt   string  `json:"created_at" example:"2023-01-01T00:00:00Z"`
	UpdatedAt   string  `json:"updated_at" example:"2023-01-01T00:00:00Z"`
}

func NewMovieResponse(movie *models.Movie) *MovieResponse {
	return &MovieResponse{
		ID:          movie.ID,
		Title:       movie.Title,
		Director:    movie.Director,
		ReleaseYear: movie.ReleaseYear,
		Genre:       movie.Genre,
		Plot:        movie.Plot,
		Rating:      movie.Rating,
		Duration:    movie.Duration,
		Language:    movie.Language,
		Country:     movie.Country,
		CreatedAt:   movie.CreatedAt.Format(time.RFC3339),
		UpdatedAt:   movie.UpdatedAt.Format(time.RFC3339),
	}
}
