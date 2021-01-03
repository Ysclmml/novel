package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"novel/app/utils/response"
)

// 测试控制器

type TestController struct {
	BaseController
}

func (t *TestController) TestGet(context *gin.Context)  {
	fmt.Println(context.Get("test"))
	response.Success(context, "ok", map[string]string{"a": "aa", "b": "bb"})
}