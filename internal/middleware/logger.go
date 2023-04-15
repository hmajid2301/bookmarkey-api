package middleware

import (
	"net/http"

	"github.com/labstack/echo/v5"
	"github.com/rs/zerolog"
	"gitlab.com/bookmarkey/api/internal/log"
)

func logger(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		l := log.NewLogger()
		requestID := w.Header().Get(echo.HeaderXRequestID)
		l.UpdateContext(func(c zerolog.Context) zerolog.Context {
			return c.Str("request_id", requestID)
		})

		r = r.WithContext(l.WithContext(r.Context()))

		next.ServeHTTP(w, r)
	})
}
