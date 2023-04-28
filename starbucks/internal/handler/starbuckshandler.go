package handler

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"starbucks/starbucks/internal/logic"
	"starbucks/starbucks/internal/svc"
	"starbucks/starbucks/internal/types"
)

func StarbucksHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.Request
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := logic.NewStarbucksLogic(r.Context(), svcCtx)
		resp, err := l.Starbucks(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
