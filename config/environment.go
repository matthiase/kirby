package config

import (
	"os"
	"strconv"
	"time"

	"github.com/joho/godotenv"
)

func init() {
	godotenv.Load()
}

// Host returns the hostname the application is running on
func Host() string {
	host := "127.0.0.1"
	if env, ok := os.LookupEnv("HOST"); ok {
		host = env
	}
	return host
}

// Port returns the port the application is listening on
func Port() uint64 {
	var port uint64 = 8000
	if env, ok := os.LookupEnv("PORT"); ok {
		if value, err := strconv.ParseUint(env, 10, 64); err == nil {
			port = value
		}
	}
	return port
}

// ServerTimeout returns the duration until a operation times out
func ServerTimeout() time.Duration {
	var duration time.Duration = 30 * time.Second
	if env, ok := os.LookupEnv("TIMEOUT_SERVER"); ok {
		if value, err := time.ParseDuration(env); err == nil {
			duration = value
		}
	}
	return duration
}

// ReadTimeout returns the duration until a read operation times out
func ReadTimeout() time.Duration {
	var duration time.Duration = 15 * time.Second
	if env, ok := os.LookupEnv("TIMEOUT_READ"); ok {
		if value, err := time.ParseDuration(env); err == nil {
			duration = value
		}
	}
	return duration
}

// WriteTimeout returns the duration until a write operation times out
func WriteTimeout() time.Duration {
	var duration time.Duration = 10 * time.Second
	if env, ok := os.LookupEnv("TIMEOUT_WRITE"); ok {
		if value, err := time.ParseDuration(env); err == nil {
			duration = value
		}
	}
	return duration
}

// IdleTimeout returns the duration until and idle operation times out
func IdleTimeout() time.Duration {
	var duration time.Duration = 5 * time.Second
	if env, ok := os.LookupEnv("TIMEOUT_IDLE"); ok {
		if value, err := time.ParseDuration(env); err == nil {
			duration = value
		}
	}
	return duration
}

// DatabaseURI returns the database URI as a string
func DatabaseURI() string {
	var uri string = ""
	if env, ok := os.LookupEnv("DATABASE_URI"); ok {
		uri = env
	}
	return uri
}

// JWTSecret used to sign tokens
func JWTSecret() string {
	var secret = ""
	if env, ok := os.LookupEnv("JWT_SECRET"); ok {
		secret = env
	}
	return secret
}

// JWTAccessTokenTimeout determines the amount of time before access tokens expire
func JWTAccessTokenTimeout() time.Duration {
	var duration time.Duration = 15 * time.Minute
	if env, ok := os.LookupEnv("JWT_ACCESS_TOKEN_TIMEOUT"); ok {
		if value, err := time.ParseDuration(env); err == nil {
			duration = value
		}
	}
	return duration
}

// JWTRefreshTokenTimeout determines the amount of time before access tokens expire
func JWTRefreshTokenTimeout() time.Duration {
	var duration time.Duration = 10 * 24 * time.Hour
	if env, ok := os.LookupEnv("JWT_REFRESH_TOKEN_TIMEOUT"); ok {
		if value, err := time.ParseDuration(env); err == nil {
			duration = value
		}
	}
	return duration
}
