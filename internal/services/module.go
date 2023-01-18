package services

import (
	imagefile "github.com/TiregeRRR/image_service/internal/services/imagesrv"
	imagev1 "github.com/TiregeRRR/image_service/proto/image/v1"
	"go.uber.org/fx"
	"google.golang.org/grpc"
)

var Module = fx.Options(
	imagefile.Module,

	fx.Invoke(func(
		grpcServer *grpc.Server,
		imageService *imagefile.Service,
	) {
		imagev1.RegisterImageServiceServer(grpcServer, imageService)
	}),
)
