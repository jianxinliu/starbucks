// Code generated by goctl. DO NOT EDIT.

package model

import (
	"context"
	"database/sql"
	"fmt"
	"strings"

	"github.com/zeromicro/go-zero/core/stores/builder"
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlc"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"github.com/zeromicro/go-zero/core/stringx"
)

var (
	walletFieldNames          = builder.RawFieldNames(&Wallet{})
	walletRows                = strings.Join(walletFieldNames, ",")
	walletRowsExpectAutoSet   = strings.Join(stringx.Remove(walletFieldNames, "`id`", "`create_at`", "`create_time`", "`created_at`", "`update_at`", "`update_time`", "`updated_at`"), ",")
	walletRowsWithPlaceHolder = strings.Join(stringx.Remove(walletFieldNames, "`id`", "`create_at`", "`create_time`", "`created_at`", "`update_at`", "`update_time`", "`updated_at`"), "=?,") + "=?"

	cacheStarbucksWalletIdPrefix     = "cache:starbucks:wallet:id:"
	cacheStarbucksWalletUserIdPrefix = "cache:starbucks:wallet:userId:"
)

type (
	walletModel interface {
		Insert(ctx context.Context, data *Wallet) (sql.Result, error)
		FindOne(ctx context.Context, id int64) (*Wallet, error)
		FindOneByUserId(ctx context.Context, userId string) (*Wallet, error)
		Update(ctx context.Context, data *Wallet) error
		Delete(ctx context.Context, id int64) error
	}

	defaultWalletModel struct {
		sqlc.CachedConn
		table string
	}

	Wallet struct {
		Id        int64        `db:"id"`
		CreatedAt sql.NullTime `db:"created_at"`
		UpdatedAt sql.NullTime `db:"updated_at"`
		DeletedAt sql.NullTime `db:"deleted_at"`
		UserId    string       `db:"user_id"`
		Balance   int64        `db:"balance"` // 用户余额，单位： 分
	}
)

func newWalletModel(conn sqlx.SqlConn, c cache.CacheConf, opts ...cache.Option) *defaultWalletModel {
	return &defaultWalletModel{
		CachedConn: sqlc.NewConn(conn, c, opts...),
		table:      "`wallet`",
	}
}

func (m *defaultWalletModel) Delete(ctx context.Context, id int64) error {
	data, err := m.FindOne(ctx, id)
	if err != nil {
		return err
	}

	starbucksWalletIdKey := fmt.Sprintf("%s%v", cacheStarbucksWalletIdPrefix, id)
	starbucksWalletUserIdKey := fmt.Sprintf("%s%v", cacheStarbucksWalletUserIdPrefix, data.UserId)
	_, err = m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("delete from %s where `id` = ?", m.table)
		return conn.ExecCtx(ctx, query, id)
	}, starbucksWalletIdKey, starbucksWalletUserIdKey)
	return err
}

func (m *defaultWalletModel) FindOne(ctx context.Context, id int64) (*Wallet, error) {
	starbucksWalletIdKey := fmt.Sprintf("%s%v", cacheStarbucksWalletIdPrefix, id)
	var resp Wallet
	err := m.QueryRowCtx(ctx, &resp, starbucksWalletIdKey, func(ctx context.Context, conn sqlx.SqlConn, v any) error {
		query := fmt.Sprintf("select %s from %s where `id` = ? limit 1", walletRows, m.table)
		return conn.QueryRowCtx(ctx, v, query, id)
	})
	switch err {
	case nil:
		return &resp, nil
	case sqlc.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}

func (m *defaultWalletModel) FindOneByUserId(ctx context.Context, userId string) (*Wallet, error) {
	starbucksWalletUserIdKey := fmt.Sprintf("%s%v", cacheStarbucksWalletUserIdPrefix, userId)
	var resp Wallet
	err := m.QueryRowIndexCtx(ctx, &resp, starbucksWalletUserIdKey, m.formatPrimary, func(ctx context.Context, conn sqlx.SqlConn, v any) (i any, e error) {
		query := fmt.Sprintf("select %s from %s where `user_id` = ? limit 1", walletRows, m.table)
		if err := conn.QueryRowCtx(ctx, &resp, query, userId); err != nil {
			return nil, err
		}
		return resp.Id, nil
	}, m.queryPrimary)
	switch err {
	case nil:
		return &resp, nil
	case sqlc.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}

func (m *defaultWalletModel) Insert(ctx context.Context, data *Wallet) (sql.Result, error) {
	starbucksWalletIdKey := fmt.Sprintf("%s%v", cacheStarbucksWalletIdPrefix, data.Id)
	starbucksWalletUserIdKey := fmt.Sprintf("%s%v", cacheStarbucksWalletUserIdPrefix, data.UserId)
	ret, err := m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("insert into %s (%s) values (?, ?, ?)", m.table, walletRowsExpectAutoSet)
		return conn.ExecCtx(ctx, query, data.DeletedAt, data.UserId, data.Balance)
	}, starbucksWalletIdKey, starbucksWalletUserIdKey)
	return ret, err
}

func (m *defaultWalletModel) Update(ctx context.Context, newData *Wallet) error {
	data, err := m.FindOne(ctx, newData.Id)
	if err != nil {
		return err
	}

	starbucksWalletIdKey := fmt.Sprintf("%s%v", cacheStarbucksWalletIdPrefix, data.Id)
	starbucksWalletUserIdKey := fmt.Sprintf("%s%v", cacheStarbucksWalletUserIdPrefix, data.UserId)
	_, err = m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("update %s set %s where `id` = ?", m.table, walletRowsWithPlaceHolder)
		return conn.ExecCtx(ctx, query, newData.DeletedAt, newData.UserId, newData.Balance, newData.Id)
	}, starbucksWalletIdKey, starbucksWalletUserIdKey)
	return err
}

func (m *defaultWalletModel) formatPrimary(primary any) string {
	return fmt.Sprintf("%s%v", cacheStarbucksWalletIdPrefix, primary)
}

func (m *defaultWalletModel) queryPrimary(ctx context.Context, conn sqlx.SqlConn, v, primary any) error {
	query := fmt.Sprintf("select %s from %s where `id` = ? limit 1", walletRows, m.table)
	return conn.QueryRowCtx(ctx, v, query, primary)
}

func (m *defaultWalletModel) tableName() string {
	return m.table
}