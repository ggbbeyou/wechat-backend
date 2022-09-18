package user

import (
	"github.com/zeromicro/go-zero/rest/httpx"
	"go-back/common"
	"net/http"

	"go-back/user/api/internal/logic/user"
	"go-back/user/api/internal/svc"
	"go-back/user/api/internal/types"
)

func SaveChatListHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.SaveChatListRequest
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := user.NewSaveChatListLogic(r.Context(), svcCtx)
		err := l.SaveChatList(&req)
		common.Response(w, nil, err)

	}
}
