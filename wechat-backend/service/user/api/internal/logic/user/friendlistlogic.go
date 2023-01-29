package user

import (
	"context"
	"github.com/zeromicro/go-zero/core/stores/sqlc"
	"wechat-backend/common/utils"

	"wechat-backend/service/user/api/internal/svc"
	"wechat-backend/service/user/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type FriendListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewFriendListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FriendListLogic {
	return &FriendListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *FriendListLogic) FriendList(req *types.UidRequest) (resp *types.FriendListResponse, err error) {
	friendList, err := l.svcCtx.FriendShipsModel.SelectFriendList(l.ctx, req.Uid)
	if err != nil && err != sqlc.ErrNotFound {
		return nil, err
	}

	var userDetails []types.UserDetail

	for _, friendInfo := range friendList {
		userInfo := types.UserDetail{}
		utils.StrCpy(friendInfo, &userInfo)
		userDetails = append(userDetails, userInfo)
	}
	return &types.FriendListResponse{
		FriendList: userDetails,
	}, err
}
