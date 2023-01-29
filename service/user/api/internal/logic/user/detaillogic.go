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

type DetailLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewDetailLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DetailLogic {
	return &DetailLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DetailLogic) Detail(req *types.UserNameRequest) (resp *types.UserDetailResponse, err error) {
	user, err := l.svcCtx.UserModel.FindOneByUsername(l.ctx, req.Username)

	if err != nil && err != sqlc.ErrNotFound {
		return nil, response.NewRespError(response.USER_NOT_EXIST, response.USER_NOT_EXIST_MESSAGE)
	}

	userInfo := types.UserDetail{}
	utils.StrCpy(user, &userInfo)

	return &types.UserDetailResponse{
		UserInfo: userInfo,
	}, nil
}
