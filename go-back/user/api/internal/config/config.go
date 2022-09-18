package config

import "github.com/zeromicro/go-zero/rest"

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
}
