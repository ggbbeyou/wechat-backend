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

type GroupCrewLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGroupCrewLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GroupCrewLogic {
	return &GroupCrewLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GroupCrewLogic) GroupCrew(req *types.GroupIdRequest) (resp *types.GroupCrewResponse, err error) {
	groupCrew, err := l.svcCtx.GroupFriendModel.SelectGroupCrew(l.ctx, req.Gid)
	if err != nil && err != sqlc.ErrNotFound {
		return nil, response.NewRespError(response.FAIL, response.SYSTEM_FAIL_MESSAGE)
	}

	var userDetail []types.UserDetail
	for _, crew := range groupCrew {
		f := types.UserDetail{}
		utils.StrCpy(crew, &f)
		userDetail = append(userDetail, f)
	}

	return &types.GroupCrewResponse{
		Crew: userDetail,
	}, err
}
