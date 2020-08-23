package api

import (
	"errors"
	"kirby/api/user"
	"kirby/httputil"
	"kirby/jwtutil"
	"net/http"
	"strings"

	"github.com/jinzhu/gorm"
)

// JwtAuthentication authenticates requests using the token found in the Authentication header
func JwtAuthentication(next http.Handler) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		exempt := []string{
			"GET /health/server",
			"GET /health/database",
			"POST /users",
			"POST /tokens",
			"PUT /tokens",
		}

		path := r.URL.Path
		for _, v := range exempt {
			params := strings.Split(v, " ")
			if r.Method == params[0] && strings.EqualFold(path, params[1]) {
				next.ServeHTTP(w, r)
				return
			}
		}

		authenticationHeader := strings.Split(r.Header.Get("Authorization"), "Bearer ")
		if len(authenticationHeader) != 2 {
			httputil.RespondWithError(w, http.StatusUnauthorized, errors.New("Malformed authentication token"))
		} else {
			claims, err := jwtutil.Decode(authenticationHeader[1])
			if err != nil {
				httputil.RespondWithError(w, http.StatusUnauthorized, errors.New("Invalid or expired access token"))
				return
			}

			currentUser := user.User{
				Model: gorm.Model{ID: claims.ID},
				Name:  claims.Name,
				Email: claims.Email,
			}

			ctx := user.NewContext(r.Context(), currentUser)

			next.ServeHTTP(w, r.WithContext(ctx))
		}
	}
	return http.HandlerFunc(fn)
}
