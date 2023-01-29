package logic

import (
	"context"

	"wechat-backend/service/talk/internal/svc"
	"wechat-backend/service/talk/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type TalkLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewTalkLogic(ctx context.Context, svcCtx *svc.ServiceContext) *TalkLogic {
	return &TalkLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *TalkLogic) Talk(req *types.Request) (resp *types.Response, err error) {
	// todo: add your logic here and delete this line

	return
}
