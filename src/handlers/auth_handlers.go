package handlers

import (
	"github.com/Mirxan/itv-test/controllers"
	"github.com/gin-gonic/gin"
)

func AuthHandlers(r *gin.RouterGroup, authController *controllers.AuthController) {
	auth := r.Group("/auth")
	{
		auth.POST("/register", authController.Register)
		auth.POST("/login", authController.Login)
	}
}
