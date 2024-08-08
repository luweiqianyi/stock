package logic

import (
	"context"
	"fmt"
	"stock/cmd/common"
	"stock/pkg/fee"

	"stock/cmd/transaction/api/internal/svc"
	"stock/cmd/transaction/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type CalculateTransactionProfitLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCalculateTransactionProfitLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CalculateTransactionProfitLogic {
	return &CalculateTransactionProfitLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CalculateTransactionProfitLogic) CalculateTransactionProfit(req *types.CalculateTransactionProfitReq) (resp *types.CalculateTransactionProfitResp, err error) {
	tr := fee.Transaction{
		Entry: fee.StockEntry{
			Code:   req.StockCode,
			Name:   req.StockName,
			Market: fee.Market(req.Market),
		},
		BuyPrice:  req.BuyPrice,
		SellPrice: req.SellPrice,
		Number:    req.Number,
	}

	transactionResult := types.TransactionRecordResult{
		TransactionDetailResp: types.TransactionDetailResp{
			StockCode: req.StockCode,
			StockName: req.StockName,
			Market:    req.Market,
			BuyPrice:  req.BuyPrice,
			Number:    req.Number,
			SellPrice: req.SellPrice,
			BuyDate:   req.BuyDate,
			SellDate:  req.SellDate,
		},
		TransactionResultResp: types.TransactionResultResp{
			BuyCost:     tr.BuyFee(),
			SellCost:    tr.SellFee(),
			TotalCost:   tr.TotalFee(),
			Rate:        tr.Ratio(),
			GainLoss:    tr.ProfitAndLoss(),
			FinalProfit: tr.FinalProfit(),
		},
	}

	tx := l.svcCtx.DB.Exec("insert into transaction_result(`stock_code`,`stock_name`,`market`,`buy_price`,`number`,`sell_price`,`buy_date`,`sell_date`,`buy_cost`,`sell_cost`,`total_cost`,`rate`,`gain_loss`,`final_profit`)values(?,?,?,?,?,?,?,?,?,?,?,?,?,?)",
		req.StockCode,
		req.StockName,
		req.Market,
		req.BuyPrice,
		req.Number,
		req.SellPrice,
		req.BuyDate,
		req.SellDate,
		tr.BuyFee(),
		tr.SellFee(),
		tr.TotalFee(),
		tr.Ratio(),
		tr.ProfitAndLoss(),
		tr.FinalProfit(),
	)
	if tx.Error != nil {
		// 重复计算的话也返回计算结果，但是如果插入数据库失败，告知接口请求方失败原因
		resp = &types.CalculateTransactionProfitResp{
			CommonResp: types.CommonResp{
				Result:  common.FAILED,
				Message: fmt.Sprintf("error: %v", tx.Error),
			},
			TransactionRecordResult: transactionResult,
		}
		return
	}

	resp = &types.CalculateTransactionProfitResp{
		CommonResp: types.CommonResp{
			Result: common.SUCCESS,
		},
		TransactionRecordResult: transactionResult,
	}

	return
}
