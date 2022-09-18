package common

import "time"

/*
* @Author: chuang
* @Date:   2022/9/16 15:48
 */

// Message 消息格式
type Message struct {
	Content  string        `json:"content"`
	From     int64         `json:"from"`
	To       int64         `json:"to"`
	Type     int64         `json:"type"` //群消息还是用户消息
	SendTime time.Duration `json:"send_time"`
}

//消息类型

const (
	USER_TYPE  = 1
	GROUP_TYPE = 2
)
