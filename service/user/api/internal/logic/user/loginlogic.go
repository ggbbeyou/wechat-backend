package user

import (
	"context"
	"strings"
	"wechat-backend/common/response"
	"wechat-backend/common/utils"
	"wechat-backend/service/user/api/internal/svc"
	"wechat-backend/service/user/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type LoginLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewLoginLogic(ctx context.Context, svcCtx *svc.ServiceContext) *LoginLogic {
	return &LoginLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *LoginLogic) Login(req *types.LoginRequest) (resp *types.LoginResponse, err error) {

	user, err := l.svcCtx.UserModel.FindOneByUsername(l.ctx, req.Username)
	if err != nil || user == nil || strings.Compare(user.Password, req.Password) != 0 {
		return nil, response.NewRespError(response.LOGIN_FAIL, response.LOGIN_FAIL_MESSAGE)
	}

	jwtToken, err := utils.GenerateToken(user.Nickname.String, user.Username, l.svcCtx.Config.Auth.AccessSecret, l.svcCtx.Config.Auth.AccessExpire)
	if err != nil {
		return nil, response.NewRespError(response.JWT_FAIL, response.JWT_FAIL_MESSAGE)
	}

	userInfo := types.UserDetail{}
	utils.StrCpy(user, &userInfo)

	return &types.LoginResponse{
		Token:    jwtToken,
		UserInfo: userInfo,
	}, nil
}
