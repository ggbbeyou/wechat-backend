package logic

import (
	"context"
	"database/sql"
	"fmt"
	"time"
	"wechat-backend/common/utils"
	"wechat-backend/service/talk/model"

	"wechat-backend/service/talk/api/internal/svc"
	"wechat-backend/service/talk/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type PublishLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewPublishLogic(ctx context.Context, svcCtx *svc.ServiceContext) *PublishLogic {
	return &PublishLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *PublishLogic) Publish(req *types.PublishRequest) (resp *types.EmptyResponse, err error) {
	//机器号  当前数据中心ID号
	snowFlake, _ := utils.NewSnowFlake(0, 0)
	snowId, err := snowFlake.NextId()
	talk := model.Talks{
		Tid:        fmt.Sprintf("%s%d", utils.TALK_PREFIX, snowId),
		Uid:        req.Uid,
		Content:    req.Content,
		Imgs:       sql.NullString{String: req.Imgs, Valid: false},
		CreateTime: sql.NullTime{Time: time.Now(), Valid: false},
		State:      req.State,
	}

	_, err = l.svcCtx.TalkModel.Insert(l.ctx, &talk)
	if err != nil {
		return nil, err
	}
	return &types.EmptyResponse{}, nil
}
