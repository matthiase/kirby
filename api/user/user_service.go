package user

import (
	"encoding/json"
	"kirby/config"
	"kirby/database"
	"kirby/errors"
	"net/http"

	"github.com/go-redis/redis"

	"github.com/jinzhu/gorm"
	"golang.org/x/crypto/bcrypt"
)

// ServiceInterface declaration
type ServiceInterface interface {
	Find(uint) (*User, error)
	Create(*CreateUserRequest) (*User, error)
	Authenticate(AuthenticationRequest) (*TokenPair, error)
	RefreshAccessToken(RefreshTokenRequest) (string, error)
}

// Service struct
type Service struct {
	DB    *gorm.DB
	Redis *redis.Client
}

// Find a user by id
func (s *Service) Find(id uint) (*User, error) {
	user := User{}
	if err := s.DB.Where("id = ?", id).First(&user).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return &User{}, errors.ApplicationError{
				Status:  http.StatusNotFound,
				Message: "User not found",
			}
		}
		return &user, err
	}
	return &user, nil
}

// Create a new user record
func (s *Service) Create(createUserRequest *CreateUserRequest) (*User, error) {
	if err := createUserRequest.Validate(); err != nil {
		return &User{}, errors.ApplicationError{
			Status:  http.StatusBadRequest,
			Message: "User request validation failed",
		}
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(createUserRequest.Password), bcrypt.DefaultCost)
	if err != nil {
		return &User{}, err
	}

	user := User{
		Name:           createUserRequest.Name,
		Email:          createUserRequest.Email,
		HashedPassword: string(hashedPassword),
	}

	if err := s.DB.Create(&user).Error; err != nil {
		if database.IsUniqueConstraintError(err, database.UniqueConstraintUserEmail) {
			return &User{}, errors.ApplicationError{
				Status:  http.StatusBadRequest,
				Source:  "user/email",
				Message: "A user with that email address already exists",
			}
		}
		return &User{}, err
	}
	return &user, nil
}

// Authenticate a user using their credentials and return a JWT token pair
func (s *Service) Authenticate(authenticationRequest AuthenticationRequest) (*TokenPair, error) {
	if err := authenticationRequest.Validate(); err != nil {
		return &TokenPair{}, errors.ApplicationError{
			Status:  http.StatusBadRequest,
			Message: "Authentication request validation failed",
		}
	}

	user := User{}
	if err := s.DB.Where("email = ?", authenticationRequest.Email).First(&user).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return &TokenPair{}, errors.ApplicationError{
				Status:  http.StatusUnauthorized,
				Message: "Invalid email address or password",
			}
		}
		return &TokenPair{}, err
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.HashedPassword), []byte(authenticationRequest.Password)); err != nil {
		if err == bcrypt.ErrMismatchedHashAndPassword {
			return &TokenPair{}, errors.ApplicationError{
				Status:  http.StatusUnauthorized,
				Message: "Invalid email address or password",
			}
		}
		return &TokenPair{}, err
	}

	tokenPair, err := user.GenerateTokenPair()
	if err != nil {
		return &TokenPair{}, err
	}

	json, err := json.Marshal(user)
	if err != nil {
		return &TokenPair{}, err
	}

	if err := s.Redis.Set(tokenPair.RefreshToken, json, config.Env.JwtRefreshTokenTimeout).Err(); err != nil {
		return &TokenPair{}, err
	}
	return tokenPair, nil
}

// RefreshAccessToken generate a new access token
func (s *Service) RefreshAccessToken(refreshTokenRequest RefreshTokenRequest) (string, error) {
	if err := refreshTokenRequest.Validate(); err != nil {
		return "", errors.ApplicationError{
			Status:  http.StatusBadRequest,
			Message: "Refresh token request validation failed",
		}
	}

	payload, err := s.Redis.Get(refreshTokenRequest.RefreshToken).Result()
	if err != nil {
		return "", errors.ApplicationError{
			Status:  http.StatusUnauthorized,
			Message: "Invalid or expired refresh token",
		}
	}

	user := User{}
	if err := json.Unmarshal([]byte(payload), &user); err != nil {
		return "", errors.ApplicationError{
			Status:  http.StatusUnauthorized,
			Message: "Invalid or expired refresh token",
		}
	}

	accessToken, err := user.GenerateAccessToken()
	if err != nil {
		return "", err
	}
	return accessToken, nil
}
