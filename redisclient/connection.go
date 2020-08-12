package redisclient

import (
	"github.com/go-redis/redis"
)

// Connection wraps the underlying connection to the Redis datastore
type Connection struct {
	*redis.Client
}

// Connect opens a connection to the Redis datastore
func Connect(url string, password string) (*Connection, error) {
	options, err := redis.ParseURL(url)
	if err != nil {
		return nil, err
	}
	options.Password = password
	conn := redis.NewClient(options)

	if _, err := conn.Ping().Result(); err != nil {
		return nil, err
	}
	return &Connection{conn}, nil
}
