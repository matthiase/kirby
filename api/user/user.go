package user

import (
	"context"
	"kirby/config"
	"kirby/jwtutil"
	"time"

	"github.com/adjust/uniuri"
	"github.com/dgrijalva/jwt-go"
	"github.com/jinzhu/gorm"
)

// User struct
type User struct {
	gorm.Model
	Email          string `gorm:"unique_index;not null"`
	HashedPassword string `gorm:"not null"`
	Name           string `gorm:"not null"`
}

type contextKey string

const userCtxKey contextKey = "currentUser"

// NewContext returns a new context that carries the provided user value
func NewContext(ctx context.Context, user User) context.Context {
	return context.WithValue(ctx, userCtxKey, user)
}

// FromContext extracts a user from a context
func FromContext(ctx context.Context) (User, bool) {
	user, ok := ctx.Value(userCtxKey).(User)
	return user, ok
}

// GenerateAccessToken creates a new JWT access token
func (u *User) GenerateAccessToken() (string, error) {
	claims := jwtutil.UserClaims{
		ID:    u.ID,
		Name:  u.Name,
		Email: u.Email,
		StandardClaims: jwt.StandardClaims{
			Subject:   u.Email,
			ExpiresAt: time.Now().Add(config.Env.JwtAccessTokenTimeout).Unix(),
		},
	}

	accessToken, err := jwtutil.Encode(&claims)
	if err != nil {
		return "", err
	}

	return accessToken, nil
}

// GenerateTokenPair creates a new JWT access and refresh token
func (u *User) GenerateTokenPair() (*TokenPair, error) {
	accessToken, err := u.GenerateAccessToken()
	if err != nil {
		return &TokenPair{}, err
	}

	refreshToken := uniuri.NewLen(24)
	tokenPair := TokenPair{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}

	return &tokenPair, nil
}

// TokenPair struct
type TokenPair struct {
	AccessToken  string `json:"accessToken,omitempty"`
	RefreshToken string `json:"refreshToken,omitempty"`
}
