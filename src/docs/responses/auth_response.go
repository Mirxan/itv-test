package responses

import "github.com/Mirxan/itv-test/models"

type SuccessResponse struct {
	Message string `json:"message"`
}

func NewSuccessResponse(message string) *SuccessResponse {
	return &SuccessResponse{Message: message}
}

type LoginResponse struct {
	Token string       `json:"token"`
	User  UserResponse `json:"user"`
}

func NewLoginResponse(token string, user *models.User) *LoginResponse {
	return &LoginResponse{
		Token: token,
		User:  *NewUserResponse(user),
	}
}
