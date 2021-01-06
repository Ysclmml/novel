package test

import (
	"fmt"
	"github.com/go-playground/validator/v10"
	"novel/app/dto"
	_ "novel/bootstrap"
	"testing"
)

func TestValidator(t *testing.T) {
	validate := validator.New()

	d := dto.BookRankDto{Limit: 30, Type: 1}
	err := validate.Struct(d)
	fmt.Println(err)

	err = validate.Struct(&dto.BookRankDto{Limit: 66, Type: -1})
	fmt.Println(err)
}
