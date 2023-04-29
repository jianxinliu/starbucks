package model

import (
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ ProductGroupModel = (*customProductGroupModel)(nil)

type (
	// ProductGroupModel is an interface to be customized, add more methods here,
	// and implement the added methods in customProductGroupModel.
	ProductGroupModel interface {
		productGroupModel
	}

	customProductGroupModel struct {
		*defaultProductGroupModel
	}
)

// NewProductGroupModel returns a model for the database table.
func NewProductGroupModel(conn sqlx.SqlConn, c cache.CacheConf, opts ...cache.Option) ProductGroupModel {
	return &customProductGroupModel{
		defaultProductGroupModel: newProductGroupModel(conn, c, opts...),
	}
}
