package friend

import (
	"context"
	"github.com/zeromicro/go-zero/core/stores/sqlc"
	utils "go-back/common"
	"go-back/user/api/internal/svc"
	"go-back/user/api/internal/types"

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

func (l *FriendListLogic) FriendList(req *types.UserId) (resp *types.FriendListResponse, err error) {
	FriendList, err := l.svcCtx.FriendshipsModel.FriendListByUid(l.ctx, req.Uid)
	if err != nil && err != sqlc.ErrNotFound {
		return nil, err
	}
	resp = new(types.FriendListResponse)
	var FriendInfoArr []types.FriendInfo
	for _, friendInfo := range FriendList {
		f := types.FriendInfo{}
		utils.ObjCopy(friendInfo.BaseInfo, &f.BaseInfo)
		FriendInfoArr = append(FriendInfoArr, f)
	}
	resp.FriendList = FriendInfoArr
	return resp, err
}
