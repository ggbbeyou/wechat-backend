package svc

import (
	"github.com/zeromicro/go-queue/kq"
	"github.com/zeromicro/go-zero/core/stores/redis"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"github.com/zeromicro/go-zero/zrpc"
	"go-back/chat/api/internal/config"
	"go-back/chat/model"
	"go-back/user/rpc/userclient"
)

type ServiceContext struct {
	Config config.Config
	//配置远程rpc服务
	//配置redis
	Redis        *redis.Redis
	UserRpc      userclient.User
	KqPusher     *kq.Pusher
	RecodesModel model.RecodesModel
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config: c,
		//配置远程rpc服务
		UserRpc:  userclient.NewUser(zrpc.MustNewClient(c.UserRpcConf)),
		KqPusher: kq.NewPusher(c.KqConf.Brokers, c.KqConf.Topic),
		//配置redis
		Redis: redis.New(c.Redis.Host, func(r *redis.Redis) {
			r.Type = c.Redis.Type
			r.Pass = c.Redis.Pass
		}),
		//model
		RecodesModel: model.NewRecodesModel(sqlx.NewMysql(c.Mysql.DataSource)),
	}
}
