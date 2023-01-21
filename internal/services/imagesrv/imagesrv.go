package imagesrv

import (
	"context"
	"errors"
	"io"

	"go.uber.org/fx"
	"go.uber.org/zap"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/timestamppb"

	"github.com/TiregeRRR/image_service/internal/controllers/imagefile"
	"github.com/TiregeRRR/image_service/internal/model"
	imagev1 "github.com/TiregeRRR/image_service/proto/image/v1"
	"github.com/golang/protobuf/ptypes/empty"
)

var (
	errFirstChunkNotName = errors.New("first chunk must be name")
	errNoChunkProvided   = errors.New("chunk not provided")
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
		return errFirstChunkNotName
	}
	data := []byte{}
	for {
		d, err := str.Recv()
		if status.Code(err) == codes.Canceled {
			s.Logger.Info("connection context closed")
			return err
		}
		if err == io.EOF {
			break
		}
		if err != nil {
			s.Logger.Error("Receive msg", zap.Error(err))
			return err
		}
		if d.GetChunk() == nil {
			return errNoChunkProvided
		}
		data = append(data, d.GetChunk()...)
	}
	m, err := s.ImageController.UploadImage(d.GetName(), data)
	if err != nil {
		return err
	}
	str.SendAndClose(&imagev1.UploadImageResponse{
		Image: imageModelToProto(m),
	})
	return nil
}

func (s *Service) GetImages(context.Context, *empty.Empty) (*imagev1.GetImagesResponse, error) {
	sl, err := s.ImageController.GetImages()
	if err != nil {
		return nil, err
	}
	return &imagev1.GetImagesResponse{
		Image: imageModelSliceToProto(sl),
	}, nil
}

func (s *Service) DownloadImage(r *imagev1.DownloadImageRequest, stream imagev1.ImageService_DownloadImageServer) error {
	m, err := s.ImageController.DownloadImage(r.GetName())
	if err != nil {
		return err
	}
	chunkSize := 1024
	data := splitChunks(m.Data, chunkSize)
	for _, c := range data {
		if err := stream.Send(&imagev1.DownloadImageResponse{
			Chunk: c,
		}); err != nil {
			return err
		}
	}
	return nil
}

func imageModelSliceToProto(m []*model.Image) []*imagev1.Image {
	sl := []*imagev1.Image{}
	for _, v := range m {
		sl = append(sl, imageModelToProto(v))
	}
	return sl
}

func imageModelToProto(m *model.Image) *imagev1.Image {
	return &imagev1.Image{
		Name:      m.Name,
		CreatedAt: timestamppb.New(m.CreatedAt),
		UpdatedAt: timestamppb.New(m.UpdatedAt),
	}
}

func splitChunks(buf []byte, lim int) [][]byte {
	var chunk []byte
	chunks := make([][]byte, 0, len(buf)/lim+1)
	for len(buf) >= lim {
		chunk, buf = buf[:lim], buf[lim:]
		chunks = append(chunks, chunk)
	}
	if len(buf) > 0 {
		chunks = append(chunks, buf[:])
	}
	return chunks
}
