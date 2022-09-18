package user

import (
	"context"
	"github.com/zeromicro/go-zero/core/stores/sqlc"
	utils "go-back/common"
	"go-back/user/api/internal/svc"
	"go-back/user/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UserInfoLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUserInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserInfoLogic {
	return &UserInfoLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UserInfoLogic) UserInfo(req *types.UserId) (resp *types.UserInfoResponse, err error) {
	user, err := l.svcCtx.UserModel.FindOneByUid(l.ctx, req.Uid)
	if err != nil && err != sqlc.ErrNotFound {
		return nil, err
	}
	resp = new(types.UserInfoResponse)
	userInfo := types.UserInfo{}
	utils.ObjCopy(user, &userInfo)
	resp.Data = userInfo
	return resp, err
}
