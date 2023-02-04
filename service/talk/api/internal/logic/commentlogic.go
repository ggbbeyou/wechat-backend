package logic

import (
	"context"
	"fmt"
	"time"
	"wechat-backend/common/utils"
	"wechat-backend/service/talk/model"

	"wechat-backend/service/talk/api/internal/svc"
	"wechat-backend/service/talk/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type CommentLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCommentLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CommentLogic {
	return &CommentLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CommentLogic) Comment(req *types.CommentRequest) (resp *types.EmptyResponse, err error) {
	//机器号  当前数据中心ID号
	snowFlake, _ := utils.NewSnowFlake(0, 0)
	snowId, err := snowFlake.NextId()

	comment := model.Comments{
		Cid:        fmt.Sprintf("%s%d", utils.COMMENT_PREFIX, snowId),
		ParentId:   req.ParentId,
		Uid:        req.Uid,
		Tid:        req.Tid,
		Content:    req.Content,
		CreateTime: time.Now(),
	}
	_, err = l.svcCtx.CommentModel.Insert(l.ctx, &comment)
	if err != nil {
		return nil, err
	}
	return &types.EmptyResponse{}, nil
}
