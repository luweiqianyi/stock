package logic

import (
	"context"
	"fmt"
	"stock/cmd/common"
	"time"

	"stock/cmd/transaction/api/internal/svc"
	"stock/cmd/transaction/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

const (
	dateformat = "2006-01-02 15:04:05"
)

type AddOneTransactionRecordLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewAddOneTransactionRecordLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddOneTransactionRecordLogic {
	return &AddOneTransactionRecordLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *AddOneTransactionRecordLogic) AddOneTransactionRecord(req *types.AddOneTransactionRecordReq) (resp *types.AddOneTransactionRecordResp, err error) {
	timeBuyDate, err := time.Parse(dateformat, req.BuyDate)
	if err != nil {
		return &types.AddOneTransactionRecordResp{
			CommonResp: types.CommonResp{
				Result:  common.FAILED,
				Message: fmt.Sprintf("%v", err),
			},
		}, err
	}

	timeSellDate, err := time.Parse(dateformat, req.SellDate)
	if err != nil {
		return &types.AddOneTransactionRecordResp{
			CommonResp: types.CommonResp{
				Result:  common.FAILED,
				Message: fmt.Sprintf("%v", err),
			},
		}, err
	}

	tx := l.svcCtx.DB.Exec("insert into transaction(`stock_code`,`stock_name`,`market`,`buy_price`,`sell_price`,`number`,`buy_date`,`sell_date`)values(?,?,?,?,?,?,?,?)",
		req.StockCode,
		req.StockName,
		req.Market,
		req.BuyPrice,
		req.SellPrice,
		req.Number,
		timeBuyDate,
		timeSellDate,
	)

	if tx.Error != nil {
		return &types.AddOneTransactionRecordResp{
			CommonResp: types.CommonResp{
				Result:  common.FAILED,
				Message: fmt.Sprintf("%v", tx.Error),
			},
		}, err
	}

	return &types.AddOneTransactionRecordResp{
		CommonResp: types.CommonResp{
			Result:  common.SUCCESS,
			Message: "success",
		},
	}, err
}
