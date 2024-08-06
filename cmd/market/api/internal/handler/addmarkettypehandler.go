package handler

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"stock/cmd/market/api/internal/logic"
	"stock/cmd/market/api/internal/svc"
	"stock/cmd/market/api/internal/types"
)

func AddMarketTypeHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.AddMarketTypeReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := logic.NewAddMarketTypeLogic(r.Context(), svcCtx)
		resp, err := l.AddMarketType(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
