package logic

import (
	"context"

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
	// todo: add your logic here and delete this line

	return
}
