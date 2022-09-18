package common

import (
	"encoding/json"
)

/*
* @Author: chuang
* @Date:   2022/9/10 11:57
 */

func Ok() []byte {
	m := make(map[string]interface{})
	m["code"] = SUCCESS
	m["msg"] = SUCCESS_MSG
	buf, _ := json.Marshal(m)
	return buf
}

func Fail(code int, message string) []byte {
	m := make(map[string]interface{})
	m["code"] = code
	m["msg"] = message
	buf, _ := json.Marshal(m)
	return buf
}
