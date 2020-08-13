package user

import (
	"encoding/json"
	"kirby/errors"
	"kirby/httputil"
	"net/http"
)

// AuthenticationRequest struct
type AuthenticationRequest struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required"`
}

// AuthenticationResponse struct
type AuthenticationResponse struct {
	AccessToken  string `json:"accessToken"`
	RefreshToken string `json:"refreshToken,omitempty"`
}

// Validate login request
func (r *AuthenticationRequest) Validate() error {
	return httputil.Validate().Struct(r)
}

// Authenticate creates a new session
func Authenticate(userService ServiceInterface) func(http.ResponseWriter, *http.Request) {
	handler := func(w http.ResponseWriter, r *http.Request) {
		authenticationRequest := AuthenticationRequest{}
		if err := json.NewDecoder(r.Body).Decode(&authenticationRequest); err != nil {
			httputil.RespondWithError(w, http.StatusBadRequest, err)
			return
		}

		tokenPair, err := userService.Authenticate(authenticationRequest)
		if err != nil {
			switch err.(type) {
			case errors.AuthenticationError:
				httputil.RespondWithError(w, http.StatusUnauthorized, err)
			default:
				httputil.RespondWithError(w, http.StatusInternalServerError, err)
			}
			return
		}

		authenticationResponse := AuthenticationResponse{
			AccessToken:  tokenPair.AccessToken,
			RefreshToken: tokenPair.RefreshToken,
		}
		httputil.RespondWithJSON(w, http.StatusOK, authenticationResponse)
	}

	return handler
}
