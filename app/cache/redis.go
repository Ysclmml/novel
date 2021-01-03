package cache

import (
	"fmt"
	"github.com/go-redis/redis"
	"github.com/spf13/viper"
	"novel/app/log"
	"time"
)

// Redis cache implement
type Redis struct {
	Client *redis.Client
}

// Setup connection
func (r *Redis) Connect() {
	r.Client = redis.NewClient(&redis.Options{
		Addr:         viper.GetString("redis.host"),
		Password:     viper.GetString("redis.auth"),
		DB:           viper.GetInt("redis.db"),
		PoolSize:     viper.GetInt("redis.poolSize"),
		MinIdleConns: viper.GetInt("redis.idleTimeout"),
		MaxRetries:   viper.GetInt("redis.maxRetries"),
		IdleTimeout:  viper.GetDuration("redis.idleTimeout"),
	})
	_, err := r.Client.Ping().Result()
	if err != nil {
		log.Fatal(fmt.Sprintf("Could not connected to redis : %s", err.Error()))
	}
	log.Info("Successfully connected to redis")
}

// Get from key
func (r *Redis) Get(key string) (string, error) {
	return r.Client.Get(key).Result()
}

// Set value with key and expire time
func (r *Redis) Set(key string, val string, expire int) error {
	return r.Client.Set(key, val, time.Duration(expire)).Err()
}

// Del delete key in redis
func (r *Redis) Del(key string) error {
	return r.Client.Del(key).Err()
}

// HashGet from key
func (r *Redis) HashGet(hk, key string) (string, error) {
	return r.Client.HGet(hk, key).Result()
}

// HashDel delete key in specify redis's hashtable
func (r *Redis) HashDel(hk, key string) error {
	return r.Client.HDel(hk, key).Err()
}

// Increase
func (r *Redis) Increase(key string) error {
	return r.Client.Incr(key).Err()
}

// Set ttl
func (r *Redis) Expire(key string, dur time.Duration) error {
	return r.Client.Expire(key, dur).Err()
}
