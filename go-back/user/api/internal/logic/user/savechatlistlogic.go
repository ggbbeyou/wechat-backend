package user

import (
	"context"
	"encoding/json"
	"fmt"
	"go-back/common"
	"strconv"

	"go-back/user/api/internal/svc"
	"go-back/user/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type SaveChatListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewSaveChatListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SaveChatListLogic {
	return &SaveChatListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *SaveChatListLogic) SaveChatList(req *types.SaveChatListRequest) error {
	key := fmt.Sprintf("%s%d", common.CHATLIST, req.Uid)
	marshal, err := json.Marshal(req.UserChatListInfo)
	if err != nil {
		logx.Error(err)
		return err
	}
	err = l.svcCtx.Redis.HsetCtx(
		l.ctx,
		key,
		strconv.FormatInt(req.UserChatListInfo.Uid, 10),
		string(marshal),
	)
	if err != nil {
		logx.Error(err)
		return err
	}
	//每次添加列表都更新过期时间
	err = l.svcCtx.Redis.ExpireCtx(l.ctx, key, common.CHATLIST_EXPIRE)
	if err != nil {
		logx.Error(err)
		return err
	}
	return nil
}
