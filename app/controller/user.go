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
func (uc *UserController) ListReadHistoryByPage(c *gin.Context) {
	var d dto.PageDto
	userDetail := uc.GetUserDetail(c)
	if uc.BindAndValidate(c, &d) {
		list, count := userService.ListReadHistoryByPage(userDetail.Id, d.Page, d.PageSize)
		response.PageSuccess(c, list, count)
	}
}

func (uc *UserController) AddReadHistory(c *gin.Context) {
	var d dto.BookRecordDto
	userDetail := uc.GetUserDetail(c)
	if uc.BindAndValidate2(c, &d) {
		err := userService.AddReadHistory(userDetail.Id, d.BookId, d.PreContentId)
		if err != nil {
			response.Fail(c, consts.CurdCreatFailCode, consts.CurdCreatFailMsg, err.Error())
			return
		}
		response.Ok(c)
	}
}

// 添加反馈
func (uc *UserController) AddFeedBack(c *gin.Context) {
	d := struct {Content string `json:"content" validate:"required"`}{}
	userDetail := uc.GetUserDetail(c)
	if uc.BindAndValidate2(c, &d) {
		if err := userService.AddFeedBack(userDetail.Id, d.Content); err != nil {
			response.Fail(c, consts.CurdCreatFailCode, consts.CurdCreatFailMsg, err.Error())
		} else {
			response.Ok(c)
		}
	}

}

// 分页查询我的反馈列表
func (uc *UserController) ListUserFeedBackByPage(c *gin.Context) {
	var d dto.PageDto
	userDetail := uc.GetUserDetail(c)
	if uc.BindAndValidate(c, &d) {
		page, count := userService.ListUserFeedBackByPage(userDetail.Id, d.Page, d.PageSize)
		response.PageSuccess(c, page, count)
	}
}

func (uc *UserController) UserInfo(c *gin.Context) {
	userDetail := uc.GetUserDetail(c)
	userInfo := userService.UserInfo(userDetail.Id)
	response.Success(c, consts.CurdStatusOkMsg, userInfo)
}

// 更新个人信息
func (uc *UserController) UpdateUserInfo(c *gin.Context) {
	userDetail := uc.GetUserDetail(c)
	var d dto.UserUpdateDto
	if uc.BindAndValidate2(c, &d) {
		if err := userService.UpdateUserInfo(userDetail.Id, d); err != nil{
			response.Fail(c, consts.CurdUpdateFailCode, consts.CurdUpdateFailMsg, err.Error())
		} else {
			// 更新用户的token, todo: 后续需要将当前token加入黑名单
			if d.NickName != "" {
				userDetail.NickName = d.NickName
				token, _ := token_utils.GenerateToken(*userDetail)
				response.Success(c, consts.CurdStatusOkMsg, map[string]string{"token": token})
			}
			response.Success(c, consts.CurdStatusOkMsg, "")
		}
	}
}

func (uc *UserController) ListCommentByPage(c *gin.Context) {
	userDetail := uc.GetUserDetail(c)
	var d dto.PageDto
	if uc.BindAndValidate2(c, &d) {
		page, count := bookService.ListCommentByPage(userDetail.Id, 0, d.Page, d.PageSize)
		response.PageSuccess(c, page, count)
	}
}

func (uc *UserController) BuyBookIndex(c *gin.Context) {
	userDetail := uc.GetUserDetail(c)
	var d dto.BookBuyRecordDto
	if uc.BindAndValidate2(c, &d) {
		d.UserId = userDetail.Id
		if bookIndex, err := bookService.GetIndexDetail(d.BookIndexId); err != nil {
			response.Fail(c, consts.CurdUpdateFailCode, consts.CurdUpdateFailMsg, err.Error())
		} else {
			// 查看书籍id和index里的book_id是否一致
			if bookIndex.BookId != d.BookId {
				response.Fail(c, consts.CurdCreatFailCode, consts.CurdUpdateFailMsg, "书籍id错误")
			}
			d.BuyAmount = bookIndex.BookPrice
			d.BookIndexName = bookIndex.IndexName
			book, _ := bookService.GetBookDetail(d.BookId)
			d.BookName = book.BookName
			err := userService.BuyBookIndex(d)
			if err != nil {
				response.Fail(c, consts.CurdUpdateFailCode, consts.CurdUpdateFailMsg, err.Error())
			} else {
				response.Ok(c)
			}
		}
	}
}

func (uc *UserController) UpdatePassword(c *gin.Context) {
	var d dto.ChgPasswordDto
	userDetail := uc.GetUserDetail(c)
	if uc.BindAndValidate2(c, &d) {
		err := userService.UpdatePassword(userDetail.Id, d.OldPassword, d.NewPassword)
		if err != nil {
			response.Fail(c, consts.CurdUpdateFailCode, consts.CurdUpdateFailMsg, err.Error())
		} else {
			response.Ok(c)
		}
	}
}


