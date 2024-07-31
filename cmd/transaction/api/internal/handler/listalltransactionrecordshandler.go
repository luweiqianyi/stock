package handler

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"stock/cmd/transaction/api/internal/logic"
	"stock/cmd/transaction/api/internal/svc"
	"stock/cmd/transaction/api/internal/types"
)

func ListAllTransactionRecordsHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.ListAllTransactionRecordsReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := logic.NewListAllTransactionRecordsLogic(r.Context(), svcCtx)
		resp, err := l.ListAllTransactionRecords(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
