package main

import (
	"go.uber.org/fx"
	"go.uber.org/fx/fxevent"
	"go.uber.org/zap"
	"google.golang.org/grpc"

	"github.com/TiregeRRR/image_service/internal/services/image"
)

func main() {
	fx.New(
		fx.WithLogger(func(logger *zap.Logger) fxevent.Logger {
			return &fxevent.ZapLogger{
				Logger: logger,
			}
		}),
		fx.Provide(image.New),
		fx.Invoke(func(
			*grpc.Server,
		) {
		}),
	).Run()
}
