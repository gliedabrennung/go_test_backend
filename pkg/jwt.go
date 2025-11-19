package pkg

import (
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type Token struct {
	RawToken  string
	Method    jwt.SigningMethod
	Header    map[string]any
	Claims    jwt.Claims
	Signature []byte
	Valid     bool
}

type CustomClaims struct {
	Username string `json:"user"`
	Role     string `json:"role"`
	jwt.RegisteredClaims
}

func GenerateSignedToken(username string, secretKey []byte) (string, error) {
	exp := time.Now().Add(time.Hour * 1)
	claims := jwt.MapClaims{
		"user": username,
		"exp":  jwt.NewNumericDate(exp),
		"iat":  jwt.NewNumericDate(time.Now()),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(secretKey)
	if err != nil {
		return "", fmt.Errorf("error signing token: %w", err)
	}

	return tokenString, nil
}
