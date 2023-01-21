package imagefile

import (
	"errors"
	"sync"

	"github.com/TiregeRRR/image_service/internal/managers/imagefile"
	"github.com/TiregeRRR/image_service/internal/model"
	"github.com/TiregeRRR/image_service/internal/pkg/config"
	"github.com/TiregeRRR/image_service/internal/pkg/storage"
	"go.uber.org/fx"
	"go.uber.org/zap"
)

var errFileIsBusy = errors.New("file is busy")

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

func (c *Controller) UploadImage(name string, data []byte) (*model.Image, error) {
	if _, ok := c.busyFiles.Load(name); ok {
		return nil, errFileIsBusy
	}
	c.busyFiles.Store(name, struct{}{})
	defer c.busyFiles.Delete(name)
	m, err := c.imageManager.Upsert(&model.Image{
		Name: name,
		Data: data,
	})
	if err != nil {
		return nil, err
	}
	return m, nil
}

func (c *Controller) GetImages() ([]*model.Image, error) {
	return c.imageManager.GetImages()
}

func (c *Controller) DownloadImage(name string) (*model.Image, error) {
	if _, ok := c.busyFiles.Load(name); ok {
		return nil, errFileIsBusy
	}
	c.busyFiles.Store(name, struct{}{})
	defer c.busyFiles.Delete(name)
	m, err := c.imageManager.GetImageData(name)
	if err != nil {
		return nil, err
	}
	return m, nil
}
