package test

import (
	"fmt"
	"github.com/jinzhu/copier"
	"testing"
)

type User struct {
	Name string
	Age int
}

type Account struct {
	Name *string
	Age int
	Address string
}

func TestCopy(t *testing.T) {
	user := User{
		Name: "aaa",
		Age:  12,
	}
	account := Account{}
	// 支持指针飞赋值
	_ = copier.Copy(&account, &user)
	fmt.Println("user", user.Name)
	fmt.Println("account", *account.Name)
}
