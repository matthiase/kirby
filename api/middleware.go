package api

import (
	"context"
	"errors"
	"kirby/httputil"
	"kirby/jwtutil"
	"net/http"
	"strings"
)

var userCtxKey *contextKey = &contextKey{"currentUser"}

type contextKey struct {
	name string
}

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

			ctx := context.WithValue(r.Context(), userCtxKey, claims)
			next.ServeHTTP(w, r.WithContext(ctx))
		}
	}
	return http.HandlerFunc(fn)
}
