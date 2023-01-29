package user

import (
	"context"
	"github.com/zeromicro/go-zero/core/stores/sqlc"
	"wechat-backend/common/utils"

	"wechat-backend/service/user/api/internal/svc"
	"wechat-backend/service/user/api/internal/types"

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

func (l *GroupListLogic) GroupList(req *types.UidRequest) (resp *types.GroupListResponse, err error) {
	groupList, err := l.svcCtx.GroupModel.SelectGroupList(l.ctx, req.Uid)
	if err != nil && err != sqlc.ErrNotFound {
		return nil, err
	}
	var groupDetails []types.GroupDetail

	for _, groupInfo := range groupList {
		g := types.GroupDetail{}
		utils.StrCpy(groupInfo, &g)
		groupDetails = append(groupDetails, g)
	}
	return &types.GroupListResponse{
		GroupList: groupDetails,
	}, err
}
