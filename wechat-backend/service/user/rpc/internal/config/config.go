package config

import "github.com/zeromicro/go-zero/zrpc"

type Config struct {
	zrpc.RpcServerConf

	//配置mysql
	Mysql struct {
		DataSource string
	}
}
