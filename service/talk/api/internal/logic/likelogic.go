package logic

import (
	"context"
	"fmt"
	"strings"
	"wechat-backend/common/utils"

	"wechat-backend/service/talk/api/internal/svc"
	"wechat-backend/service/talk/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type LikeLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewLikeLogic(ctx context.Context, svcCtx *svc.ServiceContext) *LikeLogic {
	return &LikeLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *LikeLogic) Like(req *types.LiskeRequest) (resp *types.EmptyResponse, err error) {
	//1.判断是否点赞了
	//1.1 没有的话，写入redis
	//1.2 点赞了的话，从redis中取消
	//结构：文章id::用户id -> 1/0(1.点赞，0取消点赞)
	//2.定时任务更新缓存到数据库，并删除缓存（延时双删）
	field := fmt.Sprintf("%s::%d", req.Tid, req.Uid)
	client := l.svcCtx.Redis

	exists, err := client.Exists(utils.LIKE_KEY)
	if err != nil {
		logx.Error(err)
		return nil, err
	}
	newLikeState := utils.DEFAULT_VALUE
	if exists {
		isLike, err := client.Hget(utils.LIKE_KEY, field)
		if err != nil {
			logx.Error(err)
			return nil, err
		}
		//点赞过, 取消点赞
		if strings.Compare(isLike, utils.DEFAULT_VALUE) == 1 {
			newLikeState = utils.DEFAULT_NOT_VALUE
		}
	}
	client.Hset(utils.LIKE_KEY, field, newLikeState)
	return &types.EmptyResponse{}, nil
}
