package logic

import (
	"context"

	"go-back/chat/rpc/chat"
	"go-back/chat/rpc/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type TestGroupIdListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewTestGroupIdListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *TestGroupIdListLogic {
	return &TestGroupIdListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *TestGroupIdListLogic) TestGroupIdList(in *chat.TestRequst) (*chat.TestResponse, error) {
	// todo: add your logic here and delete this line

	return &chat.TestResponse{}, nil
}
