package controller

import (
	"github.com/gin-gonic/gin"
	"novel/app/dto"
	"novel/app/utils/response"
)

type BaseController struct {
}

func (bc *BaseController) BindAndValidate(c *gin.Context, obj interface{}) bool {
	if err := dto.Bind(c, obj); err != nil {
		// failValidate(c, err.Error())
		response.ErrorParam(c, err.Error())
		return false
	}
	return true
}

// 不适用go-playground v10作为校验器
func (bc *BaseController) BindAndValidate2(c *gin.Context, obj interface{}) bool {
	if err := dto.BindValidate(c, obj); err != nil {
		response.ErrorParam(c, err.Error())
		return false
	}
	return true
}
