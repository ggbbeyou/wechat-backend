package search

import (
	"net/http"
	"wechat-backend/common/response"
	"wechat-backend/service/search/api/internal/logic/search"
	"wechat-backend/service/search/api/internal/svc"
	"wechat-backend/service/search/api/internal/types"

	"github.com/zeromicro/go-zero/rest/httpx"
)

func SearchHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.SearchRequest
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := search.NewSearchLogic(r.Context(), svcCtx)
		resp, err := l.Search(&req)
		response.Response(w, resp, err)
	}
}
