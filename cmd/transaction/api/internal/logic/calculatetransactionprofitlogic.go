package logic

import (
	"context"
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
			Market: fee.Market(req.MarketType),
		},
		BuyPrice:  req.BuyPrice,
		SellPrice: req.SellPrice,
		Number:    req.Number,
	}

	transactionResult := types.TransactionRecordResult{
		TransactionDetailResp: types.TransactionDetailResp{
			StockCode:  req.StockCode,
			StockName:  req.StockName,
			MarketType: req.MarketType,
			BuyPrice:   req.BuyPrice,
			Number:     req.Number,
			SellPrice:  req.SellPrice,
			BuyDate:    req.BuyDate,
			SellDate:   req.SellDate,
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
	resp = &types.CalculateTransactionProfitResp{
		CommonResp: types.CommonResp{
			Result: common.SUCCESS,
		},
		TransactionRecordResult: transactionResult,
	}

	return
}
