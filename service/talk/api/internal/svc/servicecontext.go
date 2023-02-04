package svc

import (
	"github.com/zeromicro/go-zero/core/stores/redis"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"github.com/zeromicro/go-zero/rest"
	"github.com/zeromicro/go-zero/zrpc"
	"wechat-backend/service/talk/api/internal/config"
	"wechat-backend/service/talk/api/internal/middleware"
	"wechat-backend/service/talk/model"
	"wechat-backend/service/user/rpc/userclient"
)

type ServiceContext struct {
	Config                    config.Config
	VerifyAuthorityMiddleware rest.Middleware
	CommentModel              model.CommentsModel
	TalkModel                 model.TalksModel
	LikeModel                 model.LikesModel
	Redis                     *redis.Redis
	UserRpcService            userclient.User
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:                    c,
		VerifyAuthorityMiddleware: middleware.NewVerifyAuthorityMiddleware(c).HandleVerifyAuth,
		CommentModel:              model.NewCommentsModel(sqlx.NewMysql(c.Mysql.DataSource)),
		TalkModel:                 model.NewTalksModel(sqlx.NewMysql(c.Mysql.DataSource)),
		LikeModel:                 model.NewLikesModel(sqlx.NewMysql(c.Mysql.DataSource)),
		Redis:                     redis.New(c.Redis.Host, redisConfig(c)),
		UserRpcService:            userclient.NewUser(zrpc.MustNewClient(c.UserRpcService)),
	}
}

//redisConfig 配置redis
func redisConfig(c config.Config) redis.Option {
	return func(r *redis.Redis) {
		r.Type = c.Redis.Type
		r.Pass = c.Redis.Pass
	}
}
