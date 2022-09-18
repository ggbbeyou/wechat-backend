package svc

import (
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"go-back/user/model"
	"go-back/user/rpc/internal/config"
)

type ServiceContext struct {
	Config            config.Config
	GroupfriendsModel model.GroupfriendsModel
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:            c,
		GroupfriendsModel: model.NewGroupfriendsModel(sqlx.NewMysql(c.Mysql.DataSource)),
	}
}
