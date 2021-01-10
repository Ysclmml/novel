package token_utils

import (
	"errors"
	"github.com/dgrijalva/jwt-go"
	"time"
)

// --------------------  JWT   ----------------准备阶段  ↓
// 常量类型不支持  error 所以该部分错误常量没有全局统一定义
var (
	tokenNotValidYet = errors.New("token not active yet")
	tokenMalformed   = errors.New("that's not even a token")
	tokenInvalid     = errors.New("couldn't handle this token")
	tokenExpire    = errors.New("token has expired")
	signKey          = "novel"
)

// 获取signKey
func GetSignKey() string {
	return signKey
}

// 设置signKey（类似秘钥）
func SetSignKey(key string) string {
	signKey = key
	return signKey
}

// --------------------  JWT   ----------------正式阶段  ↓

// 使用工厂创建一个 JWT 结构体
func createMyJWT(signKey string) *JwtSign {
	if len(signKey) > 0 {
		SetSignKey(signKey)
	}
	return &JwtSign{
		[]byte(GetSignKey()),
	}
}

// 定义一个 JWT验签 结构体
type JwtSign struct {
	SigningKey []byte
}

// CreateToken 生成一个token
func (j *JwtSign) createToken(claims CustomClaims) (string, error) {
	// 生成jwt格式的header、claims 部分
	tokenPartA := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	// 继续添加秘钥值，生成最后一部分
	return tokenPartA.SignedString(j.SigningKey)
}

// 解析Token
func (j *JwtSign) parseToken(tokenString string) (*CustomClaims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &CustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		return j.SigningKey, nil
	})
	if token == nil {
		return nil, tokenInvalid
	}
	if err != nil {
		if ve, ok := err.(*jwt.ValidationError); ok {
			if ve.Errors&jwt.ValidationErrorMalformed != 0 {
				return nil, tokenMalformed
			} else if ve.Errors&jwt.ValidationErrorNotValidYet != 0 {
				return nil, tokenNotValidYet
			} else if ve.Errors&jwt.ValidationErrorExpired != 0 {
				// 如果 TokenExpired ,只是过期, 单独处理
				return token.Claims.(*CustomClaims), tokenExpire
			} else {
				return nil, tokenInvalid
			}
		}
	}
	claims := token.Claims.(*CustomClaims)
	return claims, nil
}

// 更新token
func (j *JwtSign) refreshToken(tokenString string, extraAddSeconds int64) (string, error) {

	if CustomClaims, err := j.parseToken(tokenString); err == nil {
		CustomClaims.ExpiresAt = time.Now().Unix() + extraAddSeconds
		return j.createToken(*CustomClaims)
	} else {
		return "", err
	}
}
