package record

import (
	"context"
	"github.com/zeromicro/go-zero/core/stores/sqlc"
	"wechat-backend/common/utils"

	"wechat-backend/service/communicate/api/internal/svc"
	"wechat-backend/service/communicate/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type RecordListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewRecordListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RecordListLogic {
	return &RecordListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *RecordListLogic) RecordList(req *types.RecordListRequest) (resp *types.RecordListResponse, err error) {
	recodeList, err := l.svcCtx.RecodesModel.SelectRecordList(l.ctx, req.Uid, req.FriendUid)
	if err != nil && err != sqlc.ErrNotFound {
		return nil, err
	}

	var messages []types.MessageInfo

	for _, recode := range recodeList {
		m := types.MessageInfo{}
		utils.StrCpy(recode, &m)
		messages = append(messages, m)
	}

	return &types.RecordListResponse{
		RecordList: messages,
	}, err
}
