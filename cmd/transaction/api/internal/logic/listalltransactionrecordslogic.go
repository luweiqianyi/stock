package logic

import (
	"context"
	"fmt"
	"stock/cmd/common"

	"stock/cmd/transaction/api/internal/svc"
	"stock/cmd/transaction/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type ListAllTransactionRecordsLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewListAllTransactionRecordsLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ListAllTransactionRecordsLogic {
	return &ListAllTransactionRecordsLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ListAllTransactionRecordsLogic) ListAllTransactionRecords(req *types.ListAllTransactionRecordsReq) (resp *types.ListAllTransactionRecordsResp, err error) {
	sql := fmt.Sprintf("select * from transaction_result")
	var results []types.TransactionRecordResult
	tx := l.svcCtx.DB.Raw(sql).Scan(&results)
	if tx.Error != nil {
		resp = &types.ListAllTransactionRecordsResp{
			CommonResp: types.CommonResp{
				Result:  common.FAILED,
				Message: fmt.Sprintf("error: %v", tx.Error),
			},
			TransactionResults: []types.TransactionRecordResult{},
		}
		return
	}

	resp = &types.ListAllTransactionRecordsResp{
		CommonResp: types.CommonResp{
			Result: common.SUCCESS,
		},
		TransactionResults: results,
	}

	return
}
