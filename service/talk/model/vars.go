package model

import "github.com/zeromicro/go-zero/core/stores/sqlx"

var ErrNotFound = sqlx.ErrNotFound

type TalkListDetail struct {
	Tid        string `db:"tid"`
	Uid        int64  `db:"uid"`
	Username   string `db:"username"`
	Content    string `db:"content"`
	Imgs       string `db:"imgs"`
	State      int64  `db:"state"`
	CreateTime string `db:"create_time"`
	UpdateTime string `db:"update_time"`
}

type CommentDetail struct {
	Cid        string `db:"cid"`
	ParentId   string `db:"parent_id"`
	Uid        int64  `db:"uid"`
	Username   string `db:"username"`
	Content    string `db:"content"`
	CreateTime string `db:"create_date"`
}
