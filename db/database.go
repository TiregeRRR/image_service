package db

import (
	"fmt"

	"github.com/TiregeRRR/image_service/internal/pkg/config"
	"go.uber.org/fx"
	"go.uber.org/zap"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type In struct {
	fx.In

	Config config.Config
	Logger *zap.Logger
}

func New(in In) (*gorm.DB, error) {
	url := generateUrl(in.Config)
	db, err := gorm.Open(
		postgres.Open(url),
		&gorm.Config{},
	)
	if err != nil {
		return nil, err
	}
	migrate(db)
	return db, nil
}

func generateUrl(conf config.Config) string {
	return fmt.Sprintf(
		"%s://%s:%s@%s:%s/%s",
		"postgresql",
		conf.PostgresUser,
		conf.PostgresPassword,
		conf.PostgresHost,
		conf.PostgresPort,
		conf.PostgresDB,
	)
}
