package logic

import (
	"context"
	"github.com/zeromicro/go-zero/core/stores/sqlc"
	"wechat-backend/common/utils"

	"wechat-backend/service/talk/api/internal/svc"
	"wechat-backend/service/talk/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetCommentListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetCommentListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetCommentListLogic {
	return &GetCommentListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetCommentListLogic) GetCommentList(req *types.GetCommentListRequest) (resp *types.GetCommentListResponse, err error) {
	// todo: add your logic here and delete this line
	commentList, err := l.svcCtx.CommentModel.SelectCommentList(l.ctx, req.Tid)
	if err != nil && err != sqlc.ErrNotFound {
		return nil, err
	}
	var CommentList []types.CommentDetail
	for _, comment := range commentList {
		var t types.CommentDetail
		utils.StrCpy(&t, comment)
		CommentList = append(CommentList, t)
	}

	return &types.GetCommentListResponse{
		CommentList: commentRelation(CommentList, utils.DEFAULT_NOT_VALUE),
	}, nil
}

//通过递归构建评论列表层级关系
func commentRelation(commentList []types.CommentDetail, parentId string) []types.CommentDetail {
	var newCommentList []types.CommentDetail

	for i := 0; i < len(commentList); i++ {
		if parentId == commentList[i].ParentId {
			newCommentList = append(newCommentList, commentList[i])
			commentList[i].ChildCommentList = commentRelation(commentList, commentList[i].Cid)
		}
	}
	return newCommentList
}
