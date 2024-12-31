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
	walletTransactionsFieldNames          = builder.RawFieldNames(&WalletTransactions{})
	walletTransactionsRows                = strings.Join(walletTransactionsFieldNames, ",")
	walletTransactionsRowsExpectAutoSet   = strings.Join(stringx.Remove(walletTransactionsFieldNames, "`id`", "`create_at`", "`create_time`", "`created_at`", "`update_at`", "`update_time`", "`updated_at`"), ",")
	walletTransactionsRowsWithPlaceHolder = strings.Join(stringx.Remove(walletTransactionsFieldNames, "`id`", "`create_at`", "`create_time`", "`created_at`", "`update_at`", "`update_time`", "`updated_at`"), "=?,") + "=?"
)

type (
	walletTransactionsModel interface {
		Insert(ctx context.Context, data *WalletTransactions) (sql.Result, error)
		FindOne(ctx context.Context, id uint64) (*WalletTransactions, error)
		FindOneByOrderId(ctx context.Context, orderId string) (*WalletTransactions, error)
		Update(ctx context.Context, data *WalletTransactions) error
		Delete(ctx context.Context, id uint64) error
	}

	defaultWalletTransactionsModel struct {
		conn  sqlx.SqlConn
		table string
	}

	WalletTransactions struct {
		Id        uint64         `db:"id"`         // äº¤æ˜“ID
		UserId    uint64         `db:"user_id"`    // ç”¨æˆ·ID
		OrderId   string         `db:"order_id"`   // è®¢å•å·
		Amount    float64        `db:"amount"`     // äº¤æ˜“é‡‘é¢
		Type      int64          `db:"type"`       // äº¤æ˜“ç±»åž‹ 1:å……å€¼ 2:æçŽ° 3:æ¶ˆè´¹ 4:é€€æ¬¾
		Status    int64          `db:"status"`     // äº¤æ˜“çŠ¶æ€ 0:å¤„ç†ä¸­ 1:æˆåŠŸ 2:å¤±è´¥
		Remark    sql.NullString `db:"remark"`     // äº¤æ˜“å¤‡æ³¨
		CreatedAt time.Time      `db:"created_at"` // åˆ›å»ºæ—¶é—´
		UpdatedAt time.Time      `db:"updated_at"` // æ›´æ–°æ—¶é—´
	}
)

func newWalletTransactionsModel(conn sqlx.SqlConn) *defaultWalletTransactionsModel {
	return &defaultWalletTransactionsModel{
		conn:  conn,
		table: "`wallet_transactions`",
	}
}

func (m *defaultWalletTransactionsModel) Delete(ctx context.Context, id uint64) error {
	query := fmt.Sprintf("delete from %s where `id` = ?", m.table)
	_, err := m.conn.ExecCtx(ctx, query, id)
	return err
}

func (m *defaultWalletTransactionsModel) FindOne(ctx context.Context, id uint64) (*WalletTransactions, error) {
	query := fmt.Sprintf("select %s from %s where `id` = ? limit 1", walletTransactionsRows, m.table)
	var resp WalletTransactions
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

func (m *defaultWalletTransactionsModel) FindOneByOrderId(ctx context.Context, orderId string) (*WalletTransactions, error) {
	var resp WalletTransactions
	query := fmt.Sprintf("select %s from %s where `order_id` = ? limit 1", walletTransactionsRows, m.table)
	err := m.conn.QueryRowCtx(ctx, &resp, query, orderId)
	switch err {
	case nil:
		return &resp, nil
	case sqlx.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}

func (m *defaultWalletTransactionsModel) Insert(ctx context.Context, data *WalletTransactions) (sql.Result, error) {
	query := fmt.Sprintf("insert into %s (%s) values (?, ?, ?, ?, ?, ?)", m.table, walletTransactionsRowsExpectAutoSet)
	ret, err := m.conn.ExecCtx(ctx, query, data.UserId, data.OrderId, data.Amount, data.Type, data.Status, data.Remark)
	return ret, err
}

func (m *defaultWalletTransactionsModel) Update(ctx context.Context, newData *WalletTransactions) error {
	query := fmt.Sprintf("update %s set %s where `id` = ?", m.table, walletTransactionsRowsWithPlaceHolder)
	_, err := m.conn.ExecCtx(ctx, query, newData.UserId, newData.OrderId, newData.Amount, newData.Type, newData.Status, newData.Remark, newData.Id)
	return err
}

func (m *defaultWalletTransactionsModel) tableName() string {
	return m.table
}
