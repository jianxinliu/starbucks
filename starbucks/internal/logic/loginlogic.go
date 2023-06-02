package logic

import (
	"context"
	"github.com/golang-jwt/jwt/v4"
	"time"

	"starbucks/starbucks/internal/svc"
	"starbucks/starbucks/internal/types"

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
	resp = new(types.LoginResponse)
	user, err := l.svcCtx.UserModel.FindOneByUserNameAndPassword(l.ctx, req.UserName, req.Password)
	if err != nil {
		resp.RetCode = 404
		resp.Message = "用户未找到"
		return resp, nil
	}

	iat := time.Now().Unix()
	exp := iat + l.svcCtx.Config.UserAuth.Expired

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"userId": user.UserId,
		"exp":    exp,
		"iat":    iat,
	})
	tokenString, err := token.SignedString([]byte(l.svcCtx.Config.UserAuth.Secret))
	if err != nil {
		return nil, err
	}

	return &types.LoginResponse{
		Token: tokenString,
	}, nil
}
