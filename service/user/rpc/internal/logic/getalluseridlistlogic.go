package logic

import (
	"context"
	"github.com/zeromicro/go-zero/core/stores/sqlc"

	"wechat-backend/service/user/rpc/internal/svc"
	"wechat-backend/service/user/rpc/user"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetAllUserIdListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetAllUserIdListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetAllUserIdListLogic {
	return &GetAllUserIdListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetAllUserIdListLogic) GetAllUserIdList(in *user.Empty) (*user.UIdListResponse, error) {
	userIdList, err := l.svcCtx.UserModel.SelectAllUserIdList(l.ctx)
	if err != nil && err != sqlc.ErrNotFound {
		return nil, err
	}
	return &user.UIdListResponse{
		UidArr: userIdList,
	}, nil
}
