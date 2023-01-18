package logger

import (
	"context"

	"go.uber.org/fx"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func New(lc fx.Lifecycle) (*zap.Logger, error) {
	config := zap.NewDevelopmentConfig()
	config.Development = false
	config.EncoderConfig.EncodeLevel = zapcore.LowercaseColorLevelEncoder
	logger, err := config.Build()
	if err != nil {
		return nil, err
	}

	lc.Append(fx.Hook{
		OnStop: func(ctx context.Context) error {
			_ = logger.Sync()
			return nil
		},
	})

	return logger, nil
}
