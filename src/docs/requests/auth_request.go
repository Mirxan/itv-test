package requests

type RegisterRequest struct {
	Username string `json:"username" binding:"required,min=3,max=50" example:"username"`
	Email    string `json:"email" binding:"required,email" example:"user@example.com"`
	Password string `json:"password" binding:"required,min=6" example:"secret"`
}

type LoginRequest struct {
	Username string `json:"username" binding:"required" example:"username"`
	Password string `json:"password" binding:"required" example:"secret"`
}
