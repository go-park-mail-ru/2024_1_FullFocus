package logger

import (
	"context"
	"log/slog"

	"github.com/grpc-ecosystem/go-grpc-middleware/v2/interceptors/logging"
	"google.golang.org/grpc"
)

func InterceptorLogger(l *slog.Logger) logging.Logger {
	return logging.LoggerFunc(func(ctx context.Context, lvl logging.Level, msg string, fields ...any) {
		l.Log(ctx, slog.Level(lvl), msg, fields...)
	})
}

// UnaryInterceptor for gRPC server that injects logger into context
func UnaryInterceptor(logger *slog.Logger) grpc.UnaryServerInterceptor {
	return func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
		ctx = WithContext(ctx, logger)
		return handler(ctx, req)
	}
}

// StreamInterceptor for gRPC server that injects logger into context
func StreamInterceptor(logger *slog.Logger) grpc.StreamServerInterceptor {
	return func(srv interface{}, stream grpc.ServerStream, info *grpc.StreamServerInfo, handler grpc.StreamHandler) error {
		ctx := WithContext(stream.Context(), logger)
		return handler(srv, &wrappedStream{ctx, stream})
	}
}

// wrappedStream wraps around the original ServerStream, embedding the new context
type wrappedStream struct {
	ctx context.Context
	grpc.ServerStream
}

func (w *wrappedStream) Context() context.Context {
	return w.ctx
}
