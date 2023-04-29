package logic

import (
	"context"

	"starbucks/internal/svc"
	"starbucks/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type OrderDoneLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewOrderDoneLogic(ctx context.Context, svcCtx *svc.ServiceContext) *OrderDoneLogic {
	return &OrderDoneLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *OrderDoneLogic) OrderDone(req *types.OrderDoneReq) (resp *types.BaseResponse, err error) {
	// todo: add your logic here and delete this line

	return
}
