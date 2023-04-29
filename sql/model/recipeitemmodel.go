package model

import (
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ RecipeItemModel = (*customRecipeItemModel)(nil)

type (
	// RecipeItemModel is an interface to be customized, add more methods here,
	// and implement the added methods in customRecipeItemModel.
	RecipeItemModel interface {
		recipeItemModel
	}

	customRecipeItemModel struct {
		*defaultRecipeItemModel
	}
)

// NewRecipeItemModel returns a model for the database table.
func NewRecipeItemModel(conn sqlx.SqlConn, c cache.CacheConf, opts ...cache.Option) RecipeItemModel {
	return &customRecipeItemModel{
		defaultRecipeItemModel: newRecipeItemModel(conn, c, opts...),
	}
}
