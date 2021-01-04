package response

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"novel/app/global/consts"
	"novel/app/global/my_errors"
)

func ReturnJson(Context *gin.Context, httpCode int, dataCode int, msg string, data interface{}) {

	//Context.Header("key2020","value2020")  	//可以根据实际情况在头部添加额外的其他信息
	Context.JSON(httpCode, gin.H{
		"code": dataCode,
		"msg":  msg,
		"data": data,
	})
}

func ReturnErrJson(Context *gin.Context, httpCode int, errCode int, msg string, detail interface{}) {

	//Context.Header("key2020","value2020")  	//可以根据实际情况在头部添加额外的其他信息
	Context.JSON(httpCode, gin.H{
		"code": errCode,
		"msg":  msg,
		"detail": detail,
	})
}

// 将json字符窜以标准json格式返回（例如，从redis读取json、格式的字符串，返回给浏览器json格式）
func ReturnJsonFromString(Context *gin.Context, httpCode int, jsonStr string) {
	Context.Header("Content-Type", "application/json; charset=utf-8")
	Context.String(httpCode, jsonStr)
}

// 语法糖函数封装

// 直接返回成功
func Success(c *gin.Context, msg string, data interface{}) {
	ReturnJson(c, http.StatusOK, consts.CurdStatusOkCode, msg, data)
}

// 直接返回不包含数据的成功
func Ok(c *gin.Context)  {
	c.JSON(http.StatusOK, gin.H{
		"code": consts.CurdStatusOkCode,
		"msg":  "ok",
	})
}

// 失败的业务逻辑
func Fail(c *gin.Context, dataCode int, msg string, detail interface{}) {
	ReturnErrJson(c, http.StatusBadRequest, dataCode, msg, detail)
	c.Abort()
}

//权限校验失败
func ErrorAuthFail(c *gin.Context) {
	ReturnErrJson(c, http.StatusUnauthorized, http.StatusUnauthorized, my_errors.ErrorsNoAuthorization, "")
	//暂停执行
	c.Abort()
}

//参数校验错误
func ErrorParam(c *gin.Context, wrongParam interface{}) {
	ReturnErrJson(c, http.StatusBadRequest, consts.ValidatorParamsCheckFailCode, consts.ValidatorParamsCheckFailMsg, wrongParam)
	c.Abort()
}

// 系统执行代码错误
func ErrorSystem(c *gin.Context, msg string, detail interface{}) {
	ReturnErrJson(c, http.StatusInternalServerError, consts.ServerOccurredErrorCode, consts.ServerOccurredErrorMsg+msg, detail)
	c.Abort()
}
