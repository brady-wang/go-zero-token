package logic

import (
	"context"
	"errors"
	"github.com/dgrijalva/jwt-go"
	"time"

	"go-zero-token/internal/svc"
	"go-zero-token/internal/types"

	"github.com/tal-tech/go-zero/core/logx"
)

type LoginLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewLoginLogic(ctx context.Context, svcCtx *svc.ServiceContext) LoginLogic {
	return LoginLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *LoginLogic) Login(req types.LoginReq) (*types.LoginReply, error) {

	userName := req.Username
	password := req.Password

	// 模拟验证通过 实际上需要查询数据库的
	if userName == "admin" && password == "password" {
		userId := 1
		// ---start---
		now := time.Now().Unix()
		accessExpire := l.svcCtx.Config.Auth.AccessExpire
		jwtToken, err := l.getJwtToken(l.svcCtx.Config.Auth.AccessSecret, now, l.svcCtx.Config.Auth.AccessExpire, int64(userId))
		if err != nil {
			return nil, errors.New("token生成失败")
		}
		// ---end---

		return &types.LoginReply{
			Id:           1,
			Name:         userName,
			AccessToken:  jwtToken,
			AccessExpire: now + accessExpire,
			RefreshAfter: now + accessExpire/2,
		}, nil
	}


	return &types.LoginReply{}, nil
}


func (l *LoginLogic) getJwtToken(secretKey string, iat, seconds, userId int64) (string, error) {
	claims := make(jwt.MapClaims)
	claims["exp"] = iat + seconds
	claims["iat"] = iat
	claims["userId"] = userId
	token := jwt.New(jwt.SigningMethodHS256)
	token.Claims = claims
	return token.SignedString([]byte(secretKey))
}

