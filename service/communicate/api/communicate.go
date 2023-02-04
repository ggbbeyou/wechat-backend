package main

import (
	"flag"
	"fmt"
	"github.com/zeromicro/go-zero/core/logx"
	"wechat-backend/service/communicate/api/internal/handler/chatconn"

	"wechat-backend/service/communicate/api/internal/config"
	"wechat-backend/service/communicate/api/internal/handler"
	"wechat-backend/service/communicate/api/internal/svc"

	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/rest"
)

//var configFile = flag.String("f", "etc/communicate-api.yaml", "the config file")
var configFile = flag.String("f", "etc/communicate-api.yaml", "the config file")

func main() {
	flag.Parse()

	var c config.Config
	conf.MustLoad(*configFile, &c)

	server := rest.MustNewServer(c.RestConf)
	defer server.Stop()

	ctx := svc.NewServiceContext(c)
	handler.RegisterHandlers(server, ctx)

	//初始化所有用户
	err := chatconn.InitAllUser(ctx)
	if err != nil {
		logx.Error(err)
		return
	}
	//初始化kafka消费者
	go func(ctx *svc.ServiceContext) {
		chatconn.StartMq(ctx)
	}(ctx)

	fmt.Printf("Starting server at %s:%d...\n", c.Host, c.Port)
	server.Start()
}
