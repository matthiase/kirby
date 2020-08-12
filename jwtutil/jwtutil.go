package jwtutil

import (
	"errors"
	"fmt"
	"kirby/config"

	"github.com/dgrijalva/jwt-go"
)

// UserClaims struct
type UserClaims struct {
	ID    uint   `json:"id,omitempty"`
	Name  string `json:"name,omitempty"`
	Email string `json:"email,omitempty"`
	jwt.StandardClaims
}

// Encode claims
func Encode(claims *UserClaims) (string, error) {
	secret := []byte(config.Env.JwtSecret)
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(secret)
}

// Decode a token string
func Decode(tokenString string) (*UserClaims, error) {
	jwtToken, _ := jwt.ParseWithClaims(tokenString, &UserClaims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(config.Env.JwtSecret), nil
	})

	if claims, ok := jwtToken.Claims.(*UserClaims); ok && jwtToken.Valid {
		return claims, nil
	}

	return &UserClaims{}, errors.New("Invalid token")
}
