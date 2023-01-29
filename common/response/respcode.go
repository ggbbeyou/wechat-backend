package response

/*
* @Author: chuang
* @Date:   2023/1/10 14:01
 */
const (
	UNAUTHORIZED    = -105
	SUCCESS         = 200
	FAIL            = 400
	LOGIN_FAIL      = -106
	JWT_FAIL        = -107
	USER_NOT_EXIST  = -108
	GROUP_NOT_EXIST = -109
	WEBCOKET_ERROR  = -120
)
const (
	UNAUTHORIZED_MESSAGE    = "请登陆后再访问"
	SUCCESS_MESSAGE         = "操作成功"
	SYSTEM_FAIL_MESSAGE     = "系统异常"
	LOGIN_FAIL_MESSAGE      = "用户名或者密码错误"
	JWT_FAIL_MESSAGE        = "tokan生成失败"
	USER_NOT_EXIST_MESSAGE  = "用户不存在"
	GROUP_NOT_EXIST_MESSAGE = "群聊不存在"
	WEBCOKET_ERROR_MESSAGE  = "websocket处理异常"
)
