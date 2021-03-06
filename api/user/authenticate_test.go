package user

import (
	"bytes"
	"encoding/json"
	"kirby/database"
	"kirby/httputil"
	"kirby/jwtutil"
	"kirby/testutil"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestAuthenticate(t *testing.T) {
	testutil.LoadEnv()
	validUser := CreateUserRequest{
		Name:     "Valid User",
		Email:    "valid@example.com",
		Password: "secretpwd",
	}

	tests := []struct {
		name               string
		payload            string
		expectedStatusCode int
	}{
		{"Invalid email address", `{"email": "invalid@example.com", "password": "secretpwd"}`, http.StatusUnauthorized},
		{"Invalid password", `{"email": "valid@example.com", "password": "invalid"}`, http.StatusUnauthorized},
		{"Valid credentials", `{"email": "valid@example.com", "password": "secretpwd"}`, http.StatusOK},
	}

	pg, err := database.PgConnect()
	if err != nil {
		log.Fatalf("Postgres connection failed: %v\n", err)
	}

	pg.AutoMigrate(&User{})

	redis, err := database.RedisConnect()
	if err != nil {
		log.Fatalf("Redis connection failed: %v\n", err)
	}

	userService := &Service{DB: pg, Redis: redis}
	handler := http.HandlerFunc(Authenticate(userService))

	setup := func() error {
		pg.Exec("TRUNCATE TABLE users")
		_, err := userService.Create(&validUser)
		return err
	}

	validateResponseBody := func(r *httptest.ResponseRecorder) {
		if r.Code == http.StatusOK {
			response := httputil.SuccessResponse{}
			if err := json.Unmarshal(r.Body.Bytes(), &response); err != nil || response.Data == nil {
				t.Errorf("handler returned an invalid success response")
			}
			data := response.Data.(map[string]interface{})
			if accessToken, ok := data["accessToken"]; ok {
				claims, err := jwtutil.Decode(string(accessToken.(string)))
				if err != nil {
					t.Errorf("handler returned an invalid access token")
				}
				email := claims.Email
				if email != validUser.Email {
					t.Errorf("handler returned an invalid access token: got %v want %v", email, validUser.Email)
				}
			}
		}
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			if err := setup(); err != nil {
				t.Errorf("Test setup failed: %v\n", err)
			}
			request, _ := http.NewRequest("POST", "/tokens", bytes.NewReader([]byte(test.payload)))
			recorder := httptest.NewRecorder()
			handler.ServeHTTP(recorder, request)
			if recorder.Code != test.expectedStatusCode {
				t.Errorf("handler returned incorrect status code: got %v want %v", recorder.Code, test.expectedStatusCode)
			}
			validateResponseBody(recorder)
		})
	}
}
