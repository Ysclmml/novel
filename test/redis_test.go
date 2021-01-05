package test

import (
	"fmt"
	"novel/app/cache"
	_ "novel/bootstrap"
	"testing"
)

type U struct {
	Name string
	Age int
}

func TestRedis(t *testing.T) {
	// var u = &User{}
	// _ = cache.SetStruct("abc", &U{Name: "hello", Age: 10}, -1)
	// fmt.Println(u)
	// cache.GetStruct("abc", u)
	// fmt.Println(*u)
	res, err := cache.Get("tttttt")
	fmt.Println(res, err == nil)

	var users []User
	users = append(users, User{Name: "hello", Age: 10}, User{Name: "hello22", Age: 120})
	cache.SetStruct("users", users, -1)
	// fmt.Println(users, users == nil)
	var users2 []User
	cache.GetStruct("users2", &users2)
	fmt.Println(users, users2, users2 == nil)
	cache.GetStruct("users", &users2)
	fmt.Println(users, users2, users2 == nil)
}
