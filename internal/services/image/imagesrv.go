package image

import (
	"context"
	"net"
	"time"

	"go.uber.org/fx"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/protobuf/types/known/timestamppb"

	"github.com/TiregeRRR/image_service/internal/controllers/image"
	"github.com/TiregeRRR/image_service/internal/model"
	imagev1 "github.com/TiregeRRR/image_service/proto/image/v1"
)

type In struct {
	fx.In

	LC              fx.Lifecycle
	Logger          *zap.Logger
	ImageController *image.Controller
}

type Service struct {
	imagev1.UnimplementedImageServiceServer

	ImageController *image.Controller
}

// TODO norm nado
func New(in In) {
	srv := &Service{
		ImageController: in.ImageController,
	}
	grpcServer := grpc.NewServer([]grpc.ServerOption{}...)
	in.LC.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			ln, err := net.Listen("tcp", "8000")
			if err != nil {
				return err
			}
			imagev1.RegisterImageServiceServer(grpcServer, srv)
			in.Logger.Info("Starting grpc server on  8000")
			return grpcServer.Serve(ln)
		},
		OnStop: func(ctx context.Context) error {
			grpcServer.Stop()
			return nil
		},
	})
}

func (s *Service) UploadImage(ctx context.Context, r *imagev1.UploadImageRequest) (*imagev1.UploadImageResponse, error) {
	t := time.Now()
	m, err := s.ImageController.UploadImage(ctx, &model.Image{
		Name:      r.GetName(),
		CreatedAt: t,
		UpdatedAt: t,
		Data:      r.GetData(),
	})
	if err != nil {
		return nil, err
	}
	return &imagev1.UploadImageResponse{
		Name:      m.Name,
		CreatedAt: timestamppb.New(m.CreatedAt),
		UpdatedAt: timestamppb.New(m.UpdatedAt),
		Data:      m.Data,
	}, nil
}
