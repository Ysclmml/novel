package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gookit/validate"
	"novel/app/dto"
	"novel/app/global/consts"
	"novel/app/service"
	"novel/app/utils/response"
	"time"
)

var newsService = service.NewsService{}

type NewsController struct {
	BaseController
}

// 查询首页新闻, 查询2条新闻信息
func (nc *NewsController) ListIndexNews(c *gin.Context) {
	indexNews := newsService.ListIndexNews()
	response.Success(c, consts.CurdStatusOkMsg, indexNews)
}

// 分页查询新闻列表
func (nc *NewsController) ListByPage(c *gin.Context) {
	var d dto.PageDto
	if nc.BindAndValidate(c, &d) {
		page, count := newsService.ListByPage(d.Page, d.PageSize)
		response.PageSuccess(c, page, count)
	}
}

// 查询新闻详情
func (nc *NewsController) QueryNewsInfo(c *gin.Context) {
	var d dto.GeneralGetDto
	if nc.BindAndValidate(c, &d) {
		if news, err := newsService.QueryNewsInfo(d.Id); err == nil {
			response.Success(c, consts.CurdStatusOkMsg, news)
		} else {
			response.Fail(c, consts.CurdSelectFailCode, consts.CurdSelectFailMsg, err.Error())
		}
	}
}

// UserForm struct
type UserForm struct {
	// Name     string    `validate:"required|minLen:7"`
	UserName     string    `validate:"required" json:"username"`
	Password string    `validate:"required" json:"password"`
	Email    string    `validate:"email" message:"email is invalid"`
	Age      int       `validate:"required|int|min:1|max:99" message:"int:age必须是整数|min: age最小值为1"`
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
		"required":      "oh! the {field} is required",
		"Name.required": "name 必填",
	}
}

// 查询新闻详情
func (nc *NewsController) Test(c *gin.Context) {
	data, _ := validate.FromRequest(c.Request)
	u := UserForm{}
	v := data.Validation()
	v.AddRule("username", "required")
	err := v.BindSafeData(&u)
	if v.Validate() {
		fmt.Println(u, "1231231230-=======", data)
		fmt.Println(v.Errors.Error(), "errrrrrrrrr", err)
		response.Ok(c)
	} else {
		fmt.Println(u, "0-=======", data)
		fmt.Println(v.Errors.Error(), "errrrrrrrrr", err)
		response.Ok(c)
	}

}
