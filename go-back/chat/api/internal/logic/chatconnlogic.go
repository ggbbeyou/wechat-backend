package logic

import (
	"context"

	"github.com/zeromicro/go-zero/core/logx"
	"go-back/chat/api/internal/svc"
)

type ChatConnLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewChatConnLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ChatConnLogic {
	return &ChatConnLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ChatConnLogic) ChatConn() error {

	return nil
}
