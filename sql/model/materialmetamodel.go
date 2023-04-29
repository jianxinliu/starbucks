package model

import (
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ MaterialMetaModel = (*customMaterialMetaModel)(nil)

type (
	// MaterialMetaModel is an interface to be customized, add more methods here,
	// and implement the added methods in customMaterialMetaModel.
	MaterialMetaModel interface {
		materialMetaModel
	}

	customMaterialMetaModel struct {
		*defaultMaterialMetaModel
	}
)

// NewMaterialMetaModel returns a model for the database table.
func NewMaterialMetaModel(conn sqlx.SqlConn, c cache.CacheConf, opts ...cache.Option) MaterialMetaModel {
	return &customMaterialMetaModel{
		defaultMaterialMetaModel: newMaterialMetaModel(conn, c, opts...),
	}
}
