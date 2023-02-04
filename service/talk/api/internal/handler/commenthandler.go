package handler

import (
	"net/http"
	"wechat-backend/common/response"

	"github.com/zeromicro/go-zero/rest/httpx"
	"wechat-backend/service/talk/api/internal/logic"
	"wechat-backend/service/talk/api/internal/svc"
	"wechat-backend/service/talk/api/internal/types"
)

func CommentHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.CommentRequest
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := logic.NewCommentLogic(r.Context(), svcCtx)
		resp, err := l.Comment(&req)
		response.Response(w, resp, err)
	}
}
