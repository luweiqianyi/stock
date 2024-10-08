// Code generated by goctl. DO NOT EDIT.

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
	transactionFieldNames          = builder.RawFieldNames(&Transaction{})
	transactionRows                = strings.Join(transactionFieldNames, ",")
	transactionRowsExpectAutoSet   = strings.Join(stringx.Remove(transactionFieldNames, "`id`", "`create_at`", "`create_time`", "`created_at`", "`update_at`", "`update_time`", "`updated_at`"), ",")
	transactionRowsWithPlaceHolder = strings.Join(stringx.Remove(transactionFieldNames, "`id`", "`create_at`", "`create_time`", "`created_at`", "`update_at`", "`update_time`", "`updated_at`"), "=?,") + "=?"
)

type (
	transactionModel interface {
		Insert(ctx context.Context, data *Transaction) (sql.Result, error)
		FindOne(ctx context.Context, id int64) (*Transaction, error)
		Update(ctx context.Context, data *Transaction) error
		Delete(ctx context.Context, id int64) error
	}

	defaultTransactionModel struct {
		conn  sqlx.SqlConn
		table string
	}

	Transaction struct {
		Id        int64     `db:"id"`
		StockCode string    `db:"stock_code"` // stock_code
		StockName string    `db:"stock_name"` // stock_name
		BuyPrice  float64   `db:"buy_price"`  // buy_price
		SellPrice float64   `db:"sell_price"` // sell_price
		Number    float64   `db:"number"`     // number
		BuyDate   time.Time `db:"buy_date"`   // buy_date
		SellDate  time.Time `db:"sell_date"`  // sell_date
	}
)

func newTransactionModel(conn sqlx.SqlConn) *defaultTransactionModel {
	return &defaultTransactionModel{
		conn:  conn,
		table: "`transaction`",
	}
}

func (m *defaultTransactionModel) Delete(ctx context.Context, id int64) error {
	query := fmt.Sprintf("delete from %s where `id` = ?", m.table)
	_, err := m.conn.ExecCtx(ctx, query, id)
	return err
}

func (m *defaultTransactionModel) FindOne(ctx context.Context, id int64) (*Transaction, error) {
	query := fmt.Sprintf("select %s from %s where `id` = ? limit 1", transactionRows, m.table)
	var resp Transaction
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

func (m *defaultTransactionModel) Insert(ctx context.Context, data *Transaction) (sql.Result, error) {
	query := fmt.Sprintf("insert into %s (%s) values (?, ?, ?, ?, ?, ?, ?)", m.table, transactionRowsExpectAutoSet)
	ret, err := m.conn.ExecCtx(ctx, query, data.StockCode, data.StockName, data.BuyPrice, data.SellPrice, data.Number, data.BuyDate, data.SellDate)
	return ret, err
}

func (m *defaultTransactionModel) Update(ctx context.Context, data *Transaction) error {
	query := fmt.Sprintf("update %s set %s where `id` = ?", m.table, transactionRowsWithPlaceHolder)
	_, err := m.conn.ExecCtx(ctx, query, data.StockCode, data.StockName, data.BuyPrice, data.SellPrice, data.Number, data.BuyDate, data.SellDate, data.Id)
	return err
}

func (m *defaultTransactionModel) tableName() string {
	return m.table
}
