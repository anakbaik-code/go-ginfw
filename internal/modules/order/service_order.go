type serviceOrder struct {
	repo       RepositoryOrder
	jwtService *jwt.Service
	qrService  *qr.Service
}

func NewServiceOrder(repo RepositoryOrder, jwtService *jwt.Service, qrService *qr.Service) ServiceOrder {
	return &serviceOrder{
		repo:       repo,
		jwtService: jwtService,
		qrService:  qrService,
	}
}

func (s *serviceOrder) GenerateTicketQR(ctx context.Context, ticketID uint64) ([]byte, error) {
	token, err := s.jwtService.GenerateTicketToken(ticketID, 30*24*time.Hour)
	if err != nil {
		return nil, err
	}
	return s.qrService.GenerateImage(token)
}