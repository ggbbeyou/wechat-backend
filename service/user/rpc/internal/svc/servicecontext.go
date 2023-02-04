package svc

import (
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"wechat-backend/service/user/model"
	"wechat-backend/service/user/rpc/internal/config"
)

type ServiceContext struct {
	Config           config.Config
	UserModel        model.UsersModel
	FriendShipsModel model.FriendshipsModel
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:           c,
		UserModel:        model.NewUsersModel(sqlx.NewMysql(c.Mysql.DataSource)),
		FriendShipsModel: model.NewFriendshipsModel(sqlx.NewMysql(c.Mysql.DataSource)),
	}
}
