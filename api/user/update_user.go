package user

import (
	"encoding/json"
	"kirby/errors"
	"kirby/httputil"
	"net/http"
)

// UpdateUserRequest struct
type UpdateUserRequest struct {
	Name  string `json:"name" validate:"required"`
	Email string `json:"email" validate:"required,email"`
}

// UpdateUserResponse struct
type UpdateUserResponse struct {
	AccessToken string `json:"accessToken"`
}

// Validate the create user request
func (r *UpdateUserRequest) Validate() error {
	return httputil.Validate().Struct(r)
}

// UpdateUser creates a new user
func UpdateUser(userService ServiceInterface) func(http.ResponseWriter, *http.Request) {
	handler := func(w http.ResponseWriter, r *http.Request) {
		updateUserRequest := UpdateUserRequest{}
		if err := json.NewDecoder(r.Body).Decode(&updateUserRequest); err != nil {
			httputil.RespondWithError(w, http.StatusBadRequest, err)
			return
		}

		currentUser, ok := FromContext(r.Context())
		if !ok {
			err := errors.ApplicationError{Message: "Invalid context for current user"}
			httputil.RespondWithError(w, http.StatusInternalServerError, err)
			return
		}

		user, err := userService.Update(currentUser.ID, &updateUserRequest)
		if err != nil {
			switch err.(type) {
			case errors.ApplicationError:
				httpStatusCode := err.(errors.ApplicationError).Status
				httputil.RespondWithError(w, httpStatusCode, err)
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

		updateUserResponse := UpdateUserResponse{
			AccessToken: tokenPair.AccessToken,
		}
		httputil.RespondWithJSON(w, http.StatusOK, updateUserResponse)
	}
	return handler
}
