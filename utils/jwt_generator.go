package utils

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
)

func GenerateJwtToken(userId string, appName string, key string, exp time.Duration) (string, error) {
	claim := jwt.MapClaims{
		"iss": appName,
		"sub": userId,
		"exp": exp,
		"iat": time.Now().Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claim)
	return token.SignedString([]byte(key))
}
