package config

import (
	"log"
	"os"
	"path/filepath"
	"runtime"
	"strconv"
	"time"

	"github.com/joho/godotenv"
)

// Environment struct
type environment struct {
	Name                   string
	Host                   string
	Port                   uint
	ServerTimeout          time.Duration
	ReadTimeout            time.Duration
	WriteTimeout           time.Duration
	IdleTimeout            time.Duration
	PostgresURI            string
	RedisURI               string
	RedisPassword          string
	JwtSecret              string
	JwtAccessTokenTimeout  time.Duration
	JwtRefreshTokenTimeout time.Duration
}

// Env contains the globally accessible environment configuration
var Env environment

// LoadEnv loads project .env files
func LoadEnv() {
	environmentName := lookup("KIRBY_ENV", "development").(string)
	_, caller, _, ok := runtime.Caller(0)
	if !ok {
		log.Fatalln("Failed to load environment configuration")
	}
	basepath := filepath.Dir(caller) + "/../"

	godotenv.Load(basepath + ".env." + environmentName + ".local")
	if environmentName != "test" {
		godotenv.Load(basepath + ".env.local")
	}
	godotenv.Load(basepath + ".env." + environmentName)
	godotenv.Load(basepath + ".env")

	Env = environment{
		Name:                   environmentName,
		Host:                   lookup("HOST", "localhost").(string),
		Port:                   lookup("PORT", uint(8000)).(uint),
		ServerTimeout:          lookup("TIMEOUT_SERVERx", 30*time.Second).(time.Duration),
		ReadTimeout:            lookup("TIMEOUT_READ", 15*time.Second).(time.Duration),
		WriteTimeout:           lookup("TIMEOUT_WRITE", 10*time.Second).(time.Duration),
		IdleTimeout:            lookup("TIMEOUT_IDLE", 5*time.Second).(time.Duration),
		PostgresURI:            lookup("POSTGRES_URI", "postgres://postgres:@127.0.0.1:5432/kirby_"+environmentName+"?sslmode=disable").(string),
		RedisURI:               lookup("REDIS_URI", "redis://127.0.0.1:6379/0").(string),
		RedisPassword:          lookup("REDIS_PASSWORD", "").(string),
		JwtSecret:              lookup("JWT_SECRET", "kirby_jwt_secret").(string),
		JwtAccessTokenTimeout:  lookup("JWT_ACCESS_TOKEN_TIMEOUT", 15*time.Minute).(time.Duration),
		JwtRefreshTokenTimeout: lookup("JWT_REFRESH_TOKEN_TIMEOUT", 10*24*time.Hour).(time.Duration),
	}
}

func lookup(key string, fallback interface{}) interface{} {
	var value interface{}
	if env, ok := os.LookupEnv(key); ok {
		value = env
	} else {
		return fallback
	}

	switch fallback.(type) {
	case string:
		return value.(string)
	case uint:
		v, _ := strconv.ParseUint(value.(string), 10, 64)
		return uint(v)
	case time.Duration:
		var duration time.Duration
		if v, err := time.ParseDuration(value.(string)); err != nil {
			duration = v
		}
		return duration
	default:
		return value
	}
}
