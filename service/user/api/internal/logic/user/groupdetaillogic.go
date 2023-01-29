package user

import (
	"context"
	"github.com/zeromicro/go-zero/core/stores/sqlc"
	"wechat-backend/common/response"
	"wechat-backend/common/utils"

	"wechat-backend/service/user/api/internal/svc"
	"wechat-backend/service/user/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GroupDetailLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGroupDetailLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GroupDetailLogic {
	return &GroupDetailLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GroupDetailLogic) GroupDetail(req *types.GroupIdRequest) (resp *types.GroupDetailResponse, err error) {

	group, err := l.svcCtx.GroupModel.SelectOne(l.ctx, req.Gid)
	if err != nil && err != sqlc.ErrNotFound {
		return nil, response.NewRespError(response.GROUP_NOT_EXIST, response.GROUP_NOT_EXIST_MESSAGE)
	}

	groupInfo := types.GroupDetail{}
	utils.StrCpy(group, &groupInfo)

	return &types.GroupDetailResponse{
		GroupInfo: groupInfo,
	}, err
}
