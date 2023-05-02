package logic

import (
	"context"
	"github.com/zeromicro/go-zero/core/logx"
	"starbucks/sql/model"
	"starbucks/starbucks/global"
	"starbucks/starbucks/internal/svc"
	"starbucks/starbucks/internal/types"
	"starbucks/starbucks/utils"
)

type RegisterLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewRegisterLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RegisterLogic {
	return &RegisterLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *RegisterLogic) Register(req *types.RegisterRequest) (resp *types.BaseResponse, err error) {
	resp = new(types.BaseResponse)
	userId := utils.NewUserId(func(id string) error {
		_, e := l.svcCtx.UserModel.FindOneByUserId(l.ctx, id)
		return e
	})
	user := &model.User{
		UserName: req.UserName,
		Password: req.Password,
		UserId:   userId,
		UserType: global.Normal.String(),
	}
	_, err = l.svcCtx.UserModel.Insert(l.ctx, user)
	if err != nil {
		resp.Message = err.Error()
		return resp, nil
	}

	wallet := &model.Wallet{
		UserId:  userId,
		Balance: 0,
	}
	l.svcCtx.WalletModel.Insert(l.ctx, wallet)
	return
}
