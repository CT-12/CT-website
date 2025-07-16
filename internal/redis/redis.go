package redis

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/go-redis/redis/v8"
	"github.com/joho/godotenv"
)

var(
	ctx = context.Background()
)

func newRedisClient() (*redis.Client, error)  {
	if err := godotenv.Load(); err != nil {
		log.Printf("Error loading .env file -> %v", err)
	}

	redis_endpoint := os.Getenv("REDIS_ENDPOINT")
	if redis_endpoint == "" {
		return nil, fmt.Errorf("redis endpoint is not found in env")
	}

	redis_password := os.Getenv("REDIS_PASSWORD")
	if redis_password == "" {
		return nil, fmt.Errorf("redis password is not found in env")
	}

	// Setting up Redis client
	redis_client := redis.NewClient(&redis.Options{
		Addr: redis_endpoint,
		Password: redis_password,
		DB: 0, // Default DB
	})


	// Check whether the connection is successful
	_, err := redis_client.Ping(ctx).Result()
	if err != nil {
		return nil, fmt.Errorf("error connecting to redis -> %v", err)
	}

	log.Printf("Connected to Redis successfully!")

	return redis_client, nil
}

func SetValue(key string, value string, expire_days time.Duration) error {
	redis_client, err := newRedisClient()
	if err != nil {
		return fmt.Errorf("error creating Redis client -> %v", err)
	}
	defer redis_client.Close()

	_, err = redis_client.Get(ctx, key).Result()
	if err == nil || err == redis.Nil {
		// if err == nil, key exists, update value
		// if err == redis.Nil, key does not exist, set new value
		err := redis_client.Set(ctx, key, value, time.Hour*24*expire_days).Err()
		return err
	} else {
		return err
	} 
}

func GetValue(key string) (string, error) {
	redis_client, err := newRedisClient()
	if err != nil {
		return "", fmt.Errorf("error creating redis client -> %v", err)
	}
	defer redis_client.Close()

	value, err := redis_client.Get(ctx, key).Result()
	if err == redis.Nil {
		return "", nil // Key does not exist
	} else if err != nil {
		return "", err // Other error
	}

	return value, nil
}