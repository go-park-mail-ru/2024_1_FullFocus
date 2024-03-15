package middleware

import (
	"bufio"
	"context"
	"github.com/pkg/errors"
	"log/slog"
	"net"
	"net/http"
	"sync/atomic"

	"github.com/gorilla/mux"

	"github.com/go-park-mail-ru/2024_1_FullFocus/internal/pkg/logger"
)

var ErrHijackAssertion = errors.New("type assertion to http.Hijacker failed")

var currRequestID uint64

type responseWriterInterceptor struct {
	http.ResponseWriter
	statusCode int
}

func (w *responseWriterInterceptor) WriteHeader(statusCode int) {
	w.statusCode = statusCode
	w.ResponseWriter.WriteHeader(statusCode)
}

func (w *responseWriterInterceptor) Write(d []byte) (int, error) {
	return w.ResponseWriter.Write(d)
}

func (w *responseWriterInterceptor) Hijack() (net.Conn, *bufio.ReadWriter, error) {
	h, ok := w.ResponseWriter.(http.Hijacker)
	if !ok {
		return nil, nil, ErrHijackAssertion
	}
	return h.Hijack()
}

func NewLoggingMiddleware(l *slog.Logger) mux.MiddlewareFunc {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			atomic.AddUint64(&currRequestID, 1)
			reqGroup := slog.Group("request", slog.Uint64("requestID", currRequestID))
			requestLogger := l.With(reqGroup)
			requestLogger.Info("new request",
				slog.String("method", r.Method),
				slog.String("uri", r.RequestURI))
			wi := &responseWriterInterceptor{
				ResponseWriter: w,
				statusCode:     200,
			}
			ctx := logger.ContextWithLogger(context.Background(), requestLogger)
			next.ServeHTTP(wi, r.WithContext(ctx))
			requestLogger.Info("request processed", slog.Int("statusCode", wi.statusCode))
		})
	}
}
