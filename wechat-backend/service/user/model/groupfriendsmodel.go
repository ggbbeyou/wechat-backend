package model

import "github.com/zeromicro/go-zero/core/stores/sqlx"

var _ GroupfriendsModel = (*customGroupfriendsModel)(nil)

type (
	// GroupfriendsModel is an interface to be customized, add more methods here,
	// and implement the added methods in customGroupfriendsModel.
	GroupfriendsModel interface {
		groupfriendsModel
	}

	customGroupfriendsModel struct {
		*defaultGroupfriendsModel
	}
)

// NewGroupfriendsModel returns a domain for the database table.
func NewGroupfriendsModel(conn sqlx.SqlConn) GroupfriendsModel {
	return &customGroupfriendsModel{
		defaultGroupfriendsModel: newGroupfriendsModel(conn),
	}
}
