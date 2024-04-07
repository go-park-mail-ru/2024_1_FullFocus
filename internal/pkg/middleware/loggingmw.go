package middleware

import (
	"bufio"
	"context"
	"log/slog"
	"net"
	"net/http"
	"sync/atomic"
	"time"

	"github.com/gorilla/mux"
	"github.com/pkg/errors"

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
			requestLogger := l.With(slog.Uint64("requestID", currRequestID))
			requestLogger.Info("new",
				slog.String("method", r.Method),
				slog.String("uri", r.RequestURI))
			wi := &responseWriterInterceptor{
				ResponseWriter: w,
				statusCode:     200,
			}
			ctx := logger.WithContext(context.Background(), requestLogger)
			start := time.Now()
			next.ServeHTTP(wi, r.WithContext(ctx))
			requestLogger.Info("response",
				slog.Int("statusCode", wi.statusCode),
				slog.String("duration", time.Since(start).String()))
		})
	}
}
