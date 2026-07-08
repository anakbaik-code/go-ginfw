package jwt

type Service struct {
	secret []byte
}

func New(secret string) *Service {
	return &Service{
		secret: []byte(secret),
	}
}