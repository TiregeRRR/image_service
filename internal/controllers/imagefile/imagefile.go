package imagefile

import (
	"context"

	"github.com/TiregeRRR/image_service/internal/model"
)

type Controller struct{}

func New() *Controller {
	return &Controller{}
}

func (c *Controller) UploadImage(ctx context.Context, m *model.Image) (*model.Image, error) {

	return m, nil
}
