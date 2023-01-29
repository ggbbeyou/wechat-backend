package response

/*
* @Author: chuang
* @Date:   2023/1/10 14:01
 */

type RespError struct {
	Code    int    `json:"code"`
	Message string `json:"config"`
}

func (e *RespError) Error() string {
	return e.Message
}

func NewRespError(code int, msg string) error {
	return &RespError{Code: code, Message: msg}
}

func NewDefaultError(code int, msg string) error {
	return NewRespError(FAIL, SYSTEM_FAIL_MESSAGE)
}
