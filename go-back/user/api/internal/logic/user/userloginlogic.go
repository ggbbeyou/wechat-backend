package user

import (
	"context"
	"go-back/common"
	"strings"

	"go-back/user/api/internal/svc"
	"go-back/user/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UserLoginLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUserLoginLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserLoginLogic {
	return &UserLoginLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UserLoginLogic) UserLogin(req *types.LoginRequest) (resp *types.LoginResponse, err error) {
	//查询数据库校验用户名密码
	username := req.UserName
	password := req.Password
	user, _ := l.svcCtx.UserModel.FindOneByUsername(l.ctx, username)
	resp = &types.LoginResponse{}
	//密码或者用户名错误
	if user == nil || strings.Compare(user.Password, password) != 0 {
		return nil, common.NewCodeError(common.LOGIN_FAIL, common.LOGIN_FAIL_MSG)
	}
	//生成jwt
	token, err := common.CreateToken(user.Uid, username)
	if err != nil {
		return nil, common.NewCodeError(common.JWT_FAIL, common.JWT_FAIL_MSG)
	}
	resp.Token = token
	userInfo := types.UserInfo{}
	common.ObjCopy(user, &userInfo.BaseInfo)
	resp.Data = userInfo
	return resp, err
}
