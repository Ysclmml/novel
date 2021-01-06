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
