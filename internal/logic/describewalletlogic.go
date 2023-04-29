package logic

import (
	"context"

	"starbucks/internal/svc"
	"starbucks/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type DescribeWalletLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewDescribeWalletLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DescribeWalletLogic {
	return &DescribeWalletLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DescribeWalletLogic) DescribeWallet() (resp *types.WalletDescribeResp, err error) {
	// todo: add your logic here and delete this line

	return
}
