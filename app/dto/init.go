package dto

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/gookit/validate"
	"github.com/pkg/errors"
	"novel/app/global/my_validators/binding"
	"strings"
)

// Bind : bind request dto and auto verify parameters
func Bind(c *gin.Context, obj interface{}) error {
	_ = c.ShouldBindUri(obj)
	if err := c.ShouldBind(obj); err != nil {
		if fieldErr, ok := err.(validator.ValidationErrors); ok {
			var tagErrorMsg []string
			for _, v := range fieldErr {
				if _, has := ValidateErrorMessage[v.Tag()]; has {
					tagErrorMsg = append(tagErrorMsg, fmt.Sprintf(ValidateErrorMessage[v.Tag()], v.Field(), v.Value()))
				} else {
					tagErrorMsg = append(tagErrorMsg, err.Error())
				}
			}
			return errors.New(strings.Join(tagErrorMsg, ","))
		} else if jsonErr, ok := err.(*json.UnmarshalTypeError); ok {
			return errors.New(jsonErr.Field + "类型错误")
		} else {
			// 其他错误
			fmt.Println("err", err)
		}
	}
	return nil
}

func BindValidate(c *gin.Context, obj interface{}) error {
	_ = binding.ShouldBindUri(c, obj)
	if err := binding.ShouldBind(c, obj); err != nil {
		if jsonErr, ok := err.(*json.UnmarshalTypeError); ok {
			return errors.New(jsonErr.Field + "类型错误")
		} else {
			// 其他错误
			fmt.Println("err", err)
		}
	}
	// 将数据绑定到结构体之后, 用validator库来进行校验
	v := validate.Struct(obj)
	if !v.Validate() {
		return v.Errors
	}
	return nil
}

// ValidateErrorMessage : customize error messages
var ValidateErrorMessage = map[string]string{
	"customValidate": "%s can not be %s",
	"required":       "%s is required,got empty %#v",
	"pwdValidate":    "%s is not a valid password",
	"permsValidate":  "perms [%s] is not allow to contains comma",
}
