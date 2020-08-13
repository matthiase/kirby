package user

import (
	"kirby/httputil"
	"net/http"
	"strconv"

	"github.com/go-chi/chi"
)

// FetchUserResponse struct
type FetchUserResponse struct {
	Name  string `json:"name,omitempty"`
	Email string `json:"email,omitempty"`
}

// FetchUser fetches a user record based
func FetchUser(userService ServiceInterface) func(http.ResponseWriter, *http.Request) {
	handler := func(w http.ResponseWriter, r *http.Request) {
		userID, err := strconv.ParseUint(chi.URLParam(r, "id"), 10, 32)
		if err != nil {
			httputil.RespondWithError(w, http.StatusBadRequest, err)
		}

		user, err := userService.Find(uint(userID))
		if err != nil {
			httputil.RespondWithError(w, http.StatusNotFound, err)
			return
		}

		fetchUserResponse := FetchUserResponse{
			Name:  user.Name,
			Email: user.Email,
		}
		httputil.RespondWithJSON(w, http.StatusOK, fetchUserResponse)
	}
	return handler
}
