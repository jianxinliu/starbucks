package model

import (
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ RecipeModel = (*customRecipeModel)(nil)

type (
	// RecipeModel is an interface to be customized, add more methods here,
	// and implement the added methods in customRecipeModel.
	RecipeModel interface {
		recipeModel
	}

	customRecipeModel struct {
		*defaultRecipeModel
	}
)

// NewRecipeModel returns a model for the database table.
func NewRecipeModel(conn sqlx.SqlConn, c cache.CacheConf, opts ...cache.Option) RecipeModel {
	return &customRecipeModel{
		defaultRecipeModel: newRecipeModel(conn, c, opts...),
	}
}
