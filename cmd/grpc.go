package main

import (
	"context"
	"fmt"
	"net"

	"github.com/TiregeRRR/image_service/internal/pkg/config"
	"github.com/TiregeRRR/image_service/internal/pkg/middleware"
	"go.uber.org/fx"
	"go.uber.org/zap"
	"google.golang.org/grpc"
)

func newGRPCSrv(
	lc fx.Lifecycle,
	logger *zap.Logger,
	config config.Config,
) *grpc.Server {
	srv := grpc.NewServer(middleware.NewGRPCRatelimitInterceptor()...)
	lc.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			listener, err := net.Listen("tcp", fmt.Sprintf("%s:%s", config.ServerHost, config.ServerPort))
			if err != nil {
				return err
			}
			logger.Info("Starting grpc server at: ", zap.String("", fmt.Sprintf("%s:%s", config.ServerHost, config.ServerPort)))
			go srv.Serve(listener)
			return nil
		},
		OnStop: func(ctx context.Context) error {
			srv.Stop()
			return nil
		},
	})
	return srv
}
