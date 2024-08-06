package logic

import (
	"context"
	"fmt"
	"stock/cmd/common"

	"stock/cmd/market/api/internal/svc"
	"stock/cmd/market/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type ListMarketTypeLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewListMarketTypeLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ListMarketTypeLogic {
	return &ListMarketTypeLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ListMarketTypeLogic) ListMarketType(req *types.ListMarketTypeReq) (resp *types.ListMarketTypeResp, err error) {
	rows, err := l.svcCtx.DB.Raw("select market_type from market").Rows()
	if err != nil {
		resp = &types.ListMarketTypeResp{
			CommonResp: types.CommonResp{
				Result:  common.FAILED,
				Message: fmt.Sprintf("error: %v", err),
			},
		}
		return
	}

	var marketType string
	var markets []string
	for rows.Next() {
		if err := rows.Scan(&marketType); err != nil {
			continue
		}
		markets = append(markets, marketType)
	}

	resp = &types.ListMarketTypeResp{
		CommonResp: types.CommonResp{
			Result: common.SUCCESS,
		},
		MarketType: markets,
	}
	return
}
