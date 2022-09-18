package middleware

import (
	"go-back/common"
	"net/http"
)

type AuthMiddleware struct {
}

func NewAuthMiddleware() *AuthMiddleware {
	return &AuthMiddleware{}
}

func (m *AuthMiddleware) Handle(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		//获取token
		authorization := r.Header.Get("Authorization")
		if authorization == "" {
			w.WriteHeader(http.StatusUnauthorized)
			buf := common.Fail(common.UNAUTHORIZED, common.UNAUTHORIZED_MSG)
			w.Write(buf)
			return
		}
		//校验token
		_, err := common.ParseToken(authorization)
		if err != nil {
			w.WriteHeader(http.StatusUnauthorized)
			buf := common.Fail(common.UNAUTHORIZED, common.UNAUTHORIZED_MSG)
			w.Write(buf)
			return
		}
		//r.Header.Set("UserId", string(UserClaims.Uid))
		//r.Header.Set("UserIdentity", UserClaims.Identity)
		//r.Header.Set("UserName", UserClaims.Username)
		next(w, r)
	}
}
