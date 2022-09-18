package friend

import (
	"github.com/zeromicro/go-zero/rest/httpx"
	"go-back/common"
	"net/http"

	"go-back/user/api/internal/logic/friend"
	"go-back/user/api/internal/svc"
	"go-back/user/api/internal/types"
)

func FriendDetailHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.UserId
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := friend.NewFriendDetailLogic(r.Context(), svcCtx)
		resp, err := l.FriendDetail(&req)
		common.Response(w, resp, err)

	}
}
