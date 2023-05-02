package logic

import (
	"context"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"starbucks/constants"
	"starbucks/sql/model"
	"time"

	"starbucks/starbucks/internal/svc"
	"starbucks/starbucks/internal/types"

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
	resp = new(types.BaseResponse)

	userId := l.ctx.Value("userId").(string)
	orderId := req.OrderId
	order, err := l.svcCtx.OrderModel.FindOneByOrderId(l.ctx, orderId)
	if err != nil {
		resp.Message = err.Error()
		resp.RetCode = constants.ERR_DB
		return resp, nil
	}
	switch order.Status {
	case constants.STATUS_ORDER_NOT_PAY:
		resp.Message = "订单已过期，不能支付"
		resp.RetCode = constants.ERR_PAY_NOT_ALLOWED
		return resp, nil
	case constants.STATUS_ORDER_PAYED:
		resp.Message = "订单已支付，不能重复支付"
		resp.RetCode = constants.ERR_PAY_PAYED
		return resp, nil
	case constants.STATUS_ORDER_DONE:
		resp.Message = "订单已完成，无需再支付"
		resp.RetCode = constants.ERR_PAY_NOT_ALLOWED
		return resp, nil
	}

	wallet, err := l.svcCtx.WalletModel.FindOneByUserId(l.ctx, userId)
	if err != nil {
		resp.Message = err.Error()
		resp.RetCode = constants.ERR_DB
		markOrderUnPay(l.ctx, order, l.svcCtx.OrderModel)
		return resp, nil
	}

	// pay
	product, err := l.svcCtx.ProductModel.FindOneByProductId(l.ctx, order.ProductId)
	if err != nil {
		resp.Message = err.Error()
		resp.RetCode = constants.ERR_DB
		markOrderUnPay(l.ctx, order, l.svcCtx.OrderModel)
		return resp, nil
	}

	price := l.svcCtx.ProductModel.GetPrice(product, order.Quantity)

	if wallet.Balance < price {
		resp.Message = "用户余额不足"
		resp.RetCode = constants.ERR_PAY_BALANCE_Insufficient
		markOrderUnPay(l.ctx, order, l.svcCtx.OrderModel)
		return resp, nil
	}

	payment := &model.Payment{
		UserId:      userId,
		OrderId:     orderId,
		PayedAmount: price,
		PayedTime:   time.Now(),
	}

	// 保存订单信息和
	err = l.svcCtx.DbConn.TransactCtx(l.ctx, func(ctx context.Context, session sqlx.Session) error {
		order.Status = constants.STATUS_ORDER_PAYED
		err := l.svcCtx.OrderModel.Update(l.ctx, order)
		if err != nil {
			return err
		}

		_, err = l.svcCtx.PaymentModel.Insert(l.ctx, payment)
		if err != nil {
			return err
		}
		return nil
	})

	if err != nil {
		resp.Message = err.Error()
		resp.RetCode = constants.ERR_DB
		markOrderUnPay(l.ctx, order, l.svcCtx.OrderModel)
		return resp, nil
	}
	return
}

func markOrderUnPay(ctx context.Context, order *model.Order, orderModel model.OrderModel) {
	order.Status = constants.STATUS_ORDER_NOT_PAY
	orderModel.Update(ctx, order)
}
