package utils

import "github.com/zeromicro/go-zero/core/stringx"

func NewUserId() string {
	return "user-" + stringx.Randn(6)
}

func NewOrderId() string {
	return "order-" + stringx.Randn(6)
}

func NewTrxNo() string {
	return "order-" + stringx.Randn(16)
}
