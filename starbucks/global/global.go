package global

import (
	"context"
	"github.com/zeromicro/go-zero/core/stores/redis"
	"starbucks/starbucks/internal/config"
)

var Config *config.Config

var Redis *redis.Redis

var Ctx *context.Context

type UserType string

const (
	Normal UserType = "normal"
	Vip    UserType = "vip"
)

func ParseUserType(s string) UserType {
	switch s {
	case "normal":
		return Normal
	case "vip":
		return Vip
	}
	return Normal
}

func (ut UserType) String() string {
	switch ut {
	case Normal:
		return "normal"
	case Vip:
		return "vip"
	}
	return Normal.String()
}
