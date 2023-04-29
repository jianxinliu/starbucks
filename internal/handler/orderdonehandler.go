package handler

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"starbucks/internal/logic"
	"starbucks/internal/svc"
	"starbucks/internal/types"
)

func OrderDoneHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.OrderDoneReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := logic.NewOrderDoneLogic(r.Context(), svcCtx)
		resp, err := l.OrderDone(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
