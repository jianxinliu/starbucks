package middleware

import (
	"context"
	"net/http"
	"starbucks/starbucks/utils"
)

type CustomJwtMiddleware struct {
}

func NewCustomJwtMiddleware() *CustomJwtMiddleware {
	return &CustomJwtMiddleware{}
}

func (m *CustomJwtMiddleware) Handle(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if len(r.Header.Get("Authorization")) > 0 {

			resp, err := utils.UserAuth(&utils.UserAuthRequest{
				Token: r.Header.Get("Authorization"),
			})
			if err != nil {
				w.WriteHeader(http.StatusUnauthorized)
				return
			}
			ctx := r.Context()

			ctx = context.WithValue(ctx, "userId", resp.UserId)
			next(w, r.WithContext(ctx))
		} else {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}
	}
}
