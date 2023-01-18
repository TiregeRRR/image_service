package image

import (
	"context"

	"github.com/TiregeRRR/tager_test/internal/model"
)

type Controller struct{}

func New() *Controller {
	return &Controller{}
}

func (c *Controller) UploadImage(ctx context.Context, m *model.Image) (*model.Image, error) {

	return nil, nil
}
