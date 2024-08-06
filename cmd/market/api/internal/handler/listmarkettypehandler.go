package handler

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"stock/cmd/market/api/internal/logic"
	"stock/cmd/market/api/internal/svc"
	"stock/cmd/market/api/internal/types"
)

func ListMarketTypeHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.ListMarketTypeReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := logic.NewListMarketTypeLogic(r.Context(), svcCtx)
		resp, err := l.ListMarketType(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
