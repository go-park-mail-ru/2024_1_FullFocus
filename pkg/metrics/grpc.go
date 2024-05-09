package metrics

import (
	"context"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type Middleware struct {
	metrics Collector
}

func NewMiddleware(metrics Collector) *Middleware {
	return &Middleware{
		metrics: metrics,
	}
}

func (m *Middleware) UnaryInterceptor(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (any, error) {
	start := time.Now()
	h, err := handler(ctx, req)
	dur := time.Since(start)
	statusCode := "200"
	if err != nil {
		st, ok := status.FromError(err)
		if !ok {
			statusCode = "500"
		} else {
			switch st.Code() {
			case codes.OK:
			case codes.InvalidArgument, codes.NotFound, codes.AlreadyExists:
				statusCode = "400"
			case codes.PermissionDenied:
				statusCode = "401"
			case codes.Internal:
				statusCode = "500"
			}
		}
		if statusCode != "200" {
			m.metrics.IncreaseErr(statusCode, info.FullMethod)
		}
	}
	m.metrics.AddDurationToSummary(statusCode, info.FullMethod, dur)
	m.metrics.AddDurationToHistogram(info.FullMethod, dur)
	m.metrics.IncreaseHits(info.FullMethod)
	return h, err
}
