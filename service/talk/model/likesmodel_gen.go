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
	likesFieldNames          = builder.RawFieldNames(&Likes{})
	likesRows                = strings.Join(likesFieldNames, ",")
	likesRowsExpectAutoSet   = strings.Join(stringx.Remove(likesFieldNames, "`id`", "`create_time`", "`update_time`", "`create_at`", "`update_at`"), ",")
	likesRowsWithPlaceHolder = strings.Join(stringx.Remove(likesFieldNames, "`id`", "`create_time`", "`update_time`", "`create_at`", "`update_at`"), "=?,") + "=?"
)

type (
	likesModel interface {
		Insert(ctx context.Context, data *Likes) (sql.Result, error)
		FindOne(ctx context.Context, id int64) (*Likes, error)
		Update(ctx context.Context, data *Likes) error
		Delete(ctx context.Context, id int64) error
	}

	defaultLikesModel struct {
		conn  sqlx.SqlConn
		table string
	}

	Likes struct {
		Id         int64        `db:"id"`
		Lid        string       `db:"lid"`         // 点赞编号
		Tid        string       `db:"tid"`         // 文章编号
		Uid        int64        `db:"uid"`         // 用户编号
		State      int64        `db:"state"`       // 点赞状态，0 取消点赞了，1 点赞了
		CreateTime sql.NullTime `db:"create_time"` // 点赞时间
	}
)

func newLikesModel(conn sqlx.SqlConn) *defaultLikesModel {
	return &defaultLikesModel{
		conn:  conn,
		table: "`likes`",
	}
}

func (m *defaultLikesModel) Delete(ctx context.Context, id int64) error {
	query := fmt.Sprintf("delete from %s where `id` = ?", m.table)
	_, err := m.conn.ExecCtx(ctx, query, id)
	return err
}

func (m *defaultLikesModel) FindOne(ctx context.Context, id int64) (*Likes, error) {
	query := fmt.Sprintf("select %s from %s where `id` = ? limit 1", likesRows, m.table)
	var resp Likes
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

func (m *defaultLikesModel) Insert(ctx context.Context, data *Likes) (sql.Result, error) {
	query := fmt.Sprintf("insert into %s (%s) values (?, ?, ?, ?)", m.table, likesRowsExpectAutoSet)
	ret, err := m.conn.ExecCtx(ctx, query, data.Lid, data.Tid, data.Uid, data.State)
	return ret, err
}

func (m *defaultLikesModel) Update(ctx context.Context, data *Likes) error {
	query := fmt.Sprintf("update %s set %s where `id` = ?", m.table, likesRowsWithPlaceHolder)
	_, err := m.conn.ExecCtx(ctx, query, data.Lid, data.Tid, data.Uid, data.State, data.Id)
	return err
}

func (m *defaultLikesModel) tableName() string {
	return m.table
}
