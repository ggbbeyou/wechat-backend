package main

import (
	"flag"
	"fmt"
	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/rest"
	"go-back/chat/api/internal/config"
	"go-back/chat/api/internal/handler"
	"go-back/chat/api/internal/svc"
)

var configFile = flag.String("f", "chat/api/etc/chat-api.yaml", "the config file")

func main() {
	flag.Parse()

	var c config.Config
	conf.MustLoad(*configFile, &c)

	server := rest.MustNewServer(c.RestConf)
	defer server.Stop()

	ctx := svc.NewServiceContext(c)
	handler.RegisterHandlers(server, ctx)

	fmt.Printf("Starting server at %s:%d...\n", c.Host, c.Port)
	//开启协程运行kafka

	go func() {
		handler.LoadMq(ctx)
	}()

	server.Start()
}
