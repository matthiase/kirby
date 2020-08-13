package user

import (
	"encoding/json"
	"kirby/errors"
	"kirby/httputil"
	"net/http"
)

// RefreshTokenRequest struct
type RefreshTokenRequest struct {
	RefreshToken string `json:"refreshToken" validate:"required"`
}

// Validate a RefreshTokenRequest
func (r *RefreshTokenRequest) Validate() error {
	return httputil.Validate().Struct(r)
}

// RefreshTokenResponse struct
type RefreshTokenResponse struct {
	AccessToken string `json:"accessToken,omitempty"`
}

// RefreshToken refreshes the JWT access token
func RefreshToken(userService ServiceInterface) func(http.ResponseWriter, *http.Request) {
	handler := func(w http.ResponseWriter, r *http.Request) {
		refreshTokenRequest := RefreshTokenRequest{}
		if err := json.NewDecoder(r.Body).Decode(&refreshTokenRequest); err != nil {
			httputil.RespondWithError(w, http.StatusBadRequest, err)
		}

		accessToken, err := userService.RefreshAccessToken(refreshTokenRequest)
		if err != nil {
			var status uint32
			switch err.(type) {
			case errors.ValidationError:
				status = http.StatusBadRequest
			case errors.AuthenticationError:
				status = http.StatusUnauthorized
			default:
				status = http.StatusInternalServerError
			}
			httputil.RespondWithError(w, status, err)
			return
		}
		httputil.RespondWithJSON(w, http.StatusOK, RefreshTokenResponse{accessToken})
	}
	return handler
}
