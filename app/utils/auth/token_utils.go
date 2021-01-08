package auth

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/spf13/viper"
	"novel/app/dto"
	"novel/app/global/consts"
	"novel/app/global/my_errors"
	"novel/app/log"
	"time"
)

// 创建 userToken 工厂

func CreateTokenFactory() *userToken {
	return &userToken{
		userJwt: createMyJWT(consts.JwtTokenSignKey),
	}
}

type userToken struct {
	userJwt *JwtSign
}

// 生成token
func (u *userToken) GenerateToken(userDetail dto.UserDetail) (tokens string, err error) {

	// 根据实际业务自定义token需要包含的参数，生成token，注意：用户密码请勿包含在token
	customClaims := CustomClaims{
		UserDetail: userDetail,
		// 特别注意，针对前文的匿名结构体，初始化的时候必须指定键名，并且不带 jwt. 否则报错：Mixture of field: value and value initializers
		StandardClaims: jwt.StandardClaims{
			NotBefore: time.Now().Unix() - 10,       // 生效开始时间
			ExpiresAt: time.Now().Unix() + viper.GetInt64("token.jwtTokenRefreshExpireAt"), // 失效截止时间
		},
	}
	return u.userJwt.createToken(customClaims)
}

// 用户login成功，记录用户token
func (u *userToken) RecordLoginToken(userToken, clientIp string) bool {
	if _, err := u.userJwt.parseToken(userToken); err == nil {
		// userId := customClaims.UserDetail.Id
		// expiresAt := customClaims.ExpiresAt
		return true
	} else {
		return false
	}
}

// 刷新token的有效期（默认+3600秒，参见常量配置项）
func (u *userToken) RefreshToken(oldToken string) (newToken string, res bool) {

	// 解析用户token的数据信息
	_, code := u.isNotExpired(oldToken)
	switch code {
	case consts.JwtTokenOK, consts.JwtTokenExpired:
		// 如果token已经过期，那么执行更新
		if newToken, err := u.userJwt.refreshToken(oldToken, viper.GetInt64("token.jwtTokenRefreshExpireAt")); err == nil {
			if _, err := u.userJwt.parseToken(newToken); err == nil {
				// userId := customClaims.UserDetail.Id
				// expiresAt := customClaims.ExpiresAt
				// if model.CreateUserFactory("").OauthRefreshToken(userId, expiresAt, oldToken, newToken, clientIp) {
				return newToken, true
				// }
			}
		}
	case consts.JwtTokenInvalid:
		log.Fatal(my_errors.ErrorsTokenInvalid)
	}

	return "", false
}

// 销毁token，基本用不到，因为一个网站的用户退出都是直接关闭浏览器窗口，极少有户会点击“注销、退出”等按钮，销毁token其实无多大意义
func (u *userToken) DestroyToken() {

}

// 判断token是否未过期
func (u *userToken) isNotExpired(token string) (*CustomClaims, int) {
	if customClaims, err := u.userJwt.parseToken(token); err == nil {

		if time.Now().Unix()-customClaims.ExpiresAt < 0 {
			// token有效
			return customClaims, consts.JwtTokenOK
		} else {
			// 过期的token
			return customClaims, consts.JwtTokenExpired
		}
	} else {
		// 无效的token
		return nil, consts.JwtTokenInvalid
	}
}

// 判断token是否有效（未过期+数据库用户信息正常）
func (u *userToken) IsEffective(token string) bool {
	_, code := u.isNotExpired(token)
	if consts.JwtTokenOK == code {
		// if user_item := Model.CreateUserFactory("").ShowOneItem(customClaims.UserId); user_item != nil {
		// if model.CreateUserFactory("").OauthCheckTokenIsOk(customClaims.UserId, token) {
		return true
		// }
	}
	return false
}
