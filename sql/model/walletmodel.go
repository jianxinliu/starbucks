package model

import (
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ WalletModel = (*customWalletModel)(nil)

type (
	// WalletModel is an interface to be customized, add more methods here,
	// and implement the added methods in customWalletModel.
	WalletModel interface {
		walletModel
	}

	customWalletModel struct {
		*defaultWalletModel
	}
)

// NewWalletModel returns a model for the database table.
func NewWalletModel(conn sqlx.SqlConn, c cache.CacheConf, opts ...cache.Option) WalletModel {
	return &customWalletModel{
		defaultWalletModel: newWalletModel(conn, c, opts...),
	}
}
