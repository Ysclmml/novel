package test

import (
	"fmt"
	"github.com/gookit/validate"
	_ "novel/bootstrap"
	"testing"
)


// UserForm struct
type UserDto struct {
	// Name     string    `validate:"required|minLen:7"`
	Name     string    `validate:"required"`
	Sex   	*bool	   `validate:"required"`
}


// Messages 您可以自定义验证器错误消息
func (ud UserDto) Messages() map[string]string {
	return validate.MS{
		"required": "oh! the {field} is required",
		"Name.required": "name 必填",
	}
}

func testValidate(ud *UserDto)  {
	// 创建 Validation 实例
	v := validate.Struct(ud)
	if !v.Validate() {
		fmt.Println(v.Errors)
	} else {
		fmt.Println("Success...")
	}
}

func TestMyValidator(t *testing.T)  {
	// 下面如果Sex不传或传nil, 会panic错误
	sex := false
	testValidate(&UserDto{
		Name: "abc",
		Sex: &sex,
	})
	fmt.Println("=============")
	testValidate(&UserDto{
		Name: "abc",
		Sex: nil,
	})
}
