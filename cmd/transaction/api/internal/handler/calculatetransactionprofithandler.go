package handler

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"stock/cmd/transaction/api/internal/logic"
	"stock/cmd/transaction/api/internal/svc"
	"stock/cmd/transaction/api/internal/types"
)

func CalculateTransactionProfitHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.CalculateTransactionProfitReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := logic.NewCalculateTransactionProfitLogic(r.Context(), svcCtx)
		resp, err := l.CalculateTransactionProfit(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
