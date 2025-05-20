package utils

import (
	"errors"
	"time"

	"github.com/Mirxan/itv-test/config"
	"github.com/Mirxan/itv-test/models"
	"github.com/golang-jwt/jwt/v5"
)

type JWTUtil struct {
	secret string
}

func NewJWTUtil(config *config.Config) *JWTUtil {
	return &JWTUtil{secret: config.JWTSecret}
}

type Claims struct {
	ID uint `json:"id"`
	jwt.RegisteredClaims
}

func (j *JWTUtil) GenerateToken(user *models.User) (string, error) {
	claims := &Claims{
		ID: user.ID,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(24 * time.Hour)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			Issuer:    "movie-api",
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(j.secret))
}

func (j *JWTUtil) ValidateToken(tokenString string) (*Claims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(j.secret), nil
	})

	if err != nil {
		return nil, err
	}

	if claims, ok := token.Claims.(*Claims); ok && token.Valid {
		return claims, nil
	}

	return nil, errors.New("invalid token")
}
