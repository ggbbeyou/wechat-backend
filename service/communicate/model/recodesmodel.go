package model

import "github.com/zeromicro/go-zero/core/stores/sqlx"

var _ RecodesModel = (*customRecodesModel)(nil)

type (
	// RecodesModel is an interface to be customized, add more methods here,
	// and implement the added methods in customRecodesModel.
	RecodesModel interface {
		recodesModel
	}

	customRecodesModel struct {
		*defaultRecodesModel
	}
)

// NewRecodesModel returns a domain for the database table.
func NewRecodesModel(conn sqlx.SqlConn) RecodesModel {
	return &customRecodesModel{
		defaultRecodesModel: newRecodesModel(conn),
	}
}
