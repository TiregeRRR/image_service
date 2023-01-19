package db

import (
	"github.com/TiregeRRR/image_service/internal/pkg/config"
	"go.uber.org/fx"
)

type In struct {
	fx.In

	Config config.Config
}

func New(in In) *