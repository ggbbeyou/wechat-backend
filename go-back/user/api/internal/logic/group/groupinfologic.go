package group

import (
	"context"
	utils "go-back/common"
	"go-back/user/api/internal/svc"
	"go-back/user/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GroupInfoLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGroupInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GroupInfoLogic {
	return &GroupInfoLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GroupInfoLogic) GroupInfo(req *types.GroupId) (resp *types.GroupDetailResponse, err error) {
	groupDetail, err := l.svcCtx.GroupsModel.GroupByGid(l.ctx, req.Gid)
	if err != nil {
		return nil, err
	}
	resp = new(types.GroupDetailResponse)
	resp.Nickname = groupDetail.Nickname
	utils.ObjCopy(groupDetail.GroupBaseInfo, &resp.GroupInfo)
	//TODO 群聊成员信息
	resp.FriendList = nil
	return resp, nil
}
