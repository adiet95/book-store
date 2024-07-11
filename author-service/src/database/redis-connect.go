package database

import (
	"context"
	"github.com/go-redis/redis/v8"
	"os"
	"strconv"
	"time"
)

func NewRedisClient() (*redis.Client, error) {
	var opts = &redis.Options{
		Addr:         os.Getenv("REDIS_POOL_URI"),
		Password:     os.Getenv("REDIS_POOL_PASSWORD"),
		DB:           toInteger(os.Getenv("REDIS_POOL_DB")),
		DialTimeout:  envAsDuration(os.Getenv("REDIS_POOL_DIAL_TIMEOUT"), 5*time.Second) * time.Second,
		ReadTimeout:  envAsDuration(os.Getenv("REDIS_POOL_READ_TIMEOUT"), 5*time.Second) * time.Second,
		WriteTimeout: envAsDuration(os.Getenv("REDIS_POOL_WRITE_TIMEOUT"), 5*time.Second) * time.Second,
		IdleTimeout:  envAsDuration(os.Getenv("REDIS_POOL_IDLE_TIMEOUT"), 5*time.Second) * time.Second,
		MaxConnAge:   envAsDuration(os.Getenv("REDIS_POOL_MAX_CONN_AGE"), 5*time.Second) * time.Second,
	}
	rc := redis.NewClient(opts)
	ctx, cancel := context.WithTimeout(context.Background(), envAsDuration(os.Getenv("REDIS_POOL_DIAL_TIMEOUT"), 5*time.Second)*time.Second)
	defer cancel()
	if err := rc.Ping(ctx).Err(); err != nil {
		return nil, err
	}
	return rc, nil
}

func toInteger(str string) int {
	toInt, _ := strconv.Atoi(str)
	return toInt
}

func envAsDuration(key string, defaultVal time.Duration) time.Duration {
	strVal := key
	if val, err := time.ParseDuration(strVal); err == nil {
		return val
	}
	return defaultVal
}
