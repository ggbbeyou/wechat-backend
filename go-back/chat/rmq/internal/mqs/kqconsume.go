package mqs

import (
	"context"
	"github.com/zeromicro/go-zero/core/logx"
	"go-back/chat/rmq/internal/svc"
)

/*
* @Author: chuang
* @Date:   2022/9/15 12:09
 */

//监听聊天通知消息队列

type ChatManagerMq struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewChatManagerMq(ctx context.Context, svcCtx *svc.ServiceContext) *ChatManagerMq {
	return &ChatManagerMq{
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (c *ChatManagerMq) Consume(_, value string) error {
	//kafka消费消息,业务处理
	logx.Infof("Kafka Consumer value : %s \n", value)

	return nil
}
