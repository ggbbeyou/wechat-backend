package model

import "github.com/zeromicro/go-zero/core/stores/sqlx"

var _ GroupsModel = (*customGroupsModel)(nil)

type (
	// GroupsModel is an interface to be customized, add more methods here,
	// and implement the added methods in customGroupsModel.
	GroupsModel interface {
		groupsModel
	}

	customGroupsModel struct {
		*defaultGroupsModel
	}
)

// NewGroupsModel returns a domain for the database table.
func NewGroupsModel(conn sqlx.SqlConn) GroupsModel {
	return &customGroupsModel{
		defaultGroupsModel: newGroupsModel(conn),
	}
}
