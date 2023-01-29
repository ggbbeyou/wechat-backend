package user

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"wechat-backend/service/user/api/internal/logic/user"
	"wechat-backend/service/user/api/internal/svc"
	"wechat-backend/service/user/api/internal/types"
)

func FriendListHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.UidRequest
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := user.NewFriendListLogic(r.Context(), svcCtx)
		resp, err := l.FriendList(&req)
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}