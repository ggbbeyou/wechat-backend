package main

import (
	"flag"
	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/core/service"
	"go-back/chat/rmq/internal/config"
	"go-back/chat/rmq/internal/listen"
)

/*
* @Author: chuang
* @Date:   2022/9/15 10:25
 */

var configFile = flag.String("f", "chat/rmq/etc/chat.yaml", "the config file")

func main() {
	flag.Parse()
	var c config.Config
	conf.MustLoad(*configFile, &c)
	//kafka配置
	//if err := c.SetUp(); err != nil {
	//	panic(err)
	//}
	//创建服务组，用户运行kafka 消费者队列
	serviceGroup := service.NewServiceGroup()
	defer serviceGroup.Stop()
	for _, mq := range listen.Mqs(c) {
		serviceGroup.Add(mq)
	}
	serviceGroup.Start()
}

//
//{"@timestamp":"2022-09-15T17:18:34.901+08:00",
//	"caller":"kq/queue.go:170",
//	"content":"Error on reading message, \"[15] Group Coordinator Not Available: the broker returns this error code for group coordinator requests, offset commits, and most group management requests if the offsets topic has not yet been created, or if the group coordinator is not active\"","level":"error"}
