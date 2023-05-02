package model

import (
	"context"
	"fmt"
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlc"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"starbucks/starbucks/utils"
)

var _ OrderModel = (*customOrderModel)(nil)

type (
	// OrderModel is an interface to be customized, add more methods here,
	// and implement the added methods in customOrderModel.
	OrderModel interface {
		orderModel
		FindOneBy(ctx context.Context, filedName, fieldValue string) (*Order, error)
		NewOrderId(ctx context.Context) string
		NewTrxNo(ctx context.Context) string
	}

	customOrderModel struct {
		*defaultOrderModel
	}
)

// NewOrderModel returns a model for the database table.
func NewOrderModel(conn sqlx.SqlConn, c cache.CacheConf, opts ...cache.Option) OrderModel {
	return &customOrderModel{
		defaultOrderModel: newOrderModel(conn, c, opts...),
	}
}

func (m *defaultOrderModel) FindOneBy(ctx context.Context, fieldName, fieldValue string) (*Order, error) {
	starbucksOrderOrderIdKey := fmt.Sprintf("%s%v", cacheStarbucksOrderOrderIdPrefix, fieldName)
	var resp Order
	err := m.QueryRowIndexCtx(ctx, &resp, starbucksOrderOrderIdKey, m.formatPrimary, func(ctx context.Context, conn sqlx.SqlConn, v any) (i any, e error) {
		query := fmt.Sprintf("select %s from %s where `?` = ? limit 1", orderRows, m.table)
		if err := conn.QueryRowCtx(ctx, &resp, query, fieldName, fieldValue); err != nil {
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

func (m customOrderModel) NewOrderId(ctx context.Context) string {
	id := utils.NewOrderId()
	for i := 0; i < 10; i++ {
		_, err := m.FindOneByOrderId(ctx, id)
		if err != nil && err == sqlx.ErrNotFound {
			break
		}
		id = utils.NewOrderId()
	}
	return id
}

func (m customOrderModel) NewTrxNo(ctx context.Context) string {
	id := utils.NewTrxNo()
	for i := 0; i < 10; i++ {
		_, err := m.FindOneBy(ctx, "trx_no", id)
		if err != nil && err == sqlx.ErrNotFound {
			break
		}
		id = utils.NewTrxNo()
	}
	return id
}
