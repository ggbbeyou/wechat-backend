package model

import "github.com/zeromicro/go-zero/core/stores/sqlx"

var ErrNotFound = sqlx.ErrNotFound

type UserDetail struct {
	Username    string `db:"username"`
	Nickname    string `db:"nickname"`
	Age         int64  `db:"age"`
	Avatar      string `db:"avatar"`
	Email       string `db:"email"`
	Introduce   string `db:"introduce"`
	LastTime    string `db:"last_time"`
	LastAddress string `db:"last_address"`
}

type GroupDetail struct {
	Gid        int64  `json:"gid"`
	Uid        int64  `json:"uid"`
	Avatar     string `json:"avatar"`
	Gname      string `json:"name"`
	Count      string `json:"count"`
	CreateTime string `json:"create_time"`
}
