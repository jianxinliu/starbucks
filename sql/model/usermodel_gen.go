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
	userFieldNames          = builder.RawFieldNames(&User{})
	userRows                = strings.Join(userFieldNames, ",")
	userRowsExpectAutoSet   = strings.Join(stringx.Remove(userFieldNames, "`id`", "`create_at`", "`create_time`", "`created_at`", "`update_at`", "`update_time`", "`updated_at`"), ",")
	userRowsWithPlaceHolder = strings.Join(stringx.Remove(userFieldNames, "`id`", "`create_at`", "`create_time`", "`created_at`", "`update_at`", "`update_time`", "`updated_at`"), "=?,") + "=?"

	cacheStarbucksUserIdPrefix     = "cache:starbucks:user:id:"
	cacheStarbucksUserUserIdPrefix = "cache:starbucks:user:userId:"
)

type (
	userModel interface {
		Insert(ctx context.Context, data *User) (sql.Result, error)
		FindOne(ctx context.Context, id int64) (*User, error)
		FindOneByUserId(ctx context.Context, userId string) (*User, error)
		Update(ctx context.Context, data *User) error
		Delete(ctx context.Context, id int64) error
	}

	defaultUserModel struct {
		sqlc.CachedConn
		table string
	}

	User struct {
		Id           int64        `db:"id"`
		UserId       string       `db:"user_id"`
		UserName     string       `db:"user_name"`
		Password     string       `db:"password"`
		UserType     string       `db:"user_type"`      // 用户类型 normal, vip1, vip2……
		VipStartTime sql.NullTime `db:"vip_start_time"` // 会员起始时间
		VipEndTime   sql.NullTime `db:"vip_end_time"`   // 会员到期时间
	}
)

func newUserModel(conn sqlx.SqlConn, c cache.CacheConf, opts ...cache.Option) *defaultUserModel {
	return &defaultUserModel{
		CachedConn: sqlc.NewConn(conn, c, opts...),
		table:      "`user`",
	}
}

func (m *defaultUserModel) Delete(ctx context.Context, id int64) error {
	data, err := m.FindOne(ctx, id)
	if err != nil {
		return err
	}

	starbucksUserIdKey := fmt.Sprintf("%s%v", cacheStarbucksUserIdPrefix, id)
	starbucksUserUserIdKey := fmt.Sprintf("%s%v", cacheStarbucksUserUserIdPrefix, data.UserId)
	_, err = m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("delete from %s where `id` = ?", m.table)
		return conn.ExecCtx(ctx, query, id)
	}, starbucksUserIdKey, starbucksUserUserIdKey)
	return err
}

func (m *defaultUserModel) FindOne(ctx context.Context, id int64) (*User, error) {
	starbucksUserIdKey := fmt.Sprintf("%s%v", cacheStarbucksUserIdPrefix, id)
	var resp User
	err := m.QueryRowCtx(ctx, &resp, starbucksUserIdKey, func(ctx context.Context, conn sqlx.SqlConn, v any) error {
		query := fmt.Sprintf("select %s from %s where `id` = ? limit 1", userRows, m.table)
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

func (m *defaultUserModel) FindOneByUserId(ctx context.Context, userId string) (*User, error) {
	starbucksUserUserIdKey := fmt.Sprintf("%s%v", cacheStarbucksUserUserIdPrefix, userId)
	var resp User
	err := m.QueryRowIndexCtx(ctx, &resp, starbucksUserUserIdKey, m.formatPrimary, func(ctx context.Context, conn sqlx.SqlConn, v any) (i any, e error) {
		query := fmt.Sprintf("select %s from %s where `user_id` = ? limit 1", userRows, m.table)
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

func (m *defaultUserModel) Insert(ctx context.Context, data *User) (sql.Result, error) {
	starbucksUserIdKey := fmt.Sprintf("%s%v", cacheStarbucksUserIdPrefix, data.Id)
	starbucksUserUserIdKey := fmt.Sprintf("%s%v", cacheStarbucksUserUserIdPrefix, data.UserId)
	ret, err := m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("insert into %s (%s) values (?, ?, ?, ?, ?, ?)", m.table, userRowsExpectAutoSet)
		return conn.ExecCtx(ctx, query, data.UserId, data.UserName, data.Password, data.UserType, data.VipStartTime, data.VipEndTime)
	}, starbucksUserIdKey, starbucksUserUserIdKey)
	return ret, err
}

func (m *defaultUserModel) Update(ctx context.Context, newData *User) error {
	data, err := m.FindOne(ctx, newData.Id)
	if err != nil {
		return err
	}

	starbucksUserIdKey := fmt.Sprintf("%s%v", cacheStarbucksUserIdPrefix, data.Id)
	starbucksUserUserIdKey := fmt.Sprintf("%s%v", cacheStarbucksUserUserIdPrefix, data.UserId)
	_, err = m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("update %s set %s where `id` = ?", m.table, userRowsWithPlaceHolder)
		return conn.ExecCtx(ctx, query, newData.UserId, newData.UserName, newData.Password, newData.UserType, newData.VipStartTime, newData.VipEndTime, newData.Id)
	}, starbucksUserIdKey, starbucksUserUserIdKey)
	return err
}

func (m *defaultUserModel) formatPrimary(primary any) string {
	return fmt.Sprintf("%s%v", cacheStarbucksUserIdPrefix, primary)
}

func (m *defaultUserModel) queryPrimary(ctx context.Context, conn sqlx.SqlConn, v, primary any) error {
	query := fmt.Sprintf("select %s from %s where `id` = ? limit 1", userRows, m.table)
	return conn.QueryRowCtx(ctx, v, query, primary)
}

func (m *defaultUserModel) tableName() string {
	return m.table
}
