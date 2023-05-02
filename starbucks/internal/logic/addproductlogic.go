package logic

import (
	"context"
	"starbucks/sql/model"
	"starbucks/starbucks/utils"

	"starbucks/starbucks/internal/svc"
	"starbucks/starbucks/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type AddProductLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewAddProductLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddProductLogic {
	return &AddProductLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *AddProductLogic) AddProduct(req *types.AddProductReq) (resp *types.BaseResponse, err error) {
	resp = new(types.BaseResponse)
	productId := utils.NewProductId(func(id string) error {
		_, e := l.svcCtx.ProductModel.FindOneByProductId(l.ctx, id)
		return e
	})
	product := &model.Products{
		ProductId:   productId,
		Name:        req.Name,
		Description: utils.ToSqlNullString(req.Description),
		Image:       utils.ToSqlNullString(req.Image),
		GroupId:     req.GroupId,
		Price:       utils.ToSqlNullInt64(int64(req.Price)),
		Discount:    req.Discount,
	}
	l.svcCtx.ProductModel.Insert(l.ctx, product)
	return
}
