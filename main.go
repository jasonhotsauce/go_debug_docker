package main

import (
	"fmt"

	"github.com/go-redis/redis"
)

// NewExampleClient creates a redis client and ping it to make sure it can talk to the server.
func NewExampleClient() (*redis.Client, error) {
	client := redis.NewClient(&redis.Options{
		Addr:     "redis:6379",
		Password: "",
		DB:       0,
	})
	_, err := client.Ping().Result()
	return client, err
}

func main() {
	client, err := NewExampleClient()
	if err != nil {
		panic(err)
	}
	defer client.Close()
	err = client.Set("KEY1", "VALUE", 0).Err()
	if err != nil {
		panic(err)
	}

	result := client.Get("KEY1")
	if err = result.Err(); err != nil {
		panic(err)
	}
	fmt.Println(result.Val())
}
