package utils

import (
	"github.com/golang-jwt/jwt/v4"
	"time"
	"wechat-backend/common/response"
)

/*
* @Author: chuang
* @Date:   2023/1/11 13:20
 */

// UserClaims 信息体
type UserClaims struct {
	Username string
	Nickname string
	jwt.RegisteredClaims
}

// GenerateToken 生成token
func GenerateToken(nickname, Username, accessSecret string, expire int64) (string, error) {
	claim := UserClaims{
		Nickname: nickname,
		Username: Username,
		RegisteredClaims: jwt.RegisteredClaims{
			NotBefore: jwt.NewNumericDate(time.Now()),                                          // 生效时间
			IssuedAt:  jwt.NewNumericDate(time.Now()),                                          // 签发时间
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Duration(expire) * time.Second)), //过期时间
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claim)
	tokenStr, err := token.SignedString([]byte(accessSecret))
	if err != nil {
		return "", response.NewRespError(response.JWT_FAIL, err.Error())
	}
	return tokenStr, nil
}

// ParseToken 解密
func ParseToken(token, accessSecret string) (*UserClaims, error) {
	t, err := jwt.ParseWithClaims(token, &UserClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(accessSecret), nil
	})
	if err != nil {
		if ve, ok := err.(*jwt.ValidationError); ok {
			if ve.Errors&jwt.ValidationErrorMalformed != 0 {
				return nil, response.NewRespError(response.JWT_FAIL, "that's not even a token")
			} else if ve.Errors&jwt.ValidationErrorExpired != 0 {
				return nil, response.NewRespError(response.JWT_FAIL, "token is expired")
			} else if ve.Errors&jwt.ValidationErrorNotValidYet != 0 {
				return nil, response.NewRespError(response.JWT_FAIL, "token not active yet")
			} else {
				return nil, response.NewRespError(response.JWT_FAIL, "couldn't handle this token")
			}
		}
	}
	if claims, ok := t.Claims.(*UserClaims); ok && t.Valid {
		return claims, nil
	}
	//类型断言
	return nil, response.NewRespError(response.JWT_FAIL, "couldn't handle this token")
}
