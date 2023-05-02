package server

import (
	"context"
	"log"
	"math"
	"math/rand"
	"starbucks/constants"
	"starbucks/starbucks/internal/svc"
	"strings"
	"time"

	"github.com/hibiken/asynq"
	"github.com/zeromicro/go-zero/core/stores/redis"
)

type AsynqServer struct {
	svcCtx *svc.ServiceContext
	redis  *redis.Redis
}

var asynqServer *AsynqServer

func InitAsynqServer(svcCtx *svc.ServiceContext) {
	asynqServer = &AsynqServer{
		svcCtx: svcCtx,
		redis:  svcCtx.Config.Redis.NewRedis(),
	}

	srv := asynq.NewServer(
		asynq.RedisClientOpt{
			Addr: svcCtx.Config.Redis.Host,
			DB:   0,
		},
		asynq.Config{
			Concurrency: 20,
			// LogLevel:    asynq.DebugLevel,
			Queues: map[string]int{
				constants.ORDER_Q_NAME: 10,
			},
			RetryDelayFunc: func(n int, e error, t *asynq.Task) time.Duration {

				r := rand.New(rand.NewSource(time.Now().UnixNano()))
				// Formula taken from https://github.com/mperham/sidekiq.
				s := int(math.Pow(float64(n), 4)) + 15 + (r.Intn(30) * (n + 1))
				return time.Duration(s) * time.Second

			},
		},
	)

	mux := asynq.NewServeMux()
	// 任务执行时的handle
	mux.HandleFunc(constants.Q_ORDER_EXPIRED, asynqServer.HandleOrderExpired)

	if err := srv.Start(mux); err != nil {
		log.Fatalf("could not start server: %v", err)
	}
}

func (a *AsynqServer) HandleOrderExpired(ctx context.Context, t *asynq.Task) error {
	payload := strings.Split(string(t.Payload()), "|")
	if len(payload) != 2 {
		return nil
	}
	// 第一个参数为 orderId, 第二个参数为 orderType
	orderId := payload[0]

	order, err := a.svcCtx.OrderModel.FindOneByOrderId(ctx, orderId)
	if err != nil {
		return err
	}
	if order.Status == constants.STATUS_ORDER_INIT {
		// 如果还是初始化的状态，说明没支付，也没消费成功，那就改为未支付
		order.Status = constants.STATUS_ORDER_NOT_PAY
		a.svcCtx.OrderModel.Update(ctx, order)
	}
	return nil
}
