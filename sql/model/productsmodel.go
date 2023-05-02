package model

import (
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ ProductsModel = (*customProductsModel)(nil)

type (
	// ProductsModel is an interface to be customized, add more methods here,
	// and implement the added methods in customProductsModel.
	ProductsModel interface {
		productsModel
		GetPrice(product *Products, quantity float64) int64
	}

	customProductsModel struct {
		*defaultProductsModel
	}
)

// NewProductsModel returns a model for the database table.
func NewProductsModel(conn sqlx.SqlConn, c cache.CacheConf, opts ...cache.Option) ProductsModel {
	return &customProductsModel{
		defaultProductsModel: newProductsModel(conn, c, opts...),
	}
}

func (c *customProductsModel) GetPrice(product *Products, quantity float64) int64 {
	return int64(float64(product.Price.Int64) * product.Discount * quantity)
}
