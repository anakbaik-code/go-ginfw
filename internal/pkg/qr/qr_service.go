package qr

import "github.com/skip2/go-qrcode"

type Service struct{}

func NewService() *Service {
	return &Service{}
}

func (s *Service) GenerateImage(data string) ([]byte, error) {
	return qrcode.Encode(data, qrcode.Medium, 256)
}