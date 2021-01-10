package middleware

import (
	"github.com/gin-gonic/gin"
	"novel/app/utils/response"
	"novel/app/utils/token_utils"
	"strings"
)

type HeaderParams struct {
	Authorization string `header:"Authorization"`
}

// 检查是否登录的中间件
func CheckLogin(context *gin.Context) {
	//  模拟验证token
	headerParams := HeaderParams{}

	//  推荐使用 ShouldBindHeader 方式获取头参数
	if err := context.ShouldBindHeader(&headerParams); err != nil {
		context.Abort()
	}

	if len(headerParams.Authorization) >= 20 {
		token := strings.Split(headerParams.Authorization, " ")
		if len(token) == 2 && len(token[1]) >= 20 {
			// 前缀必须是novel
			if token[0] == "novel" && token_utils.IsEffective(token[1]) {
				// 在这里把UserDetails信息设置到context中...
				userDetail, _ := token_utils.GetUserDetailFromToken(token[1])
				context.Set("userDetail", userDetail)
				context.Set("token", token[1])

				context.Next()
			} else {
				response.ErrorAuthFail(context)
			}
		}
	} else {
		response.ErrorAuthFail(context)
	}

}
