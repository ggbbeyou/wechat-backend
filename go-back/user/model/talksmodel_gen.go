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
	talksFieldNames          = builder.RawFieldNames(&Talks{})
	talksRows                = strings.Join(talksFieldNames, ",")
	talksRowsExpectAutoSet   = strings.Join(stringx.Remove(talksFieldNames, "`id`", "`create_time`", "`update_time`", "`create_at`", "`update_at`"), ",")
	talksRowsWithPlaceHolder = strings.Join(stringx.Remove(talksFieldNames, "`id`", "`create_time`", "`update_time`", "`create_at`", "`update_at`"), "=?,") + "=?"
)

type (
	talksModel interface {
		Insert(ctx context.Context, data *Talks) (sql.Result, error)
		FindOne(ctx context.Context, id int64) (*Talks, error)
		Update(ctx context.Context, data *Talks) error
		Delete(ctx context.Context, id int64) error
	}

	defaultTalksModel struct {
		conn  sqlx.SqlConn
		table string
	}

	Talks struct {
		Id         int64          `db:"id"`          // id
		Tid        int64          `db:"tid"`         // 说说id uuid
		Uid        int64          `db:"uid"`         // uid用户id
		Content    string         `db:"content"`     // 说说内容
		Imgs       sql.NullString `db:"imgs"`        // 图片地址
		Like       uint64         `db:"like"`        // 点赞数量
		CreateTime sql.NullTime   `db:"create_time"` // 发表时间
		UpdateTime sql.NullTime   `db:"update_time"` // 更新时间
		IsDelate   uint64         `db:"is_delate"`   // 逻辑删除
	}
)

func newTalksModel(conn sqlx.SqlConn) *defaultTalksModel {
	return &defaultTalksModel{
		conn:  conn,
		table: "`talks`",
	}
}

func (m *defaultTalksModel) Delete(ctx context.Context, id int64) error {
	query := fmt.Sprintf("delete from %s where `id` = ?", m.table)
	_, err := m.conn.ExecCtx(ctx, query, id)
	return err
}

func (m *defaultTalksModel) FindOne(ctx context.Context, id int64) (*Talks, error) {
	query := fmt.Sprintf("select %s from %s where `id` = ? limit 1", talksRows, m.table)
	var resp Talks
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

func (m *defaultTalksModel) Insert(ctx context.Context, data *Talks) (sql.Result, error) {
	query := fmt.Sprintf("insert into %s (%s) values (?, ?, ?, ?, ?, ?)", m.table, talksRowsExpectAutoSet)
	ret, err := m.conn.ExecCtx(ctx, query, data.Tid, data.Uid, data.Content, data.Imgs, data.Like, data.IsDelate)
	return ret, err
}

func (m *defaultTalksModel) Update(ctx context.Context, data *Talks) error {
	query := fmt.Sprintf("update %s set %s where `id` = ?", m.table, talksRowsWithPlaceHolder)
	_, err := m.conn.ExecCtx(ctx, query, data.Tid, data.Uid, data.Content, data.Imgs, data.Like, data.IsDelate, data.Id)
	return err
}

func (m *defaultTalksModel) tableName() string {
	return m.table
}