package logic

import (
	"context"

	"go-zero-token/internal/svc"
	"go-zero-token/internal/types"

	"github.com/tal-tech/go-zero/core/logx"
)

type UserInfoLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUserInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) UserInfoLogic {
	return UserInfoLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UserInfoLogic) UserInfo(req types.UserInfoReq) (*types.UserInfoResp, error) {


	return &types.UserInfoResp{
		Id:   "1",
		Name: "admin",
	}, nil
}
