package services

import (
	"errors"
	"log"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/sam33339999/go-be-starter/domains"
	"github.com/sam33339999/go-be-starter/lib"
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
	token, err := jwt.Parse(tokenString, func(t *jwt.Token) (interface{}, error) {
		return []byte(s.env.JWTSecret), nil
	})

	if token.Valid {
		return true, nil
	} else if errors.Is(err, jwt.ErrTokenMalformed) {
		s.logger.Warnln("That's not even a token: ", token)
	} else if errors.Is(err, jwt.ErrTokenSignatureInvalid) {
		// 不可用的簽名
		s.logger.Warnln("Invalid signature: ", token)
	} else if errors.Is(err, jwt.ErrTokenExpired) || errors.Is(err, jwt.ErrTokenNotValidYet) {
		// Token 處在不可用的狀態: 過期或是尚未生效
		s.logger.Warnln("Timing is everything", token)
	} else {
		// 無法處理的 Token
		s.logger.Warnf("Couldn't handle this token: %+v\n", err)
	}
	return false, errors.New("count not handle token")
}

func (s AuthService) CreateToken() string {
	mySigningKey := s.env.JWTSecret

	claims := &jwt.RegisteredClaims{
		ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * 72)),
		Issuer:    s.env.AppName,
		IssuedAt:  jwt.NewNumericDate(time.Now()),
	}

	// 目前使用最簡單的 HS256 簽名方法
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// 進行簽名時，需要將對應的字串轉為 []byte
	ss, err := token.SignedString([]byte(mySigningKey))

	if err != nil {
		log.Fatalf("Error signing token: %v", err)
	}
	return ss
}
