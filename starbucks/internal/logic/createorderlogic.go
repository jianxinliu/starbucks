package logic

import (
	"context"
	"fmt"
	"github.com/hibiken/asynq"
	"starbucks/constants"
	"starbucks/sql/model"
	"starbucks/starbucks/utils"
	"time"

	"starbucks/starbucks/internal/svc"
	"starbucks/starbucks/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type CreateOrderLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCreateOrderLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateOrderLogic {
	return &CreateOrderLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CreateOrderLogic) CreateOrder(req *types.CreateOrderRequest) (resp *types.CreateOrderResponse, err error) {
	resp = new(types.CreateOrderResponse)

	userId := l.ctx.Value("userId").(string)
	order := &model.Order{
		OrderId:    l.svcCtx.OrderModel.NewOrderId(l.ctx),
		ProductId:  req.ProductId,
		Status:     constants.STATUS_ORDER_INIT,
		UserId:     userId,
		CreateTime: time.Now(),
		OrderType:  req.OrderType,
		Quantity:   float64(req.Quantity),
		TrxNo:      utils.ToSqlNullString(l.svcCtx.OrderModel.NewTrxNo(l.ctx)),
	}
	_, err = l.svcCtx.OrderModel.Insert(l.ctx, order)
	if err != nil {
		resp.RetCode = constants.ERR_ORDER_CREATE
		resp.Message = err.Error()
		return
	}
	resp.OrderId = order.OrderId
	// 设置一个十分钟的定时任务，必须在十分钟内支付，否则订单失效

	asynqClient := asynq.NewClient(asynq.RedisClientOpt{
		Addr:     l.svcCtx.Config.Redis.Host,
		Username: "",
		Password: "",
		DB:       0,
	})

	asynqClient.Enqueue(
		asynq.NewTask(constants.Q_ORDER_EXPIRED, []byte(fmt.Sprintf("%s|%s", order.OrderId, order.OrderType))),
		asynq.ProcessAt(time.Now().Add(constants.ORDER_EXPIRED_MINUTES*time.Minute)),
		asynq.Queue(constants.ORDER_Q_NAME))

	return
}
