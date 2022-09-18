package user

import (
	"context"
	"encoding/json"
	"fmt"

	"go-back/user/api/internal/svc"
	"go-back/user/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UserChatListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUserChatListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserChatListLogic {
	return &UserChatListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

// UserChatList 通过redis获取聊天列表
func (l *UserChatListLogic) UserChatList(req *types.UserId) (resp *types.UserChatListInfoResponse, err error) {
	key := fmt.Sprintf("chatlist:%d", req.Uid)
	data, _ := l.svcCtx.Redis.HgetallCtx(l.ctx, key)
	userChat := make([]types.UserChatListInfo, len(data))
	idx := 0
	for _, v := range data {
		err := json.Unmarshal([]byte(v), &userChat[idx])
		if err != nil {
			return nil, err
		}
		idx++
	}
	return &types.UserChatListInfoResponse{
		UserChatList: userChat,
	}, nil
}
