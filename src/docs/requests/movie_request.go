package requests

type CreateMovieRequest struct {
	Title       string  `json:"title" binding:"required" example:"Title of the Movie"`
	Director    string  `json:"director" binding:"required" example:"Director Name"`
	ReleaseYear int     `json:"release_year" binding:"required" example:"2010"`
	Genre       string  `json:"genre" binding:"required" example:"Genre"`
	Plot        string  `json:"plot" binding:"required" example:"Plot description goes here."`
	Rating      float64 `json:"rating" example:"8.8"`
	Duration    int     `json:"duration" binding:"required" example:"148"`
	Language    string  `json:"language" binding:"required" example:"English"`
	Country     string  `json:"country" binding:"required" example:"USA"`
}

type UpdateMovieRequest struct {
	Title       string  `json:"title" example:"Title of the Movie"`
	Director    string  `json:"director" example:"Director Name"`
	ReleaseYear int     `json:"release_year" example:"2010"`
	Genre       string  `json:"genre" example:"Genre"`
	Plot        string  `json:"plot" example:"Plot description goes here."`
	Rating      float64 `json:"rating" example:"8.8"`
	Duration    int     `json:"duration" example:"148"`
	Language    string  `json:"language" example:"English"`
	Country     string  `json:"country" example:"USA"`
}
