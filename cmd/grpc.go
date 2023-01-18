package main

import (
	"context"
	"net"

	"github.com/TiregeRRR/image_service/internal/pkg/config"
	"go.uber.org/fx"
	"go.uber.org/zap"
	"google.golang.org/grpc"
)

type In struct {
	LC fx.Lifecycle

	Logger *zap.Logger
	Config config.Config
}

func newGRPCSrv(
	lc fx.Lifecycle,
	logger *zap.Logger,
	config config.Config,
) *grpc.Server {
	srv := grpc.NewServer()
	lc.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			listener, err := net.Listen("tcp", config.ServerAddress)
			if err != nil {
				return err
			}
			logger.Info("Starting grpc server at: ", zap.String("", config.ServerAddress))
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
