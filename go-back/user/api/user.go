package main

import (
	"flag"
	"fmt"
	"go-back/user/api/internal/config"
	"go-back/user/api/internal/handler"
	"go-back/user/api/internal/svc"

	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/rest"
)

//var configFile = flag.String("f", "/service/api/etc/service-api.yaml", "the config file")
var configFile = flag.String("f", "user/api/etc/user-api.yaml", "the config file")

func main() {
	flag.Parse()

	var c config.Config
	conf.MustLoad(*configFile, &c)

	server := rest.MustNewServer(c.RestConf)
	defer server.Stop()

	ctx := svc.NewServiceContext(c)
	handler.RegisterHandlers(server, ctx)

	fmt.Printf("Starting server at %s:%d...\n", c.Host, c.Port)
	server.Start()
}
