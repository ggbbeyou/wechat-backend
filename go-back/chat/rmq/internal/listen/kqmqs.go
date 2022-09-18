package listen

import (
	"context"
	"github.com/zeromicro/go-queue/kq"
	"github.com/zeromicro/go-zero/core/service"
	"go-back/chat/rmq/internal/config"
	"go-back/chat/rmq/internal/mqs"
	"go-back/chat/rmq/internal/svc"
)

/*
* @Author: chuang
* @Date:   2022/9/15 16:19
 */
// Mqs 返回所有消费者 ,将每一个消费者都封装成一个go-zero service

func Mqs(c config.Config) []service.Service {
	svcContext := svc.NewServiceContext(c)
	ctx := context.Background()
	var services []service.Service
	services = append(services, kqMqs(c, ctx, svcContext)...)
	return services
}

func kqMqs(c config.Config, ctx context.Context, svcContext *svc.ServiceContext) []service.Service {
	return []service.Service{
		kq.MustNewQueue(c.KqConf, mqs.NewChatManagerMq(ctx, svcContext)),
	}
}
