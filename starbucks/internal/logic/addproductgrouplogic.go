package logic

import (
	"context"
	"starbucks/sql/model"
	"starbucks/starbucks/utils"

	"starbucks/starbucks/internal/svc"
	"starbucks/starbucks/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type AddProductGroupLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewAddProductGroupLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddProductGroupLogic {
	return &AddProductGroupLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *AddProductGroupLogic) AddProductGroup(req *types.AddProductGroupReq) (resp *types.BaseResponse, err error) {
	resp = new(types.BaseResponse)
	groupId := utils.NewProductGroupId(func(id string) error {
		_, e := l.svcCtx.ProductGroupModel.FindOneByGroupId(l.ctx, id)
		return e
	})
	productGroup := &model.ProductGroup{
		GroupId:   groupId,
		GroupName: utils.ToSqlNullString(req.Name),
		GroupDesc: utils.ToSqlNullString(req.Desc),
	}
	l.svcCtx.ProductGroupModel.Insert(l.ctx, productGroup)
	return
}
