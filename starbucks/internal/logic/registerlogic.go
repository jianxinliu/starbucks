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
	user := &model.User{
		UserName: req.UserName,
		Password: req.Password,
		UserId:   utils.NewUserId(),
		UserType: global.Normal.String(),
	}
	_, err = l.svcCtx.UserModel.Insert(l.ctx, user)
	if err != nil {
		resp.Message = err.Error()
		return resp, nil
	}
	return
}
