package handler

import (
	"context"
	"database/sql"
	"encoding/json"
	"github.com/zeromicro/go-zero/core/logx"
	"go-back/chat/api/internal/svc"
	"go-back/chat/model"
	"go-back/common"
	"time"
)

/*
* @Author: chuang
* @Date:   2022/9/16 15:43
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
	logx.Infof("Kafka Consumer 收到消息: value : %s \n", value)
	//写入本地缓存
	consumeCache <- value
	logx.Infof("写入本地缓存")
	//1.协程落盘
	//2.websocket进行处理
	err := Dispatch([]byte(value), c.svcCtx)
	if err != nil {
		return err
	}
	return nil
}

// StoreMsgToMysql 存储到mysql
func (c *ChatManagerMq) StoreMsgToMysql() {
	for {
		select {
		case val := <-consumeCache:
			var msg common.Message
			err := json.Unmarshal([]byte(val), &msg)
			if err != nil {
				logx.Error(err)
				return
			}
			snowFlake, _ := common.NewSnowFlake(0, 0)
			snowId, err := snowFlake.NextId()
			if err != nil {
				logx.Error(err)
				return
			}
			recode := &model.Recodes{
				Id:       snowId,
				Content:  sql.NullString{String: msg.Content, Valid: true},
				From:     sql.NullInt64{Int64: msg.From, Valid: true},
				To:       sql.NullInt64{Int64: msg.To, Valid: true},
				SendTime: sql.NullTime{Time: time.Now(), Valid: true},
			}
			_, err = c.svcCtx.RecodesModel.Insert(c.ctx, recode)
			if err != nil {
				return
			}
		}
	}
}
