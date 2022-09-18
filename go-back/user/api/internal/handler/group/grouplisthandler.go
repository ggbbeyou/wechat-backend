package group

import (
	"github.com/zeromicro/go-zero/rest/httpx"
	"go-back/common"
	"net/http"

	"go-back/user/api/internal/logic/group"
	"go-back/user/api/internal/svc"
	"go-back/user/api/internal/types"
)

func GroupListHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.UserId
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := group.NewGroupListLogic(r.Context(), svcCtx)
		resp, err := l.GroupList(&req)
		common.Response(w, resp, err)

	}
}
