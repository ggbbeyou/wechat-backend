package record

import (
	"net/http"
	"wechat-backend/common/response"

	"github.com/zeromicro/go-zero/rest/httpx"
	"wechat-backend/service/communicate/api/internal/logic/record"
	"wechat-backend/service/communicate/api/internal/svc"
	"wechat-backend/service/communicate/api/internal/types"
)

func RecordListHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.RecordListRequest
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := record.NewRecordListLogic(r.Context(), svcCtx)
		resp, err := l.RecordList(&req)

		response.Response(w, resp, err)
	}
}
