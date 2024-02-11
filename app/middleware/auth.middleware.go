package middleware

import (
	"context"
	"net/http"

	"github.com/developertom01/post-jsonrpc-server/config"
	"github.com/developertom01/post-jsonrpc-server/utils"
)

func AuthenticationMiddleware() func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			authHeaders, ok := r.Header[config.AUTH_TOKEN_HEADER_KEY]
			if !ok {
				next.ServeHTTP(w, r)
				return
			}

			bearerToken := utils.ExtractBearerToken(authHeaders[0])
			userId, err := utils.ParseJwtToken(bearerToken, config.ACCESS_TOKEN_SECRET)
			if err != nil {
				next.ServeHTTP(w, r)
				return
			}

			ctx := context.WithValue(r.Context(), config.AUTH_CONTEXT_KEY, userId)
			r = r.WithContext(ctx)
			next.ServeHTTP(w, r)
		})
	}
}
