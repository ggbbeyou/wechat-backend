package chatconn

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/websocket"
	"github.com/zeromicro/go-queue/kq"
	"github.com/zeromicro/go-zero/core/logx"
	"wechat-backend/common/domain"
	"wechat-backend/common/utils"
	"wechat-backend/service/communicate/api/internal/svc"
)

/*
* @Author: chuang
* @Date:   2023/1/15 17:43
 */

func StartMq(svcCtx *svc.ServiceContext) {
	q := kq.MustNewQueue(svcCtx.Config.KqConf, kq.WithHandle(func(k, v string) error {
		logx.Info("kafka消费消息.......")
		err := dispatch(v)
		if err != nil {
			return err
		}
		return nil
	}))
	defer q.Stop()
	q.Start()
}

// Dispatch 处理kafka的消息，分类消费
func dispatch(value string) error {
	var message domain.Message
	err := json.Unmarshal([]byte(value), &message)
	if err != nil {
		logx.Errorf("解析kafka消息异常")
		return err
	}
	switch message.Type {
	case domain.USER_TYPE:
		sendUserMessage(message)
		break
	case domain.GROUP_TYPE:
		break
	}
	return nil
}

//sendUserMessage 处理用户消息
func sendUserMessage(message domain.Message) {
	node, ok := usersMap[message.To]
	if !ok {
		return
	}
	if node.WsConn == (*websocket.Conn)(nil) {
		//用户不在线 存储到Redis
		key := fmt.Sprintf("%s::%d", utils.OFFLINE_MESSAGE, message.To)
		_, err := node.SvcCtx.Redis.Lpush(key, message)
		if err != nil {
			logx.Error(err)
		}
		return
	}
	node.CacheOnlineMessage <- message
}
