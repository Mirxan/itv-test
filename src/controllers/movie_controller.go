package controllers

import (
	"net/http"
	"strconv"

	"github.com/Mirxan/itv-test/docs/requests"
	"github.com/Mirxan/itv-test/docs/responses"
	"github.com/Mirxan/itv-test/exceptions"
	"github.com/Mirxan/itv-test/models"
	"github.com/Mirxan/itv-test/services"
	"github.com/go-playground/validator/v10"

	"github.com/gin-gonic/gin"
)

type MovieController struct {
	service *services.MovieService
}

func NewMovieController(service *services.MovieService) *MovieController {
	return &MovieController{service: service}
}

// GetMovies godoc
// @Summary Get all movies
// @Description Get a list of all movies
// @Tags movies
// @Accept  json
// @Produce  json
// @Security BearerAuth
// @Success 200 {array} responses.MovieResponse
// @Failure 401 {object} responses.ErrorResponse
// @Failure 500 {object} responses.ErrorResponse
// @Router /movies [get]
func (c *MovieController) GetMovies(ctx *gin.Context) {
	movies, err := c.service.GetAllMovies()
	if err != nil {
		ctx.JSON(err.(*exceptions.AppError).Code, responses.NewErrorResponse(err.Error()))
		return
	}

	var response []responses.MovieResponse
	for _, movie := range movies {
		response = append(response, *responses.NewMovieResponse(&movie))
	}

	ctx.JSON(http.StatusOK, gin.H{"data": response})
}

// GetMovie godoc
// @Summary Get a movie by ID
// @Description Get a movie by its ID
// @Tags movies
// @Accept  json
// @Produce  json
// @Param id path int true "Movie ID"
// @Security BearerAuth
// @Success 200 {object} responses.MovieResponse
// @Failure 400 {object} responses.ErrorResponse
// @Failure 401 {object} responses.ErrorResponse
// @Failure 404 {object} responses.ErrorResponse
// @Failure 500 {object} responses.ErrorResponse
// @Router /movies/{id} [get]
func (c *MovieController) GetMovie(ctx *gin.Context) {
	id, err := strconv.ParseUint(ctx.Param("id"), 10, 64)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, responses.NewErrorResponse("Invalid movie ID"))
		return
	}

	movie, err := c.service.GetMovieByID(uint(id))
	if err != nil {
		ctx.JSON(err.(*exceptions.AppError).Code, responses.NewErrorResponse(err.Error()))
		return
	}

	ctx.JSON(http.StatusOK, responses.NewMovieResponse(movie))
}

// CreateMovie godoc
// @Summary Create a new movie
// @Description Create a new movie with the input payload
// @Tags movies
// @Accept  json
// @Produce  json
// @Param movie body requests.CreateMovieRequest true "Create movie"
// @Security BearerAuth
// @Success 201 {object} responses.MovieResponse
// @Failure 400 {object} responses.ErrorResponse
// @Failure 401 {object} responses.ErrorResponse
// @Failure 500 {object} responses.ErrorResponse
// @Router /movies [post]
func (c *MovieController) CreateMovie(ctx *gin.Context) {
	var req requests.CreateMovieRequest

	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, exceptions.NewValidationErrors(err))
		return
	}

	movie := models.Movie{
		Title:       req.Title,
		Director:    req.Director,
		ReleaseYear: req.ReleaseYear,
		Genre:       req.Genre,
		Plot:        req.Plot,
		Rating:      req.Rating,
		Duration:    req.Duration,
		Language:    req.Language,
		Country:     req.Country,
	}

	if err := c.service.CreateMovie(&movie); err != nil {
		ctx.JSON(err.(*exceptions.AppError).Code, responses.NewErrorResponse(err.Error()))
		return
	}

	ctx.JSON(http.StatusCreated, responses.NewMovieResponse(&movie))
}

// UpdateMovie godoc
// @Summary Update a movie
// @Description Update a movie by its ID
// @Tags movies
// @Accept  json
// @Produce  json
// @Param id path int true "Movie ID"
// @Param movie body requests.UpdateMovieRequest true "Update movie"
// @Security BearerAuth
// @Success 200 {object} responses.MovieResponse
// @Failure 400 {object} responses.ErrorResponse
// @Failure 401 {object} responses.ErrorResponse
// @Failure 404 {object} responses.ErrorResponse
// @Failure 500 {object} responses.ErrorResponse
// @Router /movies/{id} [put]
func (c *MovieController) UpdateMovie(ctx *gin.Context) {
	id, err := strconv.ParseUint(ctx.Param("id"), 10, 64)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, responses.NewErrorResponse("Invalid movie ID"))
		return
	}

	var req requests.UpdateMovieRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, exceptions.NewValidationErrors(err.(validator.ValidationErrors)))
		return
	}

	updatedMovie := models.Movie{
		Title:       req.Title,
		Director:    req.Director,
		ReleaseYear: req.ReleaseYear,
		Genre:       req.Genre,
		Plot:        req.Plot,
		Rating:      req.Rating,
		Duration:    req.Duration,
		Language:    req.Language,
		Country:     req.Country,
	}

	movie, err := c.service.UpdateMovie(uint(id), &updatedMovie)
	if err != nil {
		ctx.JSON(err.(*exceptions.AppError).Code, responses.NewErrorResponse(err.Error()))
		return
	}

	ctx.JSON(http.StatusOK, responses.NewMovieResponse(movie))
}

// DeleteMovie godoc
// @Summary Delete a movie
// @Description Delete a movie by its ID
// @Tags movies
// @Accept  json
// @Produce  json
// @Param id path int true "Movie ID"
// @Security BearerAuth
// @Success 200 {object} responses.SuccessResponse
// @Failure 400 {object} responses.ErrorResponse
// @Failure 401 {object} responses.ErrorResponse
// @Failure 404 {object} responses.ErrorResponse
// @Failure 500 {object} responses.ErrorResponse
// @Router /movies/{id} [delete]
func (c *MovieController) DeleteMovie(ctx *gin.Context) {
	id, err := strconv.ParseUint(ctx.Param("id"), 10, 64)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, responses.NewErrorResponse("Invalid movie ID"))
		return
	}

	if err := c.service.DeleteMovie(uint(id)); err != nil {
		ctx.JSON(err.(*exceptions.AppError).Code, responses.NewErrorResponse(err.Error()))
		return
	}

	ctx.Status(http.StatusOK)
}
