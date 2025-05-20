package controllers

import (
	"net/http"

	"github.com/Mirxan/itv-test/docs/requests"
	"github.com/Mirxan/itv-test/docs/responses"
	"github.com/Mirxan/itv-test/exceptions"
	"github.com/Mirxan/itv-test/services"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type AuthController struct {
	authService services.AuthService
}

func NewAuthController(authService services.AuthService) *AuthController {
	return &AuthController{authService: authService}
}

// Register godoc
// @Summary Register a new user
// @Description Register a new user with username, email and password
// @Tags auth
// @Accept json
// @Produce json
// @Param request body requests.RegisterRequest true "Register user"
// @Success 201 {object} responses.SuccessResponse
// @Failure 400 {object} responses.ErrorResponse
// @Router /auth/register [post]
func (c *AuthController) Register(ctx *gin.Context) {
	var req requests.RegisterRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, exceptions.NewValidationErrors(err))
		return
	}

	if err := c.authService.Register(req.Username, req.Email, req.Password); err != nil {
		status := http.StatusInternalServerError
		if err.Error() == "username already exists" || err.Error() == "email already exists" {
			status = http.StatusConflict
		}
		ctx.JSON(status, responses.NewErrorResponse(err.Error()))
		return
	}

	ctx.JSON(http.StatusCreated, responses.NewSuccessResponse("Registered successfully"))
}

// Login godoc
// @Summary Login a user
// @Description Login with username and password
// @Tags auth
// @Accept json
// @Produce json
// @Param credentials body requests.LoginRequest true "Login credentials"
// @Success 200 {object} responses.LoginResponse
// @Failure 400 {object} responses.ErrorResponse
// @Failure 401 {object} responses.ErrorResponse
// @Failure 500 {object} responses.ErrorResponse
// @Router /auth/login [post]
func (c *AuthController) Login(ctx *gin.Context) {
	var req requests.LoginRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, exceptions.NewValidationErrors(err.(validator.ValidationErrors)))
		return
	}

	token, user, err := c.authService.Login(req.Username, req.Password)
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, responses.NewErrorResponse(err.Error()))
		return
	}

	ctx.JSON(http.StatusOK, responses.NewLoginResponse(token, user))
}
