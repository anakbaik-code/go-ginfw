package attendance

import (
	"context"
	"database/sql"
	"go-fwgin/internal/database"
	"go-fwgin/internal/pkg/jwt"
)

type ServiceAttendance interface {
	CheckIn(ctx context.Context, req CheckInRequest) error
	CheckInByQR(ctx context.Context, req CheckInByQRRequest) error
	ManualCheckIn(ctx context.Context, req ManualCheckInRequest) error

	GetAttendanceByOrderID(ctx context.Context, orderID uint64) (AttendanceDetailResponse, error)
	GetAttendanceByEventID(ctx context.Context, eventID uint64) ([]AttendanceByEventResponse, error)
	GetAttendanceByUserID(ctx context.Context, userID uint64) ([]AttendanceByUserResponse, error)

	CountAttendeesByEvent(ctx context.Context, eventID uint64) (int64, error)
	CountAttendeesByOrganizer(ctx context.Context, organizerID uint64) (int64, error)
	CountAttendeesByTicketType(ctx context.Context, ticketTypeID uint64) (int64, error)
	GetEventAttendanceRate(ctx context.Context, organizerID uint64) ([]EventAttendanceRateResponse, error)
	GetCheckInTimeline(ctx context.Context, eventID uint64) ([]CheckInTimelineHourResponse, error)
	GetCheckInTimelineByDate(ctx context.Context, eventID uint64) ([]CheckInTimelineDateResponse, error)

	CancelAttendance(ctx context.Context, req CancelAttendanceRequest) error
	SoftDeleteAttendance(ctx context.Context, id uint64) error
}

type serviceAttendance struct {
	repo      RepositoryAttendance
	validator *AttendanceValidator
	jwtService *jwt.Service
}

func NewServiceAttendance(repo RepositoryAttendance,jwtService *jwt.Service) ServiceAttendance {
	return &serviceAttendance{
		repo:      repo,
		validator: NewAttendanceValidator(repo),
		jwtService: jwtService,
	}
}

// ============================================
// CREATE / CHECK-IN
// ============================================

func (s *serviceAttendance) CheckIn(ctx context.Context, req CheckInRequest) error {
	if err := s.validator.ValidateCheckIn(ctx, req.OrderID, req.EventID); err != nil {
		return err
	}

	return s.repo.CheckIn(ctx, database.CheckInParams{
		OrderID: req.OrderID,
		EventID: req.EventID,
		UserID: req.UserID,
		TicketTypeID: req.TicketTypeID,
		CheckInMethod: database.NullAttendancesCheckInMethod{
			AttendancesCheckInMethod: database.AttendancesCheckInMethodManual,
			Valid: true,
		},
		CheckedBy: sql.NullInt64{Int64: int64(req.CheckedBy), Valid: true},
	})
}

func (s *serviceAttendance) CheckInByQR(ctx context.Context, req CheckInByQRRequest) error {
	claims,err := s.jwtService.ValidateTicketToken(req.QRToken)
	if err != nil {
		return ErrInvalidQRToken
	}
	detail ,err :=s.repo.GetOrderItemDetail(ctx,claims.TicketID)
	if err != nil {
		return err
	}
	
	if err := s.validator.ValidateCheckIn(ctx, detail.OrderID, detail.EventID); err != nil {
		return err
	}

	return s.repo.CheckInByQR(ctx, database.CheckInByQRParams{
		OrderID:      detail.OrderID,
		EventID:      detail.EventID,
		UserID:       detail.UserID,
		TicketTypeID: detail.TicketTypeID,
		CheckedBy:   sql.NullInt64{Int64: int64(req.CheckedBy),Valid: true},
	})
}

func (s *serviceAttendance) ManualCheckIn(ctx context.Context, req ManualCheckInRequest) error {
	if err := s.validator.ValidateCheckIn(ctx, req.OrderID, req.EventID); err != nil {
		return err
	}

	return s.repo.ManualCheckIn(ctx, database.ManualCheckInParams{
		OrderID:      req.OrderID,
		EventID:      req.EventID,
		UserID:       req.UserID,
		TicketTypeID: req.TicketTypeID,
		CheckedBy:    sql.NullInt64{Int64: int64(req.CheckedBy),Valid: true},
	})
}

// ============================================
// READ
// ============================================

func (s *serviceAttendance) GetAttendanceByOrderID(ctx context.Context, orderID uint64) (AttendanceDetailResponse, error) {
	result, err := s.repo.GetAttendanceByOrderID(ctx, orderID)
	if err != nil {
		return AttendanceDetailResponse{}, err
	}
	return ToAttendanceDetailResponse(result), nil
}

func (s *serviceAttendance) GetAttendanceByEventID(ctx context.Context, eventID uint64) ([]AttendanceByEventResponse, error) {
	result, err := s.repo.GetAttendanceByEventID(ctx, eventID)
	if err != nil {
		return nil, err
	}
	return ToAttendanceByEventResponses(result), nil
}

func (s *serviceAttendance) GetAttendanceByUserID(ctx context.Context, userID uint64) ([]AttendanceByUserResponse, error) {
	result, err := s.repo.GetAttendanceByUserID(ctx, userID)
	if err != nil {
		return nil, err
	}
	return ToAttendanceByUserResponses(result), nil
}

// ============================================
// STATISTICS
// ============================================

func (s *serviceAttendance) CountAttendeesByEvent(ctx context.Context, eventID uint64) (int64, error) {
	return s.repo.CountAttendeesByEvent(ctx, eventID)
}

func (s *serviceAttendance) CountAttendeesByOrganizer(ctx context.Context, organizerID uint64) (int64, error) {
	return s.repo.CountAttendeesByOrganizer(ctx, organizerID)
}

func (s *serviceAttendance) CountAttendeesByTicketType(ctx context.Context, ticketTypeID uint64) (int64, error) {
	return s.repo.CountAttendeesByTicketType(ctx, ticketTypeID)
}

func (s *serviceAttendance) GetEventAttendanceRate(ctx context.Context, organizerID uint64) ([]EventAttendanceRateResponse, error) {
	result, err := s.repo.GetEventAttendanceRate(ctx, organizerID)
	if err != nil {
		return nil, err
	}
	return ToEventAttendanceRateResponses(result), nil
}

func (s *serviceAttendance) GetCheckInTimeline(ctx context.Context, eventID uint64) ([]CheckInTimelineHourResponse, error) {
	result, err := s.repo.GetCheckInTimeline(ctx, eventID)
	if err != nil {
		return nil, err
	}
	return ToCheckInTimelineHourResponses(result), nil
}

func (s *serviceAttendance) GetCheckInTimelineByDate(ctx context.Context, eventID uint64) ([]CheckInTimelineDateResponse, error) {
	result, err := s.repo.GetCheckInTimelineByDate(ctx, eventID)
	if err != nil {
		return nil, err
	}
	return ToCheckInTimelineDateResponses(result), nil
}

// ============================================
// UPDATE / DELETE
// ============================================

func (s *serviceAttendance) CancelAttendance(ctx context.Context, req CancelAttendanceRequest) error {
	return s.repo.CancelAttendance(ctx, req.OrderID, req.EventID)
}

func (s *serviceAttendance) SoftDeleteAttendance(ctx context.Context, id uint64) error {
	return s.repo.SoftDeleteAttendance(ctx, id)
}