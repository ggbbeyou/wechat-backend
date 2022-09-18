package handler

import (
	"context"
	"github.com/zeromicro/go-queue/kq"
	"github.com/zeromicro/go-zero/core/service"
	"go-back/chat/api/internal/svc"
	"go-back/common"
)

/*
* @Author: chuang
* @Date:   2022/9/17 13:17
 */
//func Mqs(c config.Config) []service.Service {
//	svcContext := svc.NewServiceContext(c)
//	ctx := context.Background()
//	var services []service.Service
//	services = append(services, kqMqs(c, ctx, svcContext)...)
//	return services
//}
//
//
//func kqMqs(c config.Config, ctx context.Context, svcContext *svc.ServiceContext) []service.Service {
//	return []service.Service{
//		kq.MustNewQueue(c.KqConf, NewChatManagerMq(ctx, svcContext)),
//	}
//}

func Mqs(svcContext *svc.ServiceContext, managerMq *ChatManagerMq) []service.Service {
	var services []service.Service
	services = append(services, kqMqs(svcContext, managerMq)...)
	return services
}

func kqMqs(svcContext *svc.ServiceContext, managerMq *ChatManagerMq) []service.Service {
	return []service.Service{
		kq.MustNewQueue(svcContext.Config.KqConf, managerMq),
	}
}

var consumeCache chan string

// LoadMq 初始化kafka服务
func LoadMq(svcContext *svc.ServiceContext) {
	ctx := context.Background()
	managerMq := NewChatManagerMq(ctx, svcContext)

	//创建服务组，用户运行kafka 消费者队列
	serviceGroup := service.NewServiceGroup()
	defer serviceGroup.Stop()
	for _, mq := range Mqs(svcContext, managerMq) {
		serviceGroup.Add(mq)
	}
	consumeCache = make(chan string, 1024)
	//开启协程保存数据到mysql
	common.Pool.Submit(managerMq.StoreMsgToMysql)
	serviceGroup.Start()
}
