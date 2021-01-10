package controller

import (
	"github.com/gin-gonic/gin"
	"novel/app/dto"
	"novel/app/global/consts"
	"novel/app/service"
	"novel/app/utils/response"
	"novel/app/utils/token_utils"
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
		token, _ := token_utils.GenerateToken(*userDetail)
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
	// 中间件李设置了token和userDetail信息
	token := c.GetString("token")
	refreshToken, err := token_utils.RefreshToken(token)
	if err != nil {
		response.Fail(c, consts.JwtTokenInvalid, consts.JwtTokenFormatErrMsg, err.Error())
		return
	}
	// 返回token
	data := map[string]string{
		"refresh_token": refreshToken,
	}
	response.Success(c, consts.CurdStatusOkMsg, data)

}

// 查询小说是否已加入书架
func (uc *UserController) QueryIsInShelf(c *gin.Context) {
	userDetail := uc.GetUserDetail(c)
	var d dto.GeneralGetDto
	if uc.BindAndValidate(c, &d) {
		status := userService.QueryIsInShelf(userDetail.Id, d.Id)
		response.Success(c, "ok", map[string]bool{"status": status})
	}
}

func (uc *UserController) AddToBookShelf(c *gin.Context) {
	userDetail := uc.GetUserDetail(c)
	var d dto.AddBookShelfDto
	if uc.BindAndValidate2(c, &d) {
		err := userService.AddToBookShelf(userDetail.Id, d.BookId, d.PreContentId)
		if err != nil {
			response.Fail(c, consts.CurdCreatFailCode, consts.CurdCreatFailMsg, err.Error())
			return
		}
		response.Ok(c)
	}
}

func (uc *UserController) RemoveFromBookShelf(c *gin.Context) {
	var d dto.GeneralPostDto
	userDetail := uc.GetUserDetail(c)
	if uc.BindAndValidate(c, &d) {
		err := userService.RemoveFromBookShelf(userDetail.Id, d.Id)
		if err != nil {
			response.Fail(c, consts.CurdDeleteFailCode, consts.CurdDeleteFailMsg, err.Error())
			return
		}
		response.Ok(c)
	}
}

func (uc *UserController) ListBookShelfByPage(c *gin.Context) {
	var d dto.PageDto
	userDetail := uc.GetUserDetail(c)
	if uc.BindAndValidate(c, &d) {
		list, count := userService.ListBookShelfByPage(userDetail.Id, d.Page, d.PageSize)
		response.PageSuccess(c, list, count)
	}
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


