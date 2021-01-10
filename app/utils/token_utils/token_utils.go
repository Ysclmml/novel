package token_utils

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/spf13/viper"
	"novel/app/dto"
	"novel/app/global/consts"
	"time"
)

var (
	userJwt = createMyJWT(consts.JwtTokenSignKey)
)

// 生成token
func GenerateToken(userDetail dto.UserDetail) (tokens string, err error) {

	// 根据实际业务自定义token需要包含的参数，生成token，注意：用户密码请勿包含在token
	expireTime := viper.GetInt64("token.expirationTime")
	customClaims := CustomClaims{
		UserDetail: userDetail,
		// 特别注意，针对前文的匿名结构体，初始化的时候必须指定键名，并且不带 jwt. 否则报错：Mixture of field: value and value initializers
		StandardClaims: jwt.StandardClaims{
			NotBefore: time.Now().Unix() - 10,       // 生效开始时间
			ExpiresAt: time.Now().Unix() + expireTime, // 失效截止时间
		},
	}
	return userJwt.createToken(customClaims)
}

// 用户login成功，记录用户token
func RecordLoginToken(userToken, clientIp string) bool {
	if _, err := userJwt.parseToken(userToken); err == nil {
		// userId := customClaims.UserDetail.Id
		// expiresAt := customClaims.ExpiresAt
		return true
	} else {
		return false
	}
}

// 从token中获取JWT中的负载
func GetClaimFromToken(token string) (*CustomClaims, error){
	return userJwt.parseToken(token)
}

// 直接获取UserDetail信息
func GetUserDetailFromToken(token string) (*dto.UserDetail, error) {
	claim, err := GetClaimFromToken(token)
	if err != nil {
		return nil, err
	}
	return &claim.UserDetail, nil
}

// 刷新token的有效期（默认+3600秒，参见常量配置项）
func RefreshToken(oldToken string) (string, error) {
	// 解析用户token的数据信息
	expireTime := viper.GetInt64("token.expirationTime")
	return userJwt.refreshToken(oldToken, expireTime)
}

// 销毁token，基本用不到，因为一个网站的用户退出都是直接关闭浏览器窗口，极少有户会点击“注销、退出”等按钮，销毁token其实无多大意义
func DestroyToken() {

}

// 判断token是否有效（未过期+数据库用户信息正常）
func IsEffective(token string) bool {
	_, err := userJwt.parseToken(token)
	return err == nil
}
