package logic

import (
	"context"

	"stock/cmd/transaction/api/internal/svc"
	"stock/cmd/transaction/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
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
	// todo: add your logic here and delete this line

	return
}
