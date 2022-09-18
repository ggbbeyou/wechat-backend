// Code generated by goctl. DO NOT EDIT!
// Source: chat.proto

package server

import (
	"context"

	"go-back/chat/rpc/chat"
	"go-back/chat/rpc/internal/logic"
	"go-back/chat/rpc/internal/svc"
)

type ChatServiceServer struct {
	svcCtx *svc.ServiceContext
	chat.UnimplementedChatServiceServer
}

func NewChatServiceServer(svcCtx *svc.ServiceContext) *ChatServiceServer {
	return &ChatServiceServer{
		svcCtx: svcCtx,
	}
}

func (s *ChatServiceServer) TestGroupIdList(ctx context.Context, in *chat.TestRequst) (*chat.TestResponse, error) {
	l := logic.NewTestGroupIdListLogic(ctx, s.svcCtx)
	return l.TestGroupIdList(in)
}
