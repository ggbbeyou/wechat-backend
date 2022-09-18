package friend

import (
	"context"
	utils "go-back/common"
	"go-back/user/api/internal/svc"
	"go-back/user/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type FriendDetailLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewFriendDetailLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FriendDetailLogic {
	return &FriendDetailLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *FriendDetailLogic) FriendDetail(req *types.UserId) (resp *types.FriendDetailResponse, err error) {
	friendDetail, err := l.svcCtx.UserModel.FindOneByUid(l.ctx, req.Uid)
	if err != nil {
		return nil, err
	}
	resp = new(types.FriendDetailResponse)
	utils.ObjCopy(friendDetail, &resp.BaseInfo)
	return resp, err
}
