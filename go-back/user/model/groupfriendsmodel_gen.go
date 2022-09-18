// Code generated by goctl. DO NOT EDIT!

package model

import (
	"context"
	"database/sql"
	"fmt"
	"strings"

	"github.com/zeromicro/go-zero/core/stores/builder"
	"github.com/zeromicro/go-zero/core/stores/sqlc"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"github.com/zeromicro/go-zero/core/stringx"
)

var (
	groupfriendsFieldNames          = builder.RawFieldNames(&Groupfriends{})
	groupfriendsRows                = strings.Join(groupfriendsFieldNames, ",")
	groupfriendsRowsExpectAutoSet   = strings.Join(stringx.Remove(groupfriendsFieldNames, "`id`", "`create_time`", "`update_time`", "`create_at`", "`update_at`"), ",")
	groupfriendsRowsWithPlaceHolder = strings.Join(stringx.Remove(groupfriendsFieldNames, "`id`", "`create_time`", "`update_time`", "`create_at`", "`update_at`"), "=?,") + "=?"
)

type (
	groupfriendsModel interface {
		Insert(ctx context.Context, data *Groupfriends) (sql.Result, error)
		FindOne(ctx context.Context, id int64) (*Groupfriends, error)
		Update(ctx context.Context, data *Groupfriends) error
		Delete(ctx context.Context, id int64) error

		// QueryGroupIdsByUid rpc方法调用
		QueryGroupIdsByUid(ctx context.Context, id int64) (GidArr, error)
	}

	defaultGroupfriendsModel struct {
		conn  sqlx.SqlConn
		table string
	}

	Groupfriends struct {
		Id         int64        `db:"id"`          // id
		Uid        int64        `db:"uid"`         // 用户id
		Gid        int64        `db:"gid"`         // 群聊id
		Indentity  uint64       `db:"indentity"`   // 0 群主 1 成员
		CreateTime sql.NullTime `db:"create_time"` // 加入群聊时间
		ExitTime   sql.NullTime `db:"exit_time"`   // 退出群聊时间
		IsDelete   uint64       `db:"is_delete"`   // 逻辑删除 用户退出群聊时间
	}
)

func newGroupfriendsModel(conn sqlx.SqlConn) *defaultGroupfriendsModel {
	return &defaultGroupfriendsModel{
		conn:  conn,
		table: "`groupfriends`",
	}
}

func (m *defaultGroupfriendsModel) Delete(ctx context.Context, id int64) error {
	query := fmt.Sprintf("delete from %s where `id` = ?", m.table)
	_, err := m.conn.ExecCtx(ctx, query, id)
	return err
}

func (m *defaultGroupfriendsModel) FindOne(ctx context.Context, id int64) (*Groupfriends, error) {
	query := fmt.Sprintf("select %s from %s where `id` = ? limit 1", groupfriendsRows, m.table)
	var resp Groupfriends
	err := m.conn.QueryRowCtx(ctx, &resp, query, id)
	switch err {
	case nil:
		return &resp, nil
	case sqlc.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}

func (m *defaultGroupfriendsModel) QueryGroupIdsByUid(ctx context.Context, id int64) (GidArr, error) {
	query := fmt.Sprintf("select `gid` from %s where `uid` = ?", m.table)
	var resp GidArr
	err := m.conn.QueryRowsCtx(ctx, &resp, query, id)
	if err != nil {
		return nil, err
	}
	switch err {
	case nil:
		return resp, nil
	case sqlc.ErrNotFound:
		return nil, nil
	default:
		return nil, err
	}
}

func (m *defaultGroupfriendsModel) Insert(ctx context.Context, data *Groupfriends) (sql.Result, error) {
	query := fmt.Sprintf("insert into %s (%s) values (?, ?, ?, ?, ?)", m.table, groupfriendsRowsExpectAutoSet)
	ret, err := m.conn.ExecCtx(ctx, query, data.Uid, data.Gid, data.Indentity, data.ExitTime, data.IsDelete)
	return ret, err
}

func (m *defaultGroupfriendsModel) Update(ctx context.Context, data *Groupfriends) error {
	query := fmt.Sprintf("update %s set %s where `id` = ?", m.table, groupfriendsRowsWithPlaceHolder)
	_, err := m.conn.ExecCtx(ctx, query, data.Uid, data.Gid, data.Indentity, data.ExitTime, data.IsDelete, data.Id)
	return err
}

func (m *defaultGroupfriendsModel) tableName() string {
	return m.table
}