package auth

import (
	"github.com/dgrijalva/jwt-go"
	"novel/app/dto"
)

// 自定义jwt的声明字段信息+标准字段，参考地址：https://blog.csdn.net/codeSquare/article/details/99288718
type CustomClaims struct {
	dto.UserDetail
	jwt.StandardClaims
}
