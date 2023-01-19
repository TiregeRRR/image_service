package imagefile

import (
	"fmt"
	"sync"

	"github.com/TiregeRRR/image_service/internal/pkg/config"
	"github.com/TiregeRRR/image_service/internal/pkg/storage"
	"go.uber.org/fx"
	"go.uber.org/zap"
)

type In struct {
	fx.In

	Logger *zap.Logger
	Config config.Config
}

type Controller struct {
	logger    *zap.Logger
	storage   *storage.DiskStorage
	busyFiles sync.Map
}

func New(in In) *Controller {
	return &Controller{
		logger:  in.Logger,
		storage: storage.New(in.Config.StorageDir),
	}
}

func (c *Controller) UploadImage(name string, data []byte) error {
	if _, ok := c.busyFiles.Load(name); ok {
		return fmt.Errorf("is busy")
	}
	c.busyFiles.Store("name", struct{}{})
	if err := c.storage.Save(name, data); err != nil {
		return err
	}
	c.logger.Info("image saved")
	return nil
}
