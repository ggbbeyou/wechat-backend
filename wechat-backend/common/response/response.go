package response

import (
	"github.com/zeromicro/go-zero/rest/httpx"
	"net/http"
)

/*
* @Author: chuang
* @Date:   2023/1/11 11:10
 */

type R struct {
	Code    int         `json:"code"`
	Message string      `json:"config"`
	Data    interface{} `json:"data,omitempty"`
}

func Response(w http.ResponseWriter, resp interface{}, err error) {
	r := R{
		Code:    SUCCESS,
		Message: SUCCESS_MESSAGE,
		Data:    resp,
	}
	if err != nil {
		switch e := err.(type) {
		case *RespError:
			r.Code = e.Code
			r.Message = e.Message
		default:
			r.Code = FAIL
			r.Message = SYSTEM_FAIL_MESSAGE
		}
	}
	httpx.OkJson(w, r)
}
