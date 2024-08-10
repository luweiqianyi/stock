package logic

import (
	"context"
	"math"
	"stock/cmd/common"
	"stock/pkg/fee"

	"stock/cmd/transaction/api/internal/svc"
	"stock/cmd/transaction/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type CalExpectedReturnsLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCalExpectedReturnsLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CalExpectedReturnsLogic {
	return &CalExpectedReturnsLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CalExpectedReturnsLogic) CalExpectedReturns(req *types.CalExpectedReturnsReq) (resp *types.CalExpectedReturnsResp, err error) {
	number := float64(int(math.Floor(req.Balance/req.BuyPrice)) / 100 * 100)
	tr := fee.Transaction{
		Entry: fee.StockEntry{
			Market: fee.Market(req.Market),
		},
		BuyPrice:  req.BuyPrice,
		SellPrice: req.SellPrice,
		Number:    number,
	}

	resp = &types.CalExpectedReturnsResp{
		CommonResp: types.CommonResp{
			Result: common.SUCCESS,
		},
		Data: types.CalExpectedReturnsData{
			Market:           req.Market,
			BuyPrice:         req.BuyPrice,
			SellPrice:        req.SellPrice,
			Number:           number,
			InvestedCaptical: tr.Cost(),
			Profit:           tr.FinalProfit(),
		},
	}
	return
}
