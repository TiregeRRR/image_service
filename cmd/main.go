package main

import (
	"fmt"

	"go.uber.org/fx"
	"go.uber.org/fx/fxevent"
	"go.uber.org/zap"
	"google.golang.org/grpc"

	"github.com/TiregeRRR/image_service/internal/controllers/imagefile"
	"github.com/TiregeRRR/image_service/internal/services"
	"github.com/TiregeRRR/image_service/pkg/logger"
)

func main() {
	conf, err := LoadConfig("./configs/")
	if err != nil {
		panic(fmt.Errorf("can't load config: %w", err))
	}
	fx.New(
		fx.WithLogger(func(lgr *zap.Logger) fxevent.Logger {
			return &fxevent.ZapLogger{Logger: lgr}
		}),
		fx.Supply(*conf),
		fx.Provide(
			logger.New,
			imagefile.New,
			newGRPCSrv,
		),

		services.Module,

		fx.Invoke(func(
			*grpc.Server,
		) {
		}),
	).Run()
}
