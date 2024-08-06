package logic

import (
	"context"
	"fmt"
	"stock/cmd/common"
	"stock/cmd/market/api/internal/svc"
	"stock/cmd/market/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type AddMarketTypeLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewAddMarketTypeLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddMarketTypeLogic {
	return &AddMarketTypeLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *AddMarketTypeLogic) AddMarketType(req *types.AddMarketTypeReq) (resp *types.AddMarketTypeResp, err error) {
	result := l.svcCtx.DB.Exec("insert into market(`market_type`)values(?)", req.MarketType)
	if result.Error != nil {
		resp = &types.AddMarketTypeResp{
			CommonResp: types.CommonResp{
				Result:  common.FAILED,
				Message: fmt.Sprintf("error: %v", result.Error),
			},
		}
		// err = result.Error // 不注释就是返回400，错误信息为：Error 1062 (23000): Duplicate entry 'xxx' for key 'market.uc_market_type'
		return
	}

	resp = &types.AddMarketTypeResp{
		CommonResp: types.CommonResp{
			Result: common.SUCCESS,
		},
	}
	return
}
