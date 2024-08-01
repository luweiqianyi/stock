package logic

import (
	"context"
	"database/sql"
	"fmt"
	"stock/cmd/common"
	"stock/cmd/transaction/model"
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

	_, err = l.svcCtx.TransactionModel.Insert(l.ctx, &model.Transaction{
		StockCode: req.StockCode,
		StockName: sql.NullString{
			String: req.StockName,
			Valid:  true,
		},
		BuyPrice: sql.NullFloat64{
			Float64: req.BuyPrice,
			Valid:   true,
		},
		SellPrice: sql.NullFloat64{
			Float64: req.SellPrice,
			Valid:   true,
		},
		Number: sql.NullFloat64{
			Float64: req.Number,
			Valid:   true,
		},
		BuyDate: sql.NullTime{
			Time:  timeBuyDate,
			Valid: true,
		},
		SellDate: sql.NullTime{
			Time:  timeSellDate,
			Valid: true,
		},
	})

	if err != nil {
		return &types.AddOneTransactionRecordResp{
			CommonResp: types.CommonResp{
				Result:  common.FAILED,
				Message: fmt.Sprintf("%v", err),
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
