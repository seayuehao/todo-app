package util

import (
	"errors"
	"github.com/golang-jwt/jwt/v5"
	"github.com/todo-app/internal/config"
	"time"
)

func GenerateToken(userID string) (string, error) {
	jwtKey := []byte(config.AppCfg.JwtConfig.Secret)

	claims := jwt.RegisteredClaims{
		Subject:   userID,
		ExpiresAt: jwt.NewNumericDate(time.Now().Add(24 * time.Hour)),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(jwtKey)
}

func ParseToken(tokenStr string) (*jwt.RegisteredClaims, error) {
	jwtKey := []byte(config.AppCfg.JwtConfig.Secret)
	token, err := jwt.ParseWithClaims(
		tokenStr,
		&jwt.RegisteredClaims{},
		func(token *jwt.Token) (interface{}, error) {
			return jwtKey, nil
		},
	)

	if err != nil {
		return nil, err
	}

	if !token.Valid {
		return nil, errors.New("invalid token")
	}

	return token.Claims.(*jwt.RegisteredClaims), nil
}
