package model

import (
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ MaterialModel = (*customMaterialModel)(nil)

type (
	// MaterialModel is an interface to be customized, add more methods here,
	// and implement the added methods in customMaterialModel.
	MaterialModel interface {
		materialModel
	}

	customMaterialModel struct {
		*defaultMaterialModel
	}
)

// NewMaterialModel returns a model for the database table.
func NewMaterialModel(conn sqlx.SqlConn, c cache.CacheConf, opts ...cache.Option) MaterialModel {
	return &customMaterialModel{
		defaultMaterialModel: newMaterialModel(conn, c, opts...),
	}
}
