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
	orderPaymentsFieldNames          = builder.RawFieldNames(&OrderPayments{})
	orderPaymentsRows                = strings.Join(orderPaymentsFieldNames, ",")
	orderPaymentsRowsExpectAutoSet   = strings.Join(stringx.Remove(orderPaymentsFieldNames, "`id`", "`create_at`", "`create_time`", "`created_at`", "`update_at`", "`update_time`", "`updated_at`"), ",")
	orderPaymentsRowsWithPlaceHolder = strings.Join(stringx.Remove(orderPaymentsFieldNames, "`id`", "`create_at`", "`create_time`", "`created_at`", "`update_at`", "`update_time`", "`updated_at`"), "=?,") + "=?"
)

type (
	orderPaymentsModel interface {
		Insert(ctx context.Context, data *OrderPayments) (sql.Result, error)
		FindOne(ctx context.Context, id uint64) (*OrderPayments, error)
		FindOneByPaymentNo(ctx context.Context, paymentNo string) (*OrderPayments, error)
		Update(ctx context.Context, data *OrderPayments) error
		Delete(ctx context.Context, id uint64) error
	}

	defaultOrderPaymentsModel struct {
		conn  sqlx.SqlConn
		table string
	}

	OrderPayments struct {
		Id            uint64       `db:"id"`             // æ”¯ä»˜ID
		OrderId       uint64       `db:"order_id"`       // è®¢å•ID
		PaymentNo     string       `db:"payment_no"`     // æ”¯ä»˜æµæ°´å·
		PaymentMethod int64        `db:"payment_method"` // æ”¯ä»˜æ–¹å¼ 1:å¾®ä¿¡ 2:æ”¯ä»˜å® 3:ä½™é¢
		Amount        float64      `db:"amount"`         // æ”¯ä»˜é‡‘é¢
		Status        int64        `db:"status"`         // æ”¯ä»˜çŠ¶æ€ 0:æœªæ”¯ä»˜ 1:å·²æ”¯ä»˜ 2:å·²é€€æ¬¾
		PayTime       sql.NullTime `db:"pay_time"`       // æ”¯ä»˜æ—¶é—´
		CreatedAt     time.Time    `db:"created_at"`     // åˆ›å»ºæ—¶é—´
		UpdatedAt     time.Time    `db:"updated_at"`     // æ›´æ–°æ—¶é—´
	}
)

func newOrderPaymentsModel(conn sqlx.SqlConn) *defaultOrderPaymentsModel {
	return &defaultOrderPaymentsModel{
		conn:  conn,
		table: "`order_payments`",
	}
}

func (m *defaultOrderPaymentsModel) Delete(ctx context.Context, id uint64) error {
	query := fmt.Sprintf("delete from %s where `id` = ?", m.table)
	_, err := m.conn.ExecCtx(ctx, query, id)
	return err
}

func (m *defaultOrderPaymentsModel) FindOne(ctx context.Context, id uint64) (*OrderPayments, error) {
	query := fmt.Sprintf("select %s from %s where `id` = ? limit 1", orderPaymentsRows, m.table)
	var resp OrderPayments
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

func (m *defaultOrderPaymentsModel) FindOneByPaymentNo(ctx context.Context, paymentNo string) (*OrderPayments, error) {
	var resp OrderPayments
	query := fmt.Sprintf("select %s from %s where `payment_no` = ? limit 1", orderPaymentsRows, m.table)
	err := m.conn.QueryRowCtx(ctx, &resp, query, paymentNo)
	switch err {
	case nil:
		return &resp, nil
	case sqlx.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}

func (m *defaultOrderPaymentsModel) Insert(ctx context.Context, data *OrderPayments) (sql.Result, error) {
	query := fmt.Sprintf("insert into %s (%s) values (?, ?, ?, ?, ?, ?)", m.table, orderPaymentsRowsExpectAutoSet)
	ret, err := m.conn.ExecCtx(ctx, query, data.OrderId, data.PaymentNo, data.PaymentMethod, data.Amount, data.Status, data.PayTime)
	return ret, err
}

func (m *defaultOrderPaymentsModel) Update(ctx context.Context, newData *OrderPayments) error {
	query := fmt.Sprintf("update %s set %s where `id` = ?", m.table, orderPaymentsRowsWithPlaceHolder)
	_, err := m.conn.ExecCtx(ctx, query, newData.OrderId, newData.PaymentNo, newData.PaymentMethod, newData.Amount, newData.Status, newData.PayTime, newData.Id)
	return err
}

func (m *defaultOrderPaymentsModel) tableName() string {
	return m.table
}