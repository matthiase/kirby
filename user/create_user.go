package user

import (
	"encoding/json"
	"kirby/errors"
	"kirby/httputil"
	"net/http"
)

// CreateUserRequest struct
type CreateUserRequest struct {
	Name     string `json:"name" validate:"required"`
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,gte=8"`
}

// CreateUserResponse struct
type CreateUserResponse struct {
	AccessToken  string `json:"accessToken"`
	RefreshToken string `json:"refreshToken,omitempty"`
}

// Validate the create user request
func (r *CreateUserRequest) Validate() error {
	return httputil.Validate().Struct(r)
}

// CreateUser creates a new user
func CreateUser(userService ServiceInterface) func(http.ResponseWriter, *http.Request) {
	handler := func(w http.ResponseWriter, r *http.Request) {
		createUserRequest := CreateUserRequest{}
		if err := json.NewDecoder(r.Body).Decode(&createUserRequest); err != nil {
			httputil.RespondWithError(w, http.StatusBadRequest, err)
			return
		}

		user, err := userService.Create(createUserRequest)
		if err != nil {
			switch err.(type) {
			case errors.ValidationError:
				httputil.RespondWithError(w, http.StatusBadRequest, err)
			default:
				httputil.RespondWithError(w, http.StatusInternalServerError, err)
			}
			return
		}

		tokenPair, err := user.GenerateTokenPair()
		if err != nil {
			httputil.RespondWithError(w, http.StatusInternalServerError, err)
			return
		}

		CreateUserResponse := CreateUserResponse{
			AccessToken:  tokenPair.AccessToken,
			RefreshToken: tokenPair.RefreshToken,
		}
		httputil.RespondWithJSON(w, http.StatusCreated, CreateUserResponse)
	}
	return handler
}
