package logic

import (
	"context"
	"github.com/zeromicro/go-zero/core/stores/sqlc"

	"wechat-backend/service/user/rpc/internal/svc"
	"wechat-backend/service/user/rpc/user"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetFriendIdListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetFriendIdListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetFriendIdListLogic {
	return &GetFriendIdListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// GetFriendIdList 获取所有好友
func (l *GetFriendIdListLogic) GetFriendIdList(in *user.GetFriendIdListRequest) (*user.UIdListResponse, error) {
	friendIdList, err := l.svcCtx.FriendShipsModel.GetFriendIdList(l.ctx, in.GetUid())
	if err != nil && err != sqlc.ErrNotFound {
		return nil, err
	}
	return &user.UIdListResponse{
		UidArr: friendIdList,
	}, nil
}
