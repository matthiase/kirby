package redisclient

import (
	"github.com/go-redis/redis"
)

// Connection wraps the underlying connection to the Redis datastore
type Connection struct {
	*redis.Client
}

// Connect opens a connection to the Redis datastore
func Connect(addr string, password string, database int) (*Connection, error) {
	conn := redis.NewClient(&redis.Options{
		Addr:     addr,
		Password: password,
		DB:       database,
	})

	_, err := conn.Ping().Result()
	if err != nil {
		return nil, err
	}
	return &Connection{conn}, nil
}
