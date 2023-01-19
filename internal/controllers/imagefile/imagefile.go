package imagefile

import (
	"fmt"
	"sync"

	"github.com/TiregeRRR/image_service/internal/managers/imagefile"
	"github.com/TiregeRRR/image_service/internal/model"
	"github.com/TiregeRRR/image_service/internal/pkg/config"
	"github.com/TiregeRRR/image_service/internal/pkg/storage"
	"go.uber.org/fx"
	"go.uber.org/zap"
)

type In struct {
	fx.In

	Logger       *zap.Logger
	Config       config.Config
	ImageManager *imagefile.Manager
}

type Controller struct {
	logger       *zap.Logger
	storage      *storage.DiskStorage
	imageManager *imagefile.Manager
	busyFiles    sync.Map
}

func New(in In) *Controller {
	return &Controller{
		logger:       in.Logger,
		storage:      storage.New(in.Config.StorageDir),
		imageManager: in.ImageManager,
	}
}

func (c *Controller) UploadImage(name string, data []byte) error {
	if _, ok := c.busyFiles.Load(name); ok {
		return fmt.Errorf("is busy")
	}
	c.busyFiles.Store("name", struct{}{})
	path, err := c.storage.Save(name, data)
	if err != nil {
		return err
	}
	err = c.imageManager.Create(&model.Image{
		Name: name,
		Path: path,
	})
	if err != nil {
		return err
	}
	return nil
}
