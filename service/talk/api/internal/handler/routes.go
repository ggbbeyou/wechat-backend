// Code generated by goctl. DO NOT EDIT.
package handler

import (
	"net/http"

	"wechat-backend/service/talk/api/internal/svc"

	"github.com/zeromicro/go-zero/rest"
)

func RegisterHandlers(server *rest.Server, serverCtx *svc.ServiceContext) {
	server.AddRoutes(
		rest.WithMiddlewares(
			[]rest.Middleware{serverCtx.VerifyAuthorityMiddleware},
			[]rest.Route{
				{
					Method:  http.MethodPost,
					Path:    "/publish",
					Handler: PublishHandler(serverCtx),
				},
				{
					Method:  http.MethodGet,
					Path:    "/GetTalkOwnList",
					Handler: GetTalkOwnListHandler(serverCtx),
				},
				{
					Method:  http.MethodGet,
					Path:    "/GetTalkList",
					Handler: GetTalkListHandler(serverCtx),
				},
				{
					Method:  http.MethodPost,
					Path:    "/like",
					Handler: LikeHandler(serverCtx),
				},
				{
					Method:  http.MethodPost,
					Path:    "/comment",
					Handler: CommentHandler(serverCtx),
				},
				{
					Method:  http.MethodGet,
					Path:    "/getcommentlist",
					Handler: GetCommentListHandler(serverCtx),
				},
			}...,
		),
		rest.WithPrefix("/v1/talk"),
	)
}
