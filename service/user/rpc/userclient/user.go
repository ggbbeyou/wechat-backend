// Code generated by goctl. DO NOT EDIT!
// Source: user.proto

package userclient

import (
	"context"

	"wechat-backend/service/user/rpc/user"

	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
)

type (
	Empty                  = user.Empty
	GetFriendIdListRequest = user.GetFriendIdListRequest
	UIdListResponse        = user.UIdListResponse

	User interface {
		// 获取所有用户
		GetAllUserIdList(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*UIdListResponse, error)
		// 获取所有好友
		GetFriendIdList(ctx context.Context, in *GetFriendIdListRequest, opts ...grpc.CallOption) (*UIdListResponse, error)
	}

	defaultUser struct {
		cli zrpc.Client
	}
)

func NewUser(cli zrpc.Client) User {
	return &defaultUser{
		cli: cli,
	}
}

// 获取所有用户
func (m *defaultUser) GetAllUserIdList(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*UIdListResponse, error) {
	client := user.NewUserClient(m.cli.Conn())
	return client.GetAllUserIdList(ctx, in, opts...)
}

// 获取所有好友
func (m *defaultUser) GetFriendIdList(ctx context.Context, in *GetFriendIdListRequest, opts ...grpc.CallOption) (*UIdListResponse, error) {
	client := user.NewUserClient(m.cli.Conn())
	return client.GetFriendIdList(ctx, in, opts...)
}
