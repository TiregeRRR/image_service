package middleware

import (
	"context"

	"github.com/TiregeRRR/image_service/pkg/limiter"
	"google.golang.org/grpc"
)

func NewGRPCRatelimitInterceptor() []grpc.ServerOption {
	unaryLimiter := limiter.NewLimiter(100)
	streamLimiter := limiter.NewLimiter(10)
	return []grpc.ServerOption{
		grpc.UnaryInterceptor(func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp interface{}, err error) {
			unaryLimiter.AddTask()
			defer unaryLimiter.DoneTask()
			return handler(ctx, req)
		}),
		grpc.StreamInterceptor(func(srv interface{}, ss grpc.ServerStream, info *grpc.StreamServerInfo, handler grpc.StreamHandler) error {
			streamLimiter.AddTask()
			defer streamLimiter.DoneTask()
			return handler(srv, ss)
		}),
	}
}
