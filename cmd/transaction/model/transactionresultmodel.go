package model

import "github.com/zeromicro/go-zero/core/stores/sqlx"

var _ TransactionResultModel = (*customTransactionResultModel)(nil)

type (
	// TransactionResultModel is an interface to be customized, add more methods here,
	// and implement the added methods in customTransactionResultModel.
	TransactionResultModel interface {
		transactionResultModel
		withSession(session sqlx.Session) TransactionResultModel
	}

	customTransactionResultModel struct {
		*defaultTransactionResultModel
	}
)

// NewTransactionResultModel returns a model for the database table.
func NewTransactionResultModel(conn sqlx.SqlConn) TransactionResultModel {
	return &customTransactionResultModel{
		defaultTransactionResultModel: newTransactionResultModel(conn),
	}
}

func (m *customTransactionResultModel) withSession(session sqlx.Session) TransactionResultModel {
	return NewTransactionResultModel(sqlx.NewSqlConnFromSession(session))
}
