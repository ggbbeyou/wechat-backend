package chatconn

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/gorilla/websocket"
	"github.com/zeromicro/go-zero/core/logx"
	"net/http"
	"strconv"
	"time"
	"wechat-backend/common/domain"
	"wechat-backend/common/response"
	"wechat-backend/common/utils"
	"wechat-backend/service/communicate/api/internal/svc"
	"wechat-backend/service/user/rpc/user"
)

//UsersMap 存储在线用户
var usersMap map[int64]*Node

type Node struct {
	Uid    int64
	WsConn *websocket.Conn
	//缓存用户消息
	CacheOnlineMessage chan domain.Message
	//缓存离线消息
	CachOfflineMessage chan domain.Message

	SvcCtx *svc.ServiceContext
}

func ChatConnHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		query, _ := strconv.Atoi(r.URL.Query().Get("uid"))
		uid := int64(query)

		node, ok := usersMap[uid]
		if !ok {
			return
		}
		conn, err := upgrade(w, r, svcCtx)
		if err != nil {
			response.Response(w, nil, response.NewRespError(response.WEBCOKET_ERROR, response.WEBCOKET_ERROR_MESSAGE))
			return
		}
		//绑定用户和ws
		usersMap[uid].WsConn = conn
		//每个用户开启goroutine监听消息收发
		utils.Pool.Submit(node.SendMQMessage)
		utils.Pool.Submit(node.RecvMessage)
		logx.Infof("用户%d连接成功......", uid)
	}
}

/**
协议升级 time.Duration(expire) * time.Second)
*/
func upgrade(w http.ResponseWriter, r *http.Request, svcCtx *svc.ServiceContext) (*websocket.Conn, error) {
	ws := websocket.Upgrader{
		HandshakeTimeout: time.Duration(svcCtx.Config.Client.Upgrade.HandshakeTimeout) * time.Second,
		ReadBufferSize:   int(svcCtx.Config.Client.Upgrade.ReadBufferSize),
		WriteBufferSize:  int(svcCtx.Config.Client.Upgrade.WriteBufferSize),
		CheckOrigin: func(r *http.Request) bool {
			return true
		},
	}
	conn, err := ws.Upgrade(w, r, nil)
	if err != nil {
		return nil, err
	}
	return conn, nil
}

// InitAllUser 初始化用户
func InitAllUser(svcCtx *svc.ServiceContext) error {
	ctx := context.Background()
	userIdList, err := svcCtx.UserRpcService.GetAllUserIdList(ctx, &user.Empty{})
	if err != nil {
		logx.Error("初始化用户异常")
		return err
	}
	usersMap = make(map[int64]*Node, len(userIdList.UidArr))

	for _, uid := range userIdList.UidArr {
		usersMap[uid] = &Node{
			Uid:                uid,
			WsConn:             (*websocket.Conn)(nil),
			CacheOnlineMessage: make(chan domain.Message, svcCtx.Config.Client.MessageBuf),
			CachOfflineMessage: make(chan domain.Message, svcCtx.Config.Client.MessageBuf),
			SvcCtx:             svcCtx,
		}
	}

	logx.Info("初始化用户成功, 用户Uid:")
	fmt.Println(userIdList.UidArr)

	return nil
}

// SendMQMessage 接收到消息,发送到kafka
func (n *Node) SendMQMessage() {
	defer func() {
		n.WsConn.Close()
		usersMap[n.Uid].WsConn = (*websocket.Conn)(nil)
		logx.Errorf("用户%d的websocket连接断开", n.Uid)
	}()
	for {
		_, buf, err := n.WsConn.ReadMessage()
		logx.Infof(string(buf))
		if err != nil {
			return
		}
		err = n.SvcCtx.KqPusher.Push(string(buf))
		logx.Info("kafka存储消息....")
		if err != nil {
			logx.Errorf("kafka接收消息异常", n.Uid)
			return
		}
	}
}

// RecvMessage 接受调度过来的消息
func (n *Node) RecvMessage() {
	for {
		select {
		case message := <-n.CacheOnlineMessage:
			buf, err := json.Marshal(message)
			logx.Infof("用户%d收到消息：%s", n.Uid, string(buf))
			if err != nil {
				logx.Error(err)
				return
			}
			err = n.WsConn.WriteMessage(websocket.TextMessage, buf)
			if err != nil {
				logx.Error(err)
				return
			}
		}
	}
}

//处理离线消息
