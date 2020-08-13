package database

import (
	"kirby/config"

	"github.com/go-redis/redis"
)

var (
	// Redis holds information about the database connection
	Redis *redis.Client
)

// RedisConnect opens a connection to the database
func RedisConnect() (*redis.Client, error) {
	options, err := redis.ParseURL(config.Env.RedisURI)
	if err != nil {
		return nil, err
	}
	options.Password = config.Env.RedisPassword
	redis := redis.NewClient(options)
	if _, err := redis.Ping().Result(); err != nil {
		return nil, err
	}
	Redis = redis
	return Redis, nil
}
