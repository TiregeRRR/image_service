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
			if err := unaryLimiter.AddTask(); err != nil {
				return resp, err
			}
			defer unaryLimiter.DoneTask()
			return handler(ctx, req)
		}),
		grpc.StreamInterceptor(func(srv interface{}, ss grpc.ServerStream, info *grpc.StreamServerInfo, handler grpc.StreamHandler) error {
			if err := streamLimiter.AddTask(); err != nil {
				return err
			}
			defer streamLimiter.DoneTask()
			return handler(srv, ss)
		}),
	}
}
