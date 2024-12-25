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
	paymentLogsFieldNames          = builder.RawFieldNames(&PaymentLogs{})
	paymentLogsRows                = strings.Join(paymentLogsFieldNames, ",")
	paymentLogsRowsExpectAutoSet   = strings.Join(stringx.Remove(paymentLogsFieldNames, "`id`", "`create_at`", "`create_time`", "`created_at`", "`update_at`", "`update_time`", "`updated_at`"), ",")
	paymentLogsRowsWithPlaceHolder = strings.Join(stringx.Remove(paymentLogsFieldNames, "`id`", "`create_at`", "`create_time`", "`created_at`", "`update_at`", "`update_time`", "`updated_at`"), "=?,") + "=?"
)

type (
	paymentLogsModel interface {
		Insert(ctx context.Context, data *PaymentLogs) (sql.Result, error)
		FindOne(ctx context.Context, id uint64) (*PaymentLogs, error)
		Update(ctx context.Context, data *PaymentLogs) error
		Delete(ctx context.Context, id uint64) error
	}

	defaultPaymentLogsModel struct {
		conn  sqlx.SqlConn
		table string
	}

	PaymentLogs struct {
		Id           uint64         `db:"id"`            // è‡ªå¢žID
		PaymentNo    string         `db:"payment_no"`    // æ”¯ä»˜å•å·
		Type         int64          `db:"type"`          // ç±»åž‹ 1:æ”¯ä»˜ 2:é€€æ¬¾
		Channel      int64          `db:"channel"`       // æ”¯ä»˜æ¸ é“
		RequestData  sql.NullString `db:"request_data"`  // è¯·æ±‚æ•°æ®
		ResponseData sql.NullString `db:"response_data"` // å“åº”æ•°æ®
		CreatedAt    time.Time      `db:"created_at"`    // åˆ›å»ºæ—¶é—´
	}
)

func newPaymentLogsModel(conn sqlx.SqlConn) *defaultPaymentLogsModel {
	return &defaultPaymentLogsModel{
		conn:  conn,
		table: "`payment_logs`",
	}
}

func (m *defaultPaymentLogsModel) Delete(ctx context.Context, id uint64) error {
	query := fmt.Sprintf("delete from %s where `id` = ?", m.table)
	_, err := m.conn.ExecCtx(ctx, query, id)
	return err
}

func (m *defaultPaymentLogsModel) FindOne(ctx context.Context, id uint64) (*PaymentLogs, error) {
	query := fmt.Sprintf("select %s from %s where `id` = ? limit 1", paymentLogsRows, m.table)
	var resp PaymentLogs
	err := m.conn.QueryRowCtx(ctx, &resp, query, id)
	switch err {
	case nil:
		return &resp, nil
	case sqlx.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}

func (m *defaultPaymentLogsModel) Insert(ctx context.Context, data *PaymentLogs) (sql.Result, error) {
	query := fmt.Sprintf("insert into %s (%s) values (?, ?, ?, ?, ?)", m.table, paymentLogsRowsExpectAutoSet)
	ret, err := m.conn.ExecCtx(ctx, query, data.PaymentNo, data.Type, data.Channel, data.RequestData, data.ResponseData)
	return ret, err
}

func (m *defaultPaymentLogsModel) Update(ctx context.Context, data *PaymentLogs) error {
	query := fmt.Sprintf("update %s set %s where `id` = ?", m.table, paymentLogsRowsWithPlaceHolder)
	_, err := m.conn.ExecCtx(ctx, query, data.PaymentNo, data.Type, data.Channel, data.RequestData, data.ResponseData, data.Id)
	return err
}

func (m *defaultPaymentLogsModel) tableName() string {
	return m.table
}