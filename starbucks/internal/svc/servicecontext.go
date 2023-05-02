package svc

import (
	"github.com/zeromicro/go-zero/core/stores/redis"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"github.com/zeromicro/go-zero/rest"
	"starbucks/sql/model"
	"starbucks/starbucks/global"
	"starbucks/starbucks/internal/config"
	"starbucks/starbucks/internal/middleware"
)

type ServiceContext struct {
	Config    config.Config
	Redis     *redis.Redis
	CustomJwt rest.Middleware

	DbConn sqlx.SqlConn

	UserModel         model.UserModel
	OrderModel        model.OrderModel
	ProductModel      model.ProductsModel
	ProductGroupModel model.ProductGroupModel
	PaymentModel      model.PaymentModel
	WalletModel       model.WalletModel
}

func NewServiceContext(c config.Config) *ServiceContext {

	global.Redis = c.Redis.NewRedis()

	conn := sqlx.NewMysql(c.DataSource)

	return &ServiceContext{
		Config:            c,
		Redis:             global.Redis,
		CustomJwt:         middleware.NewCustomJwtMiddleware().Handle,
		DbConn:            conn,
		UserModel:         model.NewUserModel(conn, c.Cache),
		OrderModel:        model.NewOrderModel(conn, c.Cache),
		ProductModel:      model.NewProductsModel(conn, c.Cache),
		ProductGroupModel: model.NewProductGroupModel(conn, c.Cache),
		PaymentModel:      model.NewPaymentModel(conn, c.Cache),
		WalletModel:       model.NewWalletModel(conn, c.Cache),
	}
}
