// @title Movie API
// @version 1.0
// @description This is a sample server for movies.
// @BasePath /api
//
// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authorization
// @description "Type: Bearer <JWT token>"
package main

import (
	"context"

	"github.com/Mirxan/itv-test/config"
	"github.com/Mirxan/itv-test/controllers"
	_ "github.com/Mirxan/itv-test/docs"
	"github.com/Mirxan/itv-test/handlers"
	"github.com/Mirxan/itv-test/middleware"
	"github.com/Mirxan/itv-test/repositories"
	"github.com/Mirxan/itv-test/services"
	"github.com/Mirxan/itv-test/utils"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"go.uber.org/fx"
)

func main() {
	app := fx.New(
		fx.Provide(
			// Configuration and utilities
			config.LoadConfig,
			utils.NewDBConnection,
			utils.NewJWTUtil,

			// Repositories
			repositories.NewUserRepository,
			repositories.NewMovieRepository,

			// Services
			services.NewAuthService,
			services.NewMovieService,

			// Controllers and middleware
			controllers.NewAuthController,
			controllers.NewMovieController,
			middleware.NewAuthMiddleware,

			// Router
			handlers.NewGinEngine,
		),
		fx.Invoke(registerHooks),
	)

	app.Run()
}

func registerHooks(lifecycle fx.Lifecycle, router *gin.Engine) {
	lifecycle.Append(
		fx.Hook{
			OnStart: func(context.Context) error {
				go router.Run(":" + viper.GetString("APP_PORT"))
				return nil
			},
		},
	)
}
