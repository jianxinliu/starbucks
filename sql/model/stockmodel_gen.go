// Code generated by goctl. DO NOT EDIT.

package model

import (
	"context"
	"database/sql"
	"fmt"
	"strings"
	"time"

	"github.com/zeromicro/go-zero/core/stores/builder"
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlc"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"github.com/zeromicro/go-zero/core/stringx"
)

var (
	stockFieldNames          = builder.RawFieldNames(&Stock{})
	stockRows                = strings.Join(stockFieldNames, ",")
	stockRowsExpectAutoSet   = strings.Join(stringx.Remove(stockFieldNames, "`id`", "`create_at`", "`create_time`", "`created_at`", "`update_at`", "`update_time`", "`updated_at`"), ",")
	stockRowsWithPlaceHolder = strings.Join(stringx.Remove(stockFieldNames, "`id`", "`create_at`", "`create_time`", "`created_at`", "`update_at`", "`update_time`", "`updated_at`"), "=?,") + "=?"

	cacheStarbucksStockIdPrefix      = "cache:starbucks:stock:id:"
	cacheStarbucksStockStockIdPrefix = "cache:starbucks:stock:stockId:"
)

type (
	stockModel interface {
		Insert(ctx context.Context, data *Stock) (sql.Result, error)
		FindOne(ctx context.Context, id int64) (*Stock, error)
		FindOneByStockId(ctx context.Context, stockId string) (*Stock, error)
		Update(ctx context.Context, data *Stock) error
		Delete(ctx context.Context, id int64) error
	}

	defaultStockModel struct {
		sqlc.CachedConn
		table string
	}

	Stock struct {
		Id         int64     `db:"id"`
		StockId    string    `db:"stock_id"`    // 库存 id
		MaterialId string    `db:"material_id"` // 存储的哪种原料
		Amount     int64     `db:"amount"`      // 入库数量
		Count      int64     `db:"count"`       // 当前数量
		InTime     time.Time `db:"in_time"`     // 原料进库时间
	}
)

func newStockModel(conn sqlx.SqlConn, c cache.CacheConf, opts ...cache.Option) *defaultStockModel {
	return &defaultStockModel{
		CachedConn: sqlc.NewConn(conn, c, opts...),
		table:      "`stock`",
	}
}

func (m *defaultStockModel) Delete(ctx context.Context, id int64) error {
	data, err := m.FindOne(ctx, id)
	if err != nil {
		return err
	}

	starbucksStockIdKey := fmt.Sprintf("%s%v", cacheStarbucksStockIdPrefix, id)
	starbucksStockStockIdKey := fmt.Sprintf("%s%v", cacheStarbucksStockStockIdPrefix, data.StockId)
	_, err = m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("delete from %s where `id` = ?", m.table)
		return conn.ExecCtx(ctx, query, id)
	}, starbucksStockIdKey, starbucksStockStockIdKey)
	return err
}

func (m *defaultStockModel) FindOne(ctx context.Context, id int64) (*Stock, error) {
	starbucksStockIdKey := fmt.Sprintf("%s%v", cacheStarbucksStockIdPrefix, id)
	var resp Stock
	err := m.QueryRowCtx(ctx, &resp, starbucksStockIdKey, func(ctx context.Context, conn sqlx.SqlConn, v any) error {
		query := fmt.Sprintf("select %s from %s where `id` = ? limit 1", stockRows, m.table)
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

func (m *defaultStockModel) FindOneByStockId(ctx context.Context, stockId string) (*Stock, error) {
	starbucksStockStockIdKey := fmt.Sprintf("%s%v", cacheStarbucksStockStockIdPrefix, stockId)
	var resp Stock
	err := m.QueryRowIndexCtx(ctx, &resp, starbucksStockStockIdKey, m.formatPrimary, func(ctx context.Context, conn sqlx.SqlConn, v any) (i any, e error) {
		query := fmt.Sprintf("select %s from %s where `stock_id` = ? limit 1", stockRows, m.table)
		if err := conn.QueryRowCtx(ctx, &resp, query, stockId); err != nil {
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

func (m *defaultStockModel) Insert(ctx context.Context, data *Stock) (sql.Result, error) {
	starbucksStockIdKey := fmt.Sprintf("%s%v", cacheStarbucksStockIdPrefix, data.Id)
	starbucksStockStockIdKey := fmt.Sprintf("%s%v", cacheStarbucksStockStockIdPrefix, data.StockId)
	ret, err := m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("insert into %s (%s) values (?, ?, ?, ?, ?)", m.table, stockRowsExpectAutoSet)
		return conn.ExecCtx(ctx, query, data.StockId, data.MaterialId, data.Amount, data.Count, data.InTime)
	}, starbucksStockIdKey, starbucksStockStockIdKey)
	return ret, err
}

func (m *defaultStockModel) Update(ctx context.Context, newData *Stock) error {
	data, err := m.FindOne(ctx, newData.Id)
	if err != nil {
		return err
	}

	starbucksStockIdKey := fmt.Sprintf("%s%v", cacheStarbucksStockIdPrefix, data.Id)
	starbucksStockStockIdKey := fmt.Sprintf("%s%v", cacheStarbucksStockStockIdPrefix, data.StockId)
	_, err = m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("update %s set %s where `id` = ?", m.table, stockRowsWithPlaceHolder)
		return conn.ExecCtx(ctx, query, newData.StockId, newData.MaterialId, newData.Amount, newData.Count, newData.InTime, newData.Id)
	}, starbucksStockIdKey, starbucksStockStockIdKey)
	return err
}

func (m *defaultStockModel) formatPrimary(primary any) string {
	return fmt.Sprintf("%s%v", cacheStarbucksStockIdPrefix, primary)
}

func (m *defaultStockModel) queryPrimary(ctx context.Context, conn sqlx.SqlConn, v, primary any) error {
	query := fmt.Sprintf("select %s from %s where `id` = ? limit 1", stockRows, m.table)
	return conn.QueryRowCtx(ctx, v, query, primary)
}

func (m *defaultStockModel) tableName() string {
	return m.table
}
