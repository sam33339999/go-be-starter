package services

import (
	"errors"

	"github.com/golang-jwt/jwt/v5"
	"github.com/sam33339999/go-be-starter/domains"
	"github.com/sam33339999/go-be-starter/lib"
	"github.com/sam33339999/go-be-starter/models"
)

type AuthService struct {
	env    lib.Env
	logger lib.Logger
}

func NewJWTAuthService(
	env lib.Env,
	logger lib.Logger,
) domains.AuthService {
	return AuthService{
		env:    env,
		logger: logger,
	}
}

func (s AuthService) Authorize(tokenString string) (bool, error) {

	token, _ := jwt.Parse(tokenString, func(t *jwt.Token) (interface{}, error) {
		return []byte(s.env.JWTSecret), nil
	})

	if token.Valid {
		return true, nil
	}

	return false, errors.New("count not handle token")
}

func (auth AuthService) CreateToken(user models.User) string {

	return ""
}
