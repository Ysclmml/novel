package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"novel/app/dto"
	"novel/app/utils/response"
)

// 测试控制器

type TestController struct {
	BaseController
}

func (t *TestController) TestGet(c *gin.Context)  {
	var testDto dto.TestDto
	if t.BindAndValidate(c, &testDto) {
		fmt.Println(testDto, testDto.Name, testDto.Id)
		response.Success(c, "ok", testDto)
	}
	//response.Success(c, "ok", map[string]string{"a": "aa", "b": "bb"})
}