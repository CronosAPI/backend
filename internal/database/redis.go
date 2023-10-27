package database

import (
	"os"

	"github.com/go-redis/redis/v8"
)

var Client *redis.Client

func InitializeConnection() *redis.Client {
	Client = redis.NewClient(&redis.Options{
		Addr:     os.Getenv("ADDRESS"),
		Password: os.Getenv("PASSWORD"),
		DB:       0,
	})

	return Client
}
