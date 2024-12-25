// Code generated by goctl. DO NOT EDIT.
// versions:
//  goctl version: 1.7.3

package model

import (
	"context"
	"database/sql"
	"fmt"
	"strings"
	"time"

	"github.com/zeromicro/go-zero/core/stores/builder"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"github.com/zeromicro/go-zero/core/stringx"
)

var (
	cartStatisticsFieldNames          = builder.RawFieldNames(&CartStatistics{})
	cartStatisticsRows                = strings.Join(cartStatisticsFieldNames, ",")
	cartStatisticsRowsExpectAutoSet   = strings.Join(stringx.Remove(cartStatisticsFieldNames, "`create_at`", "`create_time`", "`created_at`", "`update_at`", "`update_time`", "`updated_at`"), ",")
	cartStatisticsRowsWithPlaceHolder = strings.Join(stringx.Remove(cartStatisticsFieldNames, "`user_id`", "`create_at`", "`create_time`", "`created_at`", "`update_at`", "`update_time`", "`updated_at`"), "=?,") + "=?"
)

type (
	cartStatisticsModel interface {
		Insert(ctx context.Context, data *CartStatistics) (sql.Result, error)
		FindOne(ctx context.Context, userId uint64) (*CartStatistics, error)
		Update(ctx context.Context, data *CartStatistics) error
		Delete(ctx context.Context, userId uint64) error
	}

	defaultCartStatisticsModel struct {
		conn  sqlx.SqlConn
		table string
	}

	CartStatistics struct {
		UserId           uint64    `db:"user_id"`           // ç”¨æˆ·ID
		TotalQuantity    int64     `db:"total_quantity"`    // å•†å“æ€»æ•°é‡
		SelectedQuantity int64     `db:"selected_quantity"` // å·²é€‰å•†å“æ•°é‡
		TotalAmount      float64   `db:"total_amount"`      // å•†å“æ€»é‡‘é¢
		SelectedAmount   float64   `db:"selected_amount"`   // å·²é€‰å•†å“é‡‘é¢
		UpdatedAt        time.Time `db:"updated_at"`        // æ›´æ–°æ—¶é—´
	}
)

func newCartStatisticsModel(conn sqlx.SqlConn) *defaultCartStatisticsModel {
	return &defaultCartStatisticsModel{
		conn:  conn,
		table: "`cart_statistics`",
	}
}

func (m *defaultCartStatisticsModel) Delete(ctx context.Context, userId uint64) error {
	query := fmt.Sprintf("delete from %s where `user_id` = ?", m.table)
	_, err := m.conn.ExecCtx(ctx, query, userId)
	return err
}

func (m *defaultCartStatisticsModel) FindOne(ctx context.Context, userId uint64) (*CartStatistics, error) {
	query := fmt.Sprintf("select %s from %s where `user_id` = ? limit 1", cartStatisticsRows, m.table)
	var resp CartStatistics
	err := m.conn.QueryRowCtx(ctx, &resp, query, userId)
	switch err {
	case nil:
		return &resp, nil
	case sqlx.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}

func (m *defaultCartStatisticsModel) Insert(ctx context.Context, data *CartStatistics) (sql.Result, error) {
	query := fmt.Sprintf("insert into %s (%s) values (?, ?, ?, ?, ?)", m.table, cartStatisticsRowsExpectAutoSet)
	ret, err := m.conn.ExecCtx(ctx, query, data.UserId, data.TotalQuantity, data.SelectedQuantity, data.TotalAmount, data.SelectedAmount)
	return ret, err
}

func (m *defaultCartStatisticsModel) Update(ctx context.Context, data *CartStatistics) error {
	query := fmt.Sprintf("update %s set %s where `user_id` = ?", m.table, cartStatisticsRowsWithPlaceHolder)
	_, err := m.conn.ExecCtx(ctx, query, data.TotalQuantity, data.SelectedQuantity, data.TotalAmount, data.SelectedAmount, data.UserId)
	return err
}

func (m *defaultCartStatisticsModel) tableName() string {
	return m.table
}