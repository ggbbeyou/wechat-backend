package common

import (
	"errors"
	"github.com/golang-jwt/jwt/v4"
	"time"
)

// UserClaims 信息体
type UserClaims struct {
	Uid      int64
	Username string
	jwt.RegisteredClaims
}

// SignKey 加密密钥
var SignKey = []byte("1234567")

// CreateToken 生成token
func CreateToken(uid int64, name string) (string, error) {
	claim := UserClaims{
		Uid:      uid,
		Username: name,
		RegisteredClaims: jwt.RegisteredClaims{
			NotBefore: jwt.NewNumericDate(time.Now()),                    // 生效时间
			IssuedAt:  jwt.NewNumericDate(time.Now()),                    // 签发时间
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * 2)), //过期时间
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claim)
	tokenStr, err := token.SignedString(SignKey)
	if err != nil {
		return "", err
	}
	return tokenStr, nil
}

// ParseToken 解密
func ParseToken(token string) (*UserClaims, error) {
	t, err := jwt.ParseWithClaims(token, &UserClaims{}, func(token *jwt.Token) (interface{}, error) {
		return SignKey, nil
	})
	if err != nil {
		if ve, ok := err.(*jwt.ValidationError); ok {
			if ve.Errors&jwt.ValidationErrorMalformed != 0 {
				return nil, errors.New("that's not even a token")
			} else if ve.Errors&jwt.ValidationErrorExpired != 0 {
				return nil, errors.New("token is expired")
			} else if ve.Errors&jwt.ValidationErrorNotValidYet != 0 {
				return nil, errors.New("token not active yet")
			} else {
				return nil, errors.New("couldn't handle this token")
			}
		}
	}
	if claims, ok := t.Claims.(*UserClaims); ok && t.Valid {
		return claims, nil
	}
	//类型断言
	return nil, errors.New("couldn't handle this token")
}
