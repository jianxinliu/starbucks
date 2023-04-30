package utils

import "github.com/zeromicro/go-zero/core/stringx"

func NewUserId() string {
	return "user-" + stringx.Randn(6)
}
