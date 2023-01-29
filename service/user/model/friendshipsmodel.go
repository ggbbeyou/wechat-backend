package model

import "github.com/zeromicro/go-zero/core/stores/sqlx"

var _ FriendshipsModel = (*customFriendshipsModel)(nil)

type (
	// FriendshipsModel is an interface to be customized, add more methods here,
	// and implement the added methods in customFriendshipsModel.
	FriendshipsModel interface {
		friendshipsModel
	}

	customFriendshipsModel struct {
		*defaultFriendshipsModel
	}
)

// NewFriendshipsModel returns a domain for the database table.
func NewFriendshipsModel(conn sqlx.SqlConn) FriendshipsModel {
	return &customFriendshipsModel{
		defaultFriendshipsModel: newFriendshipsModel(conn),
	}
}
