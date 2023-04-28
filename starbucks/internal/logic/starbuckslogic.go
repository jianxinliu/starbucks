package logic

import (
	"context"

	"starbucks/starbucks/internal/svc"
	"starbucks/starbucks/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type StarbucksLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewStarbucksLogic(ctx context.Context, svcCtx *svc.ServiceContext) *StarbucksLogic {
	return &StarbucksLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *StarbucksLogic) Starbucks(req *types.Request) (resp *types.Response, err error) {
	// todo: add your logic here and delete this line

	return
}
