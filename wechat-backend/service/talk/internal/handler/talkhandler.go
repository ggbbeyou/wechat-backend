package handler

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"wechat-backend/service/talk/internal/logic"
	"wechat-backend/service/talk/internal/svc"
	"wechat-backend/service/talk/internal/types"
)

func TalkHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.Request
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := logic.NewTalkLogic(r.Context(), svcCtx)
		resp, err := l.Talk(&req)
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
