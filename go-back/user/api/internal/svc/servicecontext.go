package svc

import "C"
import (
	"github.com/zeromicro/go-zero/core/stores/redis"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"github.com/zeromicro/go-zero/rest"
	"go-back/user/api/internal/config"
	"go-back/user/api/internal/middleware"
	"go-back/user/model"
)

type ServiceContext struct {
	Config         config.Config
	AuthMiddleware rest.Middleware

	UserModel         model.UsersModel
	GroupsModel       model.GroupsModel
	CommentsModel     model.CommentsModel
	FriendshipsModel  model.FriendshipsModel
	GroupfriendsModel model.GroupfriendsModel
	TalksModel        model.TalksModel
	//配置redis
	Redis *redis.Redis
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:            c,
		AuthMiddleware:    middleware.NewAuthMiddleware().Handle,
		UserModel:         model.NewUsersModel(sqlx.NewMysql(c.Mysql.DataSource)),
		GroupsModel:       model.NewGroupsModel(sqlx.NewMysql(c.Mysql.DataSource)),
		CommentsModel:     model.NewCommentsModel(sqlx.NewMysql(c.Mysql.DataSource)),
		FriendshipsModel:  model.NewFriendshipsModel(sqlx.NewMysql(c.Mysql.DataSource)),
		GroupfriendsModel: model.NewGroupfriendsModel(sqlx.NewMysql(c.Mysql.DataSource)),
		TalksModel:        model.NewTalksModel(sqlx.NewMysql(c.Mysql.DataSource)),
		//配置redis
		Redis: redis.New(c.Redis.Host, func(r *redis.Redis) {
			r.Type = c.Redis.Type
			r.Pass = c.Redis.Pass
		}),
	}
}
