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

	UserModel model.UserModel
}

func NewServiceContext(c config.Config) *ServiceContext {

	global.Redis = c.Redis.NewRedis()

	conn := sqlx.NewMysql(c.DataSource)

	return &ServiceContext{
		Config:    c,
		Redis:     global.Redis,
		CustomJwt: middleware.NewCustomJwtMiddleware().Handle,
		UserModel: model.NewUserModel(conn, c.Cache),
	}
}
