package utils

import (
	"errors"
	"fmt"
	"strings"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type JwtClaim struct {
	Iss string
	Sub string
	Exp int64
	Iat string
}

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

// Parses jwt and returns the claim subject
func ParseJwtToken(tokenString string, key string) (string, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (any, error) {
		// Validate the signing method
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(key), nil
	})

	if err != nil {
		return "", err
	}

	if !token.Valid {
		return "", errors.New("Token is invalid")
	}

	claims, ok := token.Claims.(jwt.MapClaims)

	if !ok {
		return "", fmt.Errorf("failed to extract claims")
	}

	return claims.GetSubject()

}

func ExtractBearerToken(token string) string {
	return strings.Replace(token, "Bearer ", "", 1)
}
