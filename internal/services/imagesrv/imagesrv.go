package imagesrv

import (
	"fmt"
	"io"

	"go.uber.org/fx"
	"go.uber.org/zap"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/TiregeRRR/image_service/internal/controllers/imagefile"
	imagev1 "github.com/TiregeRRR/image_service/proto/image/v1"
)

type In struct {
	fx.In

	LC              fx.Lifecycle
	Logger          *zap.Logger
	ImageController *imagefile.Controller
}

type Service struct {
	imagev1.UnimplementedImageServiceServer

	Logger          *zap.Logger
	ImageController *imagefile.Controller
}

func New(in In) *Service {
	srv := &Service{
		ImageController: in.ImageController,
		Logger:          in.Logger,
	}
	return srv
}

func (s *Service) UploadImage(str imagev1.ImageService_UploadImageServer) error {
	d, err := str.Recv()
	if err != nil {
		return err
	}
	if d.GetName() == "" {
		return fmt.Errorf("first chunk must be name")
	}
	data := []byte{}
	for {
		d, err := str.Recv()
		if status.Code(err) == codes.Canceled {
			s.Logger.Info("connection context closed")
			break
		}
		if err == io.EOF {
			s.Logger.Debug("eof")
			break
		}
		if err != nil {
			s.Logger.Error("Receive msg", zap.Error(err))
			return err
		}
		if d.GetChunk() == nil {
			s.Logger.Error("GET CHUNK", zap.Error(err))
			return fmt.Errorf("chunk not provided")
		}
		data = append(data, d.GetChunk()...)
	}
	return s.ImageController.UploadImage(d.GetName(), data)
}
