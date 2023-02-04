package logic

import (
	"context"
	"github.com/zeromicro/go-zero/core/stores/sqlc"
	"wechat-backend/common/utils"

	"wechat-backend/service/talk/api/internal/svc"
	"wechat-backend/service/talk/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetTalkOwnListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetTalkOwnListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetTalkOwnListLogic {
	return &GetTalkOwnListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetTalkOwnListLogic) GetTalkOwnList(req *types.GetTalkListRequest) (resp *types.GetTalkListResponse, err error) {

	ownTalkList, err := l.svcCtx.TalkModel.SelectOwnTalkList(l.ctx, req.Uid)
	if err != nil && err != sqlc.ErrNotFound {
		return nil, err
	}
	var talkList []types.TalkListDetail
	for _, talk := range ownTalkList {
		var t types.TalkListDetail
		utils.StrCpy(&t, talk)
		talkList = append(talkList, t)
	}
	return &types.GetTalkListResponse{
		TalkList: talkList,
	}, nil
}
