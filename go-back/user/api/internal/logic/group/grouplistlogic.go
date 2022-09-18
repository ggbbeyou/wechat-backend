package group

import (
	"context"
	utils "go-back/common"
	"go-back/user/api/internal/svc"
	"go-back/user/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GroupListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGroupListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GroupListLogic {
	return &GroupListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GroupListLogic) GroupList(req *types.UserId) (resp *types.GroupListResponse, err error) {
	groupList, err := l.svcCtx.GroupsModel.GroupListByUid(l.ctx, req.Uid)
	if err != nil {
		return nil, err
	}
	resp = new(types.GroupListResponse)
	var groupArr []types.GroupInfo
	for _, group := range groupList {
		g := types.GroupInfo{}
		utils.ObjCopy(group, &g)
		groupArr = append(groupArr, g)
	}
	resp.GroupList = groupArr
	return resp, nil
}
