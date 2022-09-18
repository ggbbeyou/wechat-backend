package common

/*
* @Author: chuang
* @Date:   2022/9/10 11:59
 */

const (
	UNAUTHORIZED = -105
	SUCCESS      = 200
	FAIL         = 400
	LOGIN_FAIL   = -106
	JWT_FAIL     = -107
)
const (
	UNAUTHORIZED_MSG = "请登陆后再访问"
	SUCCESS_MSG      = "操作成功"
	FAIL_MSG         = "系统异常"
	LOGIN_FAIL_MSG   = "用户名或者密码错误"
	JWT_FAIL_MSG     = "tokan生成失败"
)
