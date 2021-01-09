package controller

import (
	"github.com/gin-gonic/gin"
	"novel/app/dto"
	"novel/app/global/consts"
	"novel/app/service"
	"novel/app/utils/auth"
	"novel/app/utils/response"
)

var userService = service.UserService{}

type UserController struct {
	BaseController
}
func (uc *UserController) Login(c *gin.Context) {
	var d dto.LoginDto
	if uc.BindAndValidate2(c, &d) {
		userDetail, err := userService.Login(d.UserName, d.Password)
		if err != nil {
			response.Fail(c, consts.CurdLoginFailCode, consts.CurdLoginFailMsg, err.Error())
			return
		}
		// 设置token
		token, _ := auth.CreateTokenFactory().GenerateToken(*userDetail)
		// 返回token
		data := map[string]string{
			"token": token,
		}
		response.Success(c, consts.CurdStatusOkMsg, data)
	}
}

func (uc *UserController) Register(c *gin.Context) {
	var d dto.RegisterDto
	if uc.BindAndValidate2(c, &d) {
		if err := userService.Register(d); err == nil {
			response.Success(c, "注册成功", nil)
			return
		} else {
			response.Fail(c, consts.CurdCreatFailCode, consts.CurdCreatFailMsg, err.Error())
		}
	}
}

func (uc *UserController) RefreshToken(c *gin.Context) {
	
}

func (uc *UserController) QueryIsInShelf(c *gin.Context) {
	
}

func (uc *UserController) AddToBookShelf(c *gin.Context) {
	
}

func (uc *UserController) RemoveFromBookShelf(c *gin.Context) {
	
}

func (uc *UserController) ListBookShelfByPage(c *gin.Context) {
	
}

func (uc *UserController) AddReadHistory(c *gin.Context) {
	
}

func (uc *UserController) AddFeedBack(c *gin.Context) {
	
}

func (uc *UserController) ListUserFeedBackByPage(c *gin.Context) {
	
}

func (uc *UserController) UserInfo(c *gin.Context) {
	
}

func (uc *UserController) UpdateUserInfo(c *gin.Context) {
	
}

func (uc *UserController) ListCommentByPage(c *gin.Context) {
	
}

func (uc *UserController) BuyBookIndex(c *gin.Context) {
	
}


