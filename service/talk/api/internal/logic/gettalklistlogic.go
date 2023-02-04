package logic

import (
	"context"
	"github.com/zeromicro/go-zero/core/stores/sqlc"
	"wechat-backend/common/utils"
	"wechat-backend/service/user/rpc/user"

	"wechat-backend/service/talk/api/internal/svc"
	"wechat-backend/service/talk/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetTalkListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetTalkListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetTalkListLogic {
	return &GetTalkListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetTalkListLogic) GetTalkList(req *types.GetTalkListRequest) (resp *types.GetTalkListResponse, err error) {
	var talkList []types.TalkListDetail
	//1.拿到好友列表
	UIdListResponse, err := l.svcCtx.UserRpcService.GetFriendIdList(l.ctx, &user.GetFriendIdListRequest{Uid: req.Uid})
	if err != nil {
		return nil, err
	}
	friendIdList := UIdListResponse.UidArr
	friendIdList = append(friendIdList, req.Uid)
	//2.查询
	talkListDetails, err := l.svcCtx.TalkModel.SelectTalkList(l.ctx, friendIdList)
	if err != nil && err != sqlc.ErrNotFound {
		return nil, err
	}
	for _, talk := range talkListDetails {
		var t types.TalkListDetail
		utils.StrCpy(&t, talk)
		talkList = append(talkList, t)
	}
	return &types.GetTalkListResponse{
		TalkList: talkList,
	}, nil
}
