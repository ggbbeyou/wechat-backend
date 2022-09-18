package logic

import (
	"context"
	"fmt"

	"go-back/user/rpc/internal/svc"
	"go-back/user/rpc/user"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetGroupIdListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetGroupIdListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetGroupIdListLogic {
	return &GetGroupIdListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetGroupIdListLogic) GetGroupIdList(in *user.UidRequest) (*user.GroupIdListResponse, error) {
	logx.Info(in.Uid)
	fmt.Println("awdawdawd")
	//查询用户的所有群id
	gidArr, err := l.svcCtx.GroupfriendsModel.QueryGroupIdsByUid(l.ctx, in.Uid)
	logx.Info(in.Uid)
	if err != nil {
		return nil, err
	}
	return &user.GroupIdListResponse{
		GidArr: gidArr,
	}, nil
}
