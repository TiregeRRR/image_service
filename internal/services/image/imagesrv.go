package image

import imagev1 "github.com/TiregeRRR/tager_test/proto/image/v1"

type Service struct {
	imagev1.UnimplementedImageServiceServer
}

func (s *Service) UploadImage(stream imagev1.ImageService_UploadImageServer) error {
	return nil
}
