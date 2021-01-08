package test

import (
	"fmt"
	"github.com/go-playground/validator/v10"
	"github.com/gookit/validate"
	"novel/app/dto"
	_ "novel/bootstrap"
	"testing"
	"time"
)

func TestValidator(t *testing.T) {
	validate := validator.New()

	d := dto.BookRankDto{Limit: 30, Type: 1}
	err := validate.Struct(d)
	fmt.Println(err)

	err = validate.Struct(&dto.BookRankDto{Limit: 66, Type: -1})
	fmt.Println(err)
}


// UserForm struct
type UserForm struct {
	// Name     string    `validate:"required|minLen:7"`
	Name     string    `validate:"required"`
	Email    string    `validate:"email" message:"email is invalid"`
	Age      int       `validate:"required|int|min:1|max:99" message:"int:age must int| min: age min value is 1"`
	CreateAt int       `validate:"min:1"`
	Safe     int       `validate:"-"`
	UpdateAt time.Time `validate:"required"`
	Code     string    `validate:"customValidator"` // 使用自定义验证器
}

// CustomValidator 定义在结构体中的自定义验证器
func (f UserForm) CustomValidator(val string) bool {
	return len(val) == 4
}

// Messages 您可以自定义验证器错误消息
func (f UserForm) Messages() map[string]string {
	return validate.MS{
		"required": "oh! the {field} is required",
		"Name.required": "name 必填",
	}
}

// Translates 你可以自定义字段翻译
func (f UserForm) Translates() map[string]string {
	return validate.MS{
		"Name": "用户名称",
		"Email": "用户Email",
	}
}

func TestValidator2(t *testing.T)  {
	u := &UserForm{
		// Name: "inhere",
		Name: "",
	}

	// 创建 Validation 实例
	v := validate.Struct(u)
	// v := validate.New(u)

	if v.Validate() { // 验证成功
		// do something ...
	} else {
		fmt.Println(v.Errors) // 所有的错误消息
		fmt.Println(v.Errors.One()) // 返回随机一条错误消息
		fmt.Println(v.Errors.Field("Name")) // 返回该字段的错误消息
	}
}
