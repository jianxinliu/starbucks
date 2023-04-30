package handler

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"starbucks/starbucks/internal/logic"
	"starbucks/starbucks/internal/svc"
)

func DescribeWalletHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		l := logic.NewDescribeWalletLogic(r.Context(), svcCtx)
		resp, err := l.DescribeWallet()
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
