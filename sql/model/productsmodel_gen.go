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
	productsFieldNames          = builder.RawFieldNames(&Products{})
	productsRows                = strings.Join(productsFieldNames, ",")
	productsRowsExpectAutoSet   = strings.Join(stringx.Remove(productsFieldNames, "`id`", "`create_at`", "`create_time`", "`created_at`", "`update_at`", "`update_time`", "`updated_at`"), ",")
	productsRowsWithPlaceHolder = strings.Join(stringx.Remove(productsFieldNames, "`id`", "`create_at`", "`create_time`", "`created_at`", "`update_at`", "`update_time`", "`updated_at`"), "=?,") + "=?"

	cacheStarbucksProductsIdPrefix        = "cache:starbucks:products:id:"
	cacheStarbucksProductsProductIdPrefix = "cache:starbucks:products:productId:"
)

type (
	productsModel interface {
		Insert(ctx context.Context, data *Products) (sql.Result, error)
		FindOne(ctx context.Context, id int64) (*Products, error)
		FindOneByProductId(ctx context.Context, productId string) (*Products, error)
		Update(ctx context.Context, data *Products) error
		Delete(ctx context.Context, id int64) error
	}

	defaultProductsModel struct {
		sqlc.CachedConn
		table string
	}

	Products struct {
		Id          int64          `db:"id"`
		ProductId   string         `db:"product_id"`
		Name        string         `db:"name"`
		Description sql.NullString `db:"description"`
		Image       sql.NullString `db:"image"`
		GroupId     string         `db:"group_id"` // 产品分组
		Price       sql.NullInt64  `db:"price"`    // 价格，单位：分
		Discount    float64        `db:"discount"` // 折扣
	}
)

func newProductsModel(conn sqlx.SqlConn, c cache.CacheConf, opts ...cache.Option) *defaultProductsModel {
	return &defaultProductsModel{
		CachedConn: sqlc.NewConn(conn, c, opts...),
		table:      "`products`",
	}
}

func (m *defaultProductsModel) Delete(ctx context.Context, id int64) error {
	data, err := m.FindOne(ctx, id)
	if err != nil {
		return err
	}

	starbucksProductsIdKey := fmt.Sprintf("%s%v", cacheStarbucksProductsIdPrefix, id)
	starbucksProductsProductIdKey := fmt.Sprintf("%s%v", cacheStarbucksProductsProductIdPrefix, data.ProductId)
	_, err = m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("delete from %s where `id` = ?", m.table)
		return conn.ExecCtx(ctx, query, id)
	}, starbucksProductsIdKey, starbucksProductsProductIdKey)
	return err
}

func (m *defaultProductsModel) FindOne(ctx context.Context, id int64) (*Products, error) {
	starbucksProductsIdKey := fmt.Sprintf("%s%v", cacheStarbucksProductsIdPrefix, id)
	var resp Products
	err := m.QueryRowCtx(ctx, &resp, starbucksProductsIdKey, func(ctx context.Context, conn sqlx.SqlConn, v any) error {
		query := fmt.Sprintf("select %s from %s where `id` = ? limit 1", productsRows, m.table)
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

func (m *defaultProductsModel) FindOneByProductId(ctx context.Context, productId string) (*Products, error) {
	starbucksProductsProductIdKey := fmt.Sprintf("%s%v", cacheStarbucksProductsProductIdPrefix, productId)
	var resp Products
	err := m.QueryRowIndexCtx(ctx, &resp, starbucksProductsProductIdKey, m.formatPrimary, func(ctx context.Context, conn sqlx.SqlConn, v any) (i any, e error) {
		query := fmt.Sprintf("select %s from %s where `product_id` = ? limit 1", productsRows, m.table)
		if err := conn.QueryRowCtx(ctx, &resp, query, productId); err != nil {
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

func (m *defaultProductsModel) Insert(ctx context.Context, data *Products) (sql.Result, error) {
	starbucksProductsIdKey := fmt.Sprintf("%s%v", cacheStarbucksProductsIdPrefix, data.Id)
	starbucksProductsProductIdKey := fmt.Sprintf("%s%v", cacheStarbucksProductsProductIdPrefix, data.ProductId)
	ret, err := m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("insert into %s (%s) values (?, ?, ?, ?, ?, ?, ?)", m.table, productsRowsExpectAutoSet)
		return conn.ExecCtx(ctx, query, data.ProductId, data.Name, data.Description, data.Image, data.GroupId, data.Price, data.Discount)
	}, starbucksProductsIdKey, starbucksProductsProductIdKey)
	return ret, err
}

func (m *defaultProductsModel) Update(ctx context.Context, newData *Products) error {
	data, err := m.FindOne(ctx, newData.Id)
	if err != nil {
		return err
	}

	starbucksProductsIdKey := fmt.Sprintf("%s%v", cacheStarbucksProductsIdPrefix, data.Id)
	starbucksProductsProductIdKey := fmt.Sprintf("%s%v", cacheStarbucksProductsProductIdPrefix, data.ProductId)
	_, err = m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("update %s set %s where `id` = ?", m.table, productsRowsWithPlaceHolder)
		return conn.ExecCtx(ctx, query, newData.ProductId, newData.Name, newData.Description, newData.Image, newData.GroupId, newData.Price, newData.Discount, newData.Id)
	}, starbucksProductsIdKey, starbucksProductsProductIdKey)
	return err
}

func (m *defaultProductsModel) formatPrimary(primary any) string {
	return fmt.Sprintf("%s%v", cacheStarbucksProductsIdPrefix, primary)
}

func (m *defaultProductsModel) queryPrimary(ctx context.Context, conn sqlx.SqlConn, v, primary any) error {
	query := fmt.Sprintf("select %s from %s where `id` = ? limit 1", productsRows, m.table)
	return conn.QueryRowCtx(ctx, v, query, primary)
}

func (m *defaultProductsModel) tableName() string {
	return m.table
}
