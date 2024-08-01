package svc

import (
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"stock/cmd/transaction/api/internal/config"
	"stock/cmd/transaction/model"
)

type ServiceContext struct {
	Config           config.Config
	TransactionModel model.TransactionModel
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:           c,
		TransactionModel: model.NewTransactionModel(sqlx.NewMysql(c.Mysql.DataSource)),
	}
}
