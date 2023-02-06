package config

import (
	"github.com/zeromicro/go-zero/rest"
	"github.com/zeromicro/go-zero/zrpc"
)

type Config struct {
	rest.RestConf

	Auth struct {
		AccessSecret string
		AccessExpire int64
	}

	Mysql struct {
		DataSource string
	}

	Redis struct {
		Host string
		Type string
		Pass string
	}

	UserRpcService zrpc.RpcClientConf

	//Cronjob:
	//Duration: 7200
	Cronjob struct {
		Duration int64
	}
}
