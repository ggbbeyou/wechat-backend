package common

import (
	"github.com/zeromicro/go-zero/rest/httpx"
	"net/http"
)

/*
* @Author: chuang
* @Date:   2022/9/11 8:36
 */

type Body struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data,omitempty"`
}

func Response(w http.ResponseWriter, resp interface{}, err error) {
	var body Body
	if err != nil {
		switch e := err.(type) {
		case *CodeError:
			body.Code = e.Code
			body.Msg = e.Msg
		default:
			body.Code = FAIL
			body.Msg = FAIL_MSG
		}
	} else {
		body.Code = SUCCESS
		body.Msg = SUCCESS_MSG
		body.Data = resp
	}
	httpx.OkJson(w, body)
}
