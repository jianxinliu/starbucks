package config

import (
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/redis"
	"github.com/zeromicro/go-zero/rest"
)

type Config struct {
	rest.RestConf

	Redis redis.RedisConf
	Cache cache.CacheConf

	DataSource string

	UserAuth struct {
		Secret  string
		Expired int64
	}
}
