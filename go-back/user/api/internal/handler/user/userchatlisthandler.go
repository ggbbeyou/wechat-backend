package user

import (
	"github.com/zeromicro/go-zero/rest/httpx"
	utils "go-back/common"
	"net/http"

	"go-back/user/api/internal/logic/user"
	"go-back/user/api/internal/svc"
	"go-back/user/api/internal/types"
)

func UserChatListHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.UserId
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := user.NewUserChatListLogic(r.Context(), svcCtx)
		resp, err := l.UserChatList(&req)
		utils.Response(w, resp, err)
	}
}