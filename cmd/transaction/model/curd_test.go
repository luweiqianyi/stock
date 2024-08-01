package model

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"testing"
	"time"
)

const (
	url = "root:123456@tcp(127.0.0.1:3306)/stock?charset=utf8mb4&parseTime=true&loc=Asia%2FShanghai"
)

// 插入一条记录
func TestTransactionModelInsert(t *testing.T) {
	ctx := context.Background()
	tModel := NewTransactionModel(sqlx.NewMysql(url))
	result, err := tModel.Insert(ctx, &Transaction{
		StockCode: "600250",
		StockName: sql.NullString{
			String: "南京商旅",
			Valid:  true,
		},
		BuyPrice: sql.NullFloat64{
			Float64: 8.43,
			Valid:   true,
		},
		SellPrice: sql.NullFloat64{
			Float64: 9.0,
			Valid:   true,
		},
		Number: sql.NullFloat64{
			Float64: 14000,
			Valid:   true,
		},
		BuyDate: sql.NullTime{
			Time:  time.Now(),
			Valid: true,
		},
		SellDate: sql.NullTime{
			Time:  time.Now(),
			Valid: true,
		},
	})
	if err != nil {
		fmt.Printf("error: %v\n", err)
		return
	}

	fmt.Printf("result: %#v\n", result)
}

// 查找已存在的记录
func TestTransactionModelQuery(t *testing.T) {
	ctx := context.Background()
	tModel := NewTransactionModel(sqlx.NewMysql(url))

	result, err := tModel.FindOne(ctx, 1)
	if err != nil {
		fmt.Printf("error: %v\n", err)
		return
	}

	fmt.Printf("result: %#v\n", result)
}

// 查找不存在的记录
func TestTransactionModelQueryNotExist(t *testing.T) {
	ctx := context.Background()
	tModel := NewTransactionModel(sqlx.NewMysql(url))

	result, err := tModel.FindOne(ctx, 2)
	if err != nil {
		fmt.Printf("error: %v\n", err)
		return
	}

	fmt.Printf("result: %#v\n", result)
}

// 删除已存在的记录
func TestTransactionModelDelete(t *testing.T) {
	ctx := context.Background()
	tModel := NewTransactionModel(sqlx.NewMysql(url))

	err := tModel.Delete(ctx, 1)
	fmt.Printf("error: %v\n", err)
}

// 删除不存在的记录
func TestTransactionModelDeleteNotExist(t *testing.T) {
	ctx := context.Background()
	tModel := NewTransactionModel(sqlx.NewMysql(url))

	err := tModel.Delete(ctx, 2)
	fmt.Printf("error: %v\n", err)
}
