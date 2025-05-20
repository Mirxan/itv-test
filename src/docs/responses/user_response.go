package responses

import (
	"github.com/Mirxan/itv-test/models"
)

type UserResponse struct {
	ID       int64  `json:"id"`
	Email    string `json:"email"`
	Username string `json:"username"`
}

func NewUserResponse(user *models.User) *UserResponse {
	return &UserResponse{
		ID:       int64(user.ID),
		Email:    user.Email,
		Username: user.Username,
	}
}
