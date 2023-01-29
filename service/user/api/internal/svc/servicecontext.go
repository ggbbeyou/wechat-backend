package svc

import (
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"github.com/zeromicro/go-zero/rest"
	"wechat-backend/service/user/api/internal/config"
	"wechat-backend/service/user/api/internal/middleware"
	"wechat-backend/service/user/model"
)

type ServiceContext struct {
	Config                    config.Config
	VerifyAuthorityMiddleware rest.Middleware

	UserModel        model.UsersModel
	FriendShipsModel model.FriendshipsModel
	GroupModel       model.GroupsModel
	GroupFriendModel model.GroupfriendsModel
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config: c,

		VerifyAuthorityMiddleware: middleware.NewVerifyAuthorityMiddleware(c).HandleVerifyAuth,
		UserModel:                 model.NewUsersModel(sqlx.NewMysql(c.Mysql.DataSource)),
		FriendShipsModel:          model.NewFriendshipsModel(sqlx.NewMysql(c.Mysql.DataSource)),
		GroupModel:                model.NewGroupsModel(sqlx.NewMysql(c.Mysql.DataSource)),
		GroupFriendModel:          model.NewGroupfriendsModel(sqlx.NewMysql(c.Mysql.DataSource)),
	}

}
