package services

import (
	"errors"
	"fmt"
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

const (
	TYPE_ACCESS  = "access"
	TYPE_REFRESH = "refresh"
)

func (s AuthService) genTypeToken(tokenType string) (string, error) {
	var secret string = ""
	var ttl int64 = 0
	if tokenType == TYPE_ACCESS {
		secret = s.env.JWTSecret
		ttl = s.env.JWTTtl
	}

	if tokenType == TYPE_REFRESH {
		secret = s.env.RefreshJWTSecret
		ttl = s.env.RefreshJWTTtl
	}

	if secret == "" || ttl == 0 {
		return "", errors.New("invalid token type or env not setting.")
	}

	claims := &jwt.RegisteredClaims{
		// 時間計算的寫法，如果是常數可以直接寫數字；但是如果是變數，則需要使用 time.Duration() 來轉換
		ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * time.Duration(ttl))),
		Issuer:    s.env.AppName,
		IssuedAt:  jwt.NewNumericDate(time.Now()),
	}

	// 目前使用 HS256 簽名方法
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// 進行簽名時，需要將對應的字串轉為 []byte
	ss, err := token.SignedString([]byte(secret))
	return ss, err
}

func (s AuthService) CreateToken() (string, string) {
	token, err := s.genTypeToken(TYPE_ACCESS)

	if err != nil {
		log.Fatalf("Error signing token: %v", err)
	}
	refreshToken, err := s.genTypeToken(TYPE_REFRESH)

	if err != nil {
		log.Fatalf("Error signing token: %v", err)
	}

	return token, refreshToken
}

func (s AuthService) RefreshToken(refreshToken string) (string, string, error) {
	token, err := jwt.Parse(refreshToken, func(token *jwt.Token) (interface{}, error) {
		return []byte(s.env.JWTSecret), nil
	})

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		fmt.Println(claims["foo"], claims["nbf"])
	} else {
		s.logger.Errorf("Couldn't handle this token: %+v\n", err)
	}
	return "", "", nil
}
