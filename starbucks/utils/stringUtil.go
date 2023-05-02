package utils

import (
	"github.com/zeromicro/go-zero/core/stores/sqlc"
	"github.com/zeromicro/go-zero/core/stringx"
)

func NewUserId(fn IdGeneratorFn) string {
	return NewIdBy("user-", 6, fn)
}

func NewOrderId(fn IdGeneratorFn) string {
	return NewIdBy("order-", 6, fn)
}

func NewTrxNo(fn IdGeneratorFn) string {
	return NewIdBy("trx-", 16, fn)
}

func NewProductGroupId(fn IdGeneratorFn) string {
	return NewIdBy("pg-", 6, fn)
}

func NewProductId(fn IdGeneratorFn) string {
	return NewIdBy("product-", 6, fn)
}

type IdGeneratorFn func(id string) error

// NewIdBy 通过碰撞的方式创建 Id
// prefix string ID 的前缀
// rndLen int id 除去前端的随机字串长度
// fn func(curId string) bool 一个检测当前生成的 ID 是否重复的函数，返回数据库查询结果的 error 对象
func NewIdBy(prefix string, rndLen int, fn IdGeneratorFn) string {
	id := prefix + stringx.Randn(rndLen)
	for i := 0; i < 10; i++ {
		err := fn(id)
		if err != nil && err == sqlc.ErrNotFound {
			break
		}
		id = prefix + stringx.Randn(rndLen)
	}
	return id
}
