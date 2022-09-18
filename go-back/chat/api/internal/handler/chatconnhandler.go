package handler

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/websocket"
	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/rest/httpx"
	"go-back/chat/api/internal/svc"
	"go-back/common"
	"go-back/user/rpc/user"
	"gopkg.in/fatih/set.v0"
	"net/http"
	"strconv"
	"sync"
)

var (
	//存储所有在线用户
	clientMap = make(map[int64]*Node)
)

// Node 用户和ws映射
type Node struct {
	Uid    int64
	WsConn *websocket.Conn
	//缓存用户消息
	Buf chan common.Message
	//存储用户所有群号
	GroupIdArr set.Interface
	SvcCtx     *svc.ServiceContext
}

//读写锁
var rwlocker sync.RWMutex

func ChatConnHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		query, _ := strconv.Atoi(r.URL.Query().Get("uid"))
		uid := int64(query)
		//websocket协议升级
		ws := websocket.Upgrader{
			ReadBufferSize:  1024,
			WriteBufferSize: 1024,
			CheckOrigin: func(r *http.Request) bool {
				return true
			},
		}
		conn, err := ws.Upgrade(w, r, nil)
		if err != nil {
			httpx.Error(w, err)
			return
		}
		//绑定用户和ws
		node := &Node{
			Uid:        uid,
			WsConn:     conn,
			Buf:        make(chan common.Message, 128),
			GroupIdArr: set.New(set.ThreadSafe),
			SvcCtx:     svcCtx,
		}
		//远程方法获取用户的全部群id
		groupIdList, err := svcCtx.UserRpc.GetGroupIdList(r.Context(), &user.UidRequest{Uid: uid})
		if err != nil {
			return
		}
		for _, num := range groupIdList.GidArr {
			node.GroupIdArr.Add(num)
		}
		//加写锁保证线程安全
		rwlocker.Lock()
		//将用户信息和node存放在map
		clientMap[uid] = node
		rwlocker.Unlock()
		//开启协程
		common.Pool.Submit(node.sendproc)
		common.Pool.Submit(node.recvproc)
		//获取该用户的离线消息
		err = common.Pool.Submit(node.GetMsgOnCache)
		if err != nil {
			return
		}
	}
}

//收到其他用户的消息,推送到客户端 监听每一个node的ws,
func (node *Node) sendproc() {
	for {
		select {
		case data := <-node.Buf:
			marshal, err := json.Marshal(data)
			logx.Infof("用户收到消息：%s", string(marshal))
			if err != nil {
				return
			}
			err = node.WsConn.WriteMessage(websocket.TextMessage, marshal)
			if err != nil {
				logx.Error(err)
				return
			}
		}
	}
}

//接收到消息后,发送到kafka
func (n *Node) recvproc() {
	for {
		_, data, err := n.WsConn.ReadMessage()
		logx.Info(string(data))
		if err != nil {
			//遇到异常关闭ws
			_ = n.WsConn.Close()
			delete(clientMap, n.Uid)
			logx.Errorf("用户 %d 的websocket连接断开", n.Uid)
			return
		}
		//Dispatch(data, n.SvcCtx)

		//发送到kafka
		err = n.SvcCtx.KqPusher.Push(string(data))
		if err != nil {
			_ = n.SvcCtx.KqPusher.Close()
			return
		}
	}
}

// Dispatch 调度逻辑处理
func Dispatch(data []byte, svcCtx *svc.ServiceContext) error {
	var msg common.Message
	err := json.Unmarshal(data, &msg)
	if err != nil {
		logx.Error(err)
		return err
	}
	switch msg.Type {
	case common.USER_TYPE:
		//发送消息到用户
		SendUserMsg(msg, svcCtx)
	case common.GROUP_TYPE:
		//发送消息到群聊
	}
	return nil
}

// SendUserMsg 往用户的channel中发送消息
func SendUserMsg(msg common.Message, svcCtx *svc.ServiceContext) {
	rwlocker.RLock()
	node, ok := clientMap[msg.To]
	//用户不在线
	if !ok {
		//发送到redis
		marshal, err := json.Marshal(msg)
		if err != nil {
			rwlocker.RUnlock()
			return
		}
		key := fmt.Sprintf("msgcache:%d", msg.To)
		_, err = svcCtx.Redis.Lpush(key, string(marshal))
		if err != nil {
			rwlocker.RUnlock()
			return
		}
	} else {
		node.Buf <- msg
	}
	rwlocker.RUnlock()
}

// GetMsgOnCache 从redis中获取离线消息, 发送到kafka
func (n *Node) GetMsgOnCache() {
	cacheMsgs, err := n.SvcCtx.Redis.Lrange(fmt.Sprintf("%s%d", common.MSGCACHE, n.Uid), 0, -1)
	if err != nil {
		logx.Error(err)
		return
	}
	//没有缓存消息
	if len(cacheMsgs) == 0 {
		return
	}
	for _, val := range cacheMsgs {
		var msg common.Message
		err := json.Unmarshal([]byte(val), &msg)
		if err != nil {
			logx.Error(err)
			return
		}
		err = n.SvcCtx.KqPusher.Push(val)
		if err != nil {
			logx.Error(err)
			return
		}
	}
	//删除key
	_, err = n.SvcCtx.Redis.Del(fmt.Sprintf("%s%d", common.MSGCACHE, n.Uid))
	if err != nil {
		logx.Error(err)
		return
	}
}
