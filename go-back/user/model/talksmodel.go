package model

import "github.com/zeromicro/go-zero/core/stores/sqlx"

var _ TalksModel = (*customTalksModel)(nil)

type (
	// TalksModel is an interface to be customized, add more methods here,
	// and implement the added methods in customTalksModel.
	TalksModel interface {
		talksModel
	}

	customTalksModel struct {
		*defaultTalksModel
	}
)

// NewTalksModel returns a model for the database table.
func NewTalksModel(conn sqlx.SqlConn) TalksModel {
	return &customTalksModel{
		defaultTalksModel: newTalksModel(conn),
	}
}
