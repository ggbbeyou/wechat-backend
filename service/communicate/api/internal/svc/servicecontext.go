package svc

import (
	"github.com/zeromicro/go-queue/kq"
	"github.com/zeromicro/go-zero/core/stores/redis"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"github.com/zeromicro/go-zero/rest"
	"github.com/zeromicro/go-zero/zrpc"
	"wechat-backend/service/communicate/api/internal/config"
	"wechat-backend/service/communicate/api/internal/middleware"
	"wechat-backend/service/communicate/model"
	"wechat-backend/service/user/rpc/userclient"
)

type ServiceContext struct {
	Config                    config.Config
	VerifyAuthorityMiddleware rest.Middleware
	Redis                     *redis.Redis
	KqPusher                  *kq.Pusher
	RecodesModel              model.RecodesModel
	UserRpcService            userclient.User
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:                    c,
		VerifyAuthorityMiddleware: middleware.NewVerifyAuthorityMiddleware(c).HandleVerifyAuth,
		RecodesModel:              model.NewRecodesModel(sqlx.NewMysql(c.Mysql.DataSource)),
		Redis: redis.New(c.Redis.Host, func(r *redis.Redis) {
			r.Type = c.Redis.Type
			r.Pass = c.Redis.Pass
		}),
		KqPusher:       kq.NewPusher(c.KqConf.Brokers, c.KqConf.Topic),
		UserRpcService: userclient.NewUser(zrpc.MustNewClient(c.UserRpcService)),
	}

}
