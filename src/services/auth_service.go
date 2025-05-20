package services

import (
	"errors"

	"github.com/Mirxan/itv-test/models"
	"github.com/Mirxan/itv-test/repositories"
	"github.com/Mirxan/itv-test/utils"

	"golang.org/x/crypto/bcrypt"
)

type AuthService interface {
	Register(username, email, password string) error
	Login(username, password string) (string, *models.User, error)
}

type authService struct {
	userRepo *repositories.UserRepository
	jwtUtil  *utils.JWTUtil
}

func NewAuthService(userRepo *repositories.UserRepository, jwtUtil *utils.JWTUtil) AuthService {
	return &authService{
		userRepo: userRepo,
		jwtUtil:  jwtUtil,
	}
}

func (s *authService) Register(username, email, password string) error {
	// Check if username exists
	if _, err := s.userRepo.FindByUsername(username); err == nil {
		return errors.New("username already exists")
	}

	// Check if email exists
	if _, err := s.userRepo.FindByEmail(email); err == nil {
		return errors.New("email already exists")
	}

	// Hash password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	// Create user
	user := &models.User{
		Username: username,
		Email:    email,
		Password: string(hashedPassword),
	}

	return s.userRepo.Create(user)
}

func (s *authService) Login(username, password string) (string, *models.User, error) {
	// Find user
	user, err := s.userRepo.FindByUsername(username)
	if err != nil {
		return "", nil, errors.New("invalid credentials")
	}

	// Check password
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
		return "", nil, errors.New("invalid credentials")
	}

	// Generate JWT token
	token, err := s.jwtUtil.GenerateToken(user)
	if err != nil {
		return "", nil, err
	}

	return token, user, nil
}
