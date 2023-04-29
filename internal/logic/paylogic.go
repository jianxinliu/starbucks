package logic

import (
	"context"

	"starbucks/internal/svc"
	"starbucks/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type PayLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewPayLogic(ctx context.Context, svcCtx *svc.ServiceContext) *PayLogic {
	return &PayLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *PayLogic) Pay(req *types.PayReq) (resp *types.BaseResponse, err error) {
	// todo: add your logic here and delete this line

	return
}
