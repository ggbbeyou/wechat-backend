package config

import (
	"github.com/zeromicro/go-queue/kq"
	"github.com/zeromicro/go-zero/rest"
	"github.com/zeromicro/go-zero/zrpc"
)

type Config struct {
	rest.RestConf

	//配置mysql
	Mysql struct {
		DataSource string
	}
	//redis
	Redis struct {
		Host string
		Type string
		Pass string
	}

	//配置rpc远程服务
	//ChatRpcConf zrpc.RpcClientConf
	//配置user的rpc服务
	UserRpcConf zrpc.RpcClientConf

	KqConf kq.KqConf
}
