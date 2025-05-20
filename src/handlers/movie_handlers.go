package handlers

import (
	"github.com/Mirxan/itv-test/controllers"
	"github.com/Mirxan/itv-test/middleware"
	"github.com/gin-gonic/gin"
)

func Moviehandlers(r *gin.RouterGroup, movieController *controllers.MovieController, authMiddleware *middleware.AuthMiddleware) {
	movies := r.Group("/movies")
	movies.Use(authMiddleware.Authenticate())
	{
		movies.POST("/", movieController.CreateMovie)
		movies.GET("/", movieController.GetMovies)
		movies.GET("/:id", movieController.GetMovie)
		movies.PUT("/:id", movieController.UpdateMovie)
		movies.DELETE("/:id", movieController.DeleteMovie)
	}
}
