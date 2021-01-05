package cache

import (
	"github.com/go-redis/redis"
	"time"
)

var adapter Adapter

type Adapter interface {
	Connect()
	Get(key string) (string, error)
	Set(key string, val string, expireMs time.Duration) error
	Del(key string) error
	HashGet(hk, key string) (string, error)
	HashDel(hk, key string) error
	Increase(key string) error
	Expire(key string, expireMs time.Duration) error
	SetStruct(key string, obj interface{}, expireMs time.Duration) error
	GetStruct(key string, v interface{}) error
}

func SetUp() {
	adapter = &Redis{}
	adapter.Connect()
}

func GetClient() *redis.Client {
	return adapter.(*Redis).Client
}

// Set val in cache
func Set(key, val string, expireMs time.Duration) error {
	return adapter.Set(key, val, expireMs)
}

// Get val in cache
func Get(key string) (string, error) {
	return adapter.Get(key)
}

// Del delete key in cache
func Del(key string) error {
	return adapter.Del(key)
}

// HashGet get val in hashtable cache
func HashGet(hk, key string) (string, error) {
	return adapter.HashGet(hk, key)
}

// HashDel delete one key:value pair in hashtable cache
func HashDel(hk, key string) error {
	return adapter.HashDel(hk, key)
}

// Increase value
func Increase(key string) error {
	return adapter.Increase(key)
}

func Expire(key string, expireMs time.Duration) error {
	return adapter.Expire(key, expireMs)
}

func SetStruct(key string, obj interface{}, expireMs time.Duration) error {
	return adapter.SetStruct(key, obj, expireMs)
}

func GetStruct(key string, v interface{}) error {
	return adapter.GetStruct(key, v)
}
