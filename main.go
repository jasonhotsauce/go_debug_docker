package main

import (
	"fmt"

	"github.com/go-redis/redis"
	"github.com/spf13/viper"
)

const redisKey = "redis"

// NewExampleClient creates a redis client and ping it to make sure it can talk to the server.
func NewExampleClient() (*redis.Client, error) {
	redisHost := viper.Get(redisKey)
	client := redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:6379", redisHost),
		Password: "",
		DB:       0,
	})
	_, err := client.Ping().Result()
	return client, err
}

func main() {
	viper.SetEnvPrefix("demo")
	viper.BindEnv(redisKey)
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
