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
	refundOrdersFieldNames          = builder.RawFieldNames(&RefundOrders{})
	refundOrdersRows                = strings.Join(refundOrdersFieldNames, ",")
	refundOrdersRowsExpectAutoSet   = strings.Join(stringx.Remove(refundOrdersFieldNames, "`id`", "`create_at`", "`create_time`", "`created_at`", "`update_at`", "`update_time`", "`updated_at`"), ",")
	refundOrdersRowsWithPlaceHolder = strings.Join(stringx.Remove(refundOrdersFieldNames, "`id`", "`create_at`", "`create_time`", "`created_at`", "`update_at`", "`update_time`", "`updated_at`"), "=?,") + "=?"
)

type (
	refundOrdersModel interface {
		Insert(ctx context.Context, data *RefundOrders) (sql.Result, error)
		FindOne(ctx context.Context, id uint64) (*RefundOrders, error)
		FindOneByRefundNo(ctx context.Context, refundNo string) (*RefundOrders, error)
		Update(ctx context.Context, data *RefundOrders) error
		Delete(ctx context.Context, id uint64) error
	}

	defaultRefundOrdersModel struct {
		conn  sqlx.SqlConn
		table string
	}

	RefundOrders struct {
		Id          uint64         `db:"id"`           // è‡ªå¢žID
		RefundNo    string         `db:"refund_no"`    // é€€æ¬¾å•å·
		PaymentNo   string         `db:"payment_no"`   // æ”¯ä»˜å•å·
		OrderNo     string         `db:"order_no"`     // è®¢å•å·
		UserId      uint64         `db:"user_id"`      // ç”¨æˆ·ID
		Amount      float64        `db:"amount"`       // é€€æ¬¾é‡‘é¢
		Reason      string         `db:"reason"`       // é€€æ¬¾åŽŸå›
		Status      int64          `db:"status"`       // çŠ¶æ€ 1:å¾…å¤„ç† 2:å¤„ç†ä¸­ 3:å·²é€€æ¬¾ 4:é€€æ¬¾å¤±è´¥
		ChannelData sql.NullString `db:"channel_data"` // é€€æ¬¾æ¸ é“æ•°æ®
		NotifyUrl   sql.NullString `db:"notify_url"`   // å›žè°ƒåœ°å€
		RefundTime  sql.NullTime   `db:"refund_time"`  // é€€æ¬¾æ—¶é—´
		CreatedAt   time.Time      `db:"created_at"`   // åˆ›å»ºæ—¶é—´
		UpdatedAt   time.Time      `db:"updated_at"`   // æ›´æ–°æ—¶é—´
	}
)

func newRefundOrdersModel(conn sqlx.SqlConn) *defaultRefundOrdersModel {
	return &defaultRefundOrdersModel{
		conn:  conn,
		table: "`refund_orders`",
	}
}

func (m *defaultRefundOrdersModel) Delete(ctx context.Context, id uint64) error {
	query := fmt.Sprintf("delete from %s where `id` = ?", m.table)
	_, err := m.conn.ExecCtx(ctx, query, id)
	return err
}

func (m *defaultRefundOrdersModel) FindOne(ctx context.Context, id uint64) (*RefundOrders, error) {
	query := fmt.Sprintf("select %s from %s where `id` = ? limit 1", refundOrdersRows, m.table)
	var resp RefundOrders
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

func (m *defaultRefundOrdersModel) FindOneByRefundNo(ctx context.Context, refundNo string) (*RefundOrders, error) {
	var resp RefundOrders
	query := fmt.Sprintf("select %s from %s where `refund_no` = ? limit 1", refundOrdersRows, m.table)
	err := m.conn.QueryRowCtx(ctx, &resp, query, refundNo)
	switch err {
	case nil:
		return &resp, nil
	case sqlx.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}

func (m *defaultRefundOrdersModel) Insert(ctx context.Context, data *RefundOrders) (sql.Result, error) {
	query := fmt.Sprintf("insert into %s (%s) values (?, ?, ?, ?, ?, ?, ?, ?, ?, ?)", m.table, refundOrdersRowsExpectAutoSet)
	ret, err := m.conn.ExecCtx(ctx, query, data.RefundNo, data.PaymentNo, data.OrderNo, data.UserId, data.Amount, data.Reason, data.Status, data.ChannelData, data.NotifyUrl, data.RefundTime)
	return ret, err
}

func (m *defaultRefundOrdersModel) Update(ctx context.Context, newData *RefundOrders) error {
	query := fmt.Sprintf("update %s set %s where `id` = ?", m.table, refundOrdersRowsWithPlaceHolder)
	_, err := m.conn.ExecCtx(ctx, query, newData.RefundNo, newData.PaymentNo, newData.OrderNo, newData.UserId, newData.Amount, newData.Reason, newData.Status, newData.ChannelData, newData.NotifyUrl, newData.RefundTime, newData.Id)
	return err
}

func (m *defaultRefundOrdersModel) tableName() string {
	return m.table
}