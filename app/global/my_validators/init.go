package my_validators

import (
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
	"github.com/gookit/validate"
	"log"
	"regexp"
)

func SetUp() {
	// Register custom validate methods
	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		_ = v.RegisterValidation("pwdValidate", PasswordValidate)
	} else {
		log.Fatal("Gin fail to registered custom validator(v10)")
	}
	// 添加全局校验器
	validate.AddValidator("phone", Phone)
	// 添加全局的错误消息
	validate.AddGlobalMessages(map[string]string{
		"phone": "{field}必须是一个合法的手机号",
	})
}

func PasswordValidate(fl validator.FieldLevel) bool {
	val := fl.Field().String()
	reg := regexp.MustCompile(`^1[3-9][0-9]{9}$`)
	return reg.Match([]byte(val))
}

func Phone(val string) bool {
	reg := regexp.MustCompile(`^1[3-9][0-9]{9}$`)
	return reg.Match([]byte(val))
}
