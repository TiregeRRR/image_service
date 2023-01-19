package imagefile

import (
	"github.com/TiregeRRR/image_service/internal/model"
	"go.uber.org/fx"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

type In struct {
	fx.In

	DB     *gorm.DB
	Logger *zap.Logger
}

type Manager struct {
	logger *zap.Logger
	db     *gorm.DB
}

func New(in In) *Manager {
	return &Manager{
		logger: in.Logger,
		db:     in.DB,
	}
}

func (mgr *Manager) Create(m *model.Image) error {
	mgr.logger.Info("insert image", zap.Any("model", m))
	if r := mgr.db.Create(m); r.Error != nil {
		mgr.logger.Error("insert image", zap.Error(r.Error))
	}
	return nil
}
