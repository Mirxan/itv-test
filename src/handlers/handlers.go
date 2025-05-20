package handlers

import (
	"github.com/Mirxan/itv-test/controllers"
	"github.com/Mirxan/itv-test/middleware"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func NewGinEngine(
	authController *controllers.AuthController,
	movieController *controllers.MovieController,
	authMiddleware *middleware.AuthMiddleware,
) *gin.Engine {
	router := gin.Default()
	router.Use(gin.Logger())
	router.Use(gin.Recovery())

	api := router.Group("/api")
	// Swagger docs
	api.GET("/docs/*any", gin.BasicAuth(gin.Accounts{
		"admin": "secret",
	}), ginSwagger.WrapHandler(swaggerFiles.Handler, ginSwagger.URL("/api/docs/doc.json")))

	// Handlers
	AuthHandlers(api, authController)
	Moviehandlers(api, movieController, authMiddleware)

	return router
}
