package attendance

import (
	"context"
	"database/sql"
	"go-fwgin/internal/database"
	"time"
)

type RepositoryAttendance interface {
	// Create / Check-in
	CheckIn(ctx context.Context, arg database.CheckInParams) error
	CheckInByQR(ctx context.Context, arg database.CheckInByQRParams) error
	ManualCheckIn(ctx context.Context, arg database.ManualCheckInParams) error

	// Read
	GetAttendanceByOrderID(ctx context.Context, orderID uint64) (AttendanceDetail, error)
	GetAttendanceByEventID(ctx context.Context, eventID uint64) ([]AttendanceByEvent, error)
	GetAttendanceByUserID(ctx context.Context, userID uint64) ([]AttendanceByUser, error)

	// Statistics
	CountAttendeesByEvent(ctx context.Context, eventID uint64) (int64, error)
	CountAttendeesByOrganizer(ctx context.Context, organizerID uint64) (int64, error)
	CountAttendeesByTicketType(ctx context.Context, ticketTypeID uint64) (int64, error)
	GetEventAttendanceRate(ctx context.Context, organizerID uint64) ([]EventAttendanceRate, error)
	GetCheckInTimeline(ctx context.Context, eventID uint64) ([]CheckInTimelineHour, error)
	GetCheckInTimelineByDate(ctx context.Context, eventID uint64) ([]CheckInTimelineDate, error)

	// Update / Delete
	CancelAttendance(ctx context.Context, orderID uint64, eventID uint64) error
	SoftDeleteAttendance(ctx context.Context, id uint64) error

	// Validation
	CheckAlreadyCheckedIn(ctx context.Context, orderID uint64, eventID uint64) (bool, error)
	GetOrderStatus(ctx context.Context, orderID uint64) (string, error)
	GetEventTimeRange(ctx context.Context, eventID uint64) (start time.Time, end time.Time, err error)
	GetOrderItemDetail(ctx context.Context, ticketID uint64) (*OrderItemDetail, error)
}

type repositoryAttendance struct {
	db      *sql.DB
	queries *database.Queries
}

func NewRepositoryAttendance(queries *database.Queries) RepositoryAttendance {
	return &repositoryAttendance{queries: queries}
}

// CREATE/CHECKIN
func (r *repositoryAttendance) CheckIn(ctx context.Context, arg database.CheckInParams) error {
	tx, err := r.db.BeginTx(ctx, nil)
	if err != nil {
		return err
	}
	defer tx.Rollback()

	qtx := r.queries.WithTx(tx)

	// Lock existing row (kalau ada) biar request lain menunggu
	alreadyCheckedIn, err := qtx.CheckAlreadyCheckedInForUpdate(ctx, database.CheckAlreadyCheckedInForUpdateParams{
		OrderID: arg.OrderID,
		EventID: arg.EventID,
	})
	if err != nil && err != sql.ErrNoRows {
		return err
	}
	if alreadyCheckedIn {
		return err
	}

	if err := qtx.CheckIn(ctx, database.CheckInParams{
		OrderID:       arg.OrderID,
		EventID:       arg.EventID,
		UserID:        arg.UserID,
		TicketTypeID:  arg.TicketTypeID,
		CheckInMethod: arg.CheckInMethod,
		CheckedBy:     sql.NullInt64{Int64: arg.CheckedBy.Int64, Valid: arg.CheckedBy.Int64 != 0},
	}); err != nil {
		return err
	}

	return tx.Commit()
}

func (r *repositoryAttendance) CheckInByQR(ctx context.Context, arg database.CheckInByQRParams) error {
	return r.queries.CheckInByQR(ctx, database.CheckInByQRParams{
		OrderID:      arg.OrderID,
		EventID:      arg.EventID,
		UserID:       arg.UserID,
		TicketTypeID: arg.TicketTypeID,
		CheckedBy:    sql.NullInt64{Int64: arg.CheckedBy.Int64, Valid: arg.CheckedBy.Int64 != 0},
	})
}

func (r *repositoryAttendance) ManualCheckIn(ctx context.Context, arg database.ManualCheckInParams) error {
	return r.queries.ManualCheckIn(ctx, database.ManualCheckInParams{
		OrderID:      arg.OrderID,
		EventID:      arg.EventID,
		UserID:       arg.UserID,
		TicketTypeID: arg.TicketTypeID,
		CheckedBy:    sql.NullInt64{Int64: arg.CheckedBy.Int64, Valid: arg.CheckedBy.Int64 != 0},
	})
}

// READ
func (r *repositoryAttendance) GetAttendanceByOrderID(ctx context.Context, orderID uint64) (AttendanceDetail, error) {
	row, err := r.queries.GetAttendanceByOrderID(ctx, orderID)
	if err != nil {
		return AttendanceDetail{}, err
	}

	return AttendanceDetail{
		Attendance: Attendance{
			ID:            uint64(row.ID),
			OrderID:       row.OrderID,
			EventID:       row.EventID,
			UserID:        row.UserID,
			TicketTypeID:  row.TicketTypeID,
			CheckInTime:   nullTimeToPtr(row.CheckInTime),
			CheckInMethod: row.CheckInMethod.AttendancesCheckInMethod,
			Status:        row.Status.AttendancesStatus,
			CheckedBy:     NullInt64ToUint64Ptr(row.CheckedBy),
			CreatedAt:     row.CheckInTime.Time,
			UpdatedAt:     row.UpdatedAt.Time,
		},
		UserName:       row.UserName,
		UserEmail:      row.UserEmail,
		EventTitle:     row.EventTitle,
		TicketTypeName: row.TicketTypeName,
		CheckedByName:  nullStringToPtr(row.CheckedByName),
	}, nil
}

func (r *repositoryAttendance) GetAttendanceByEventID(ctx context.Context, eventID uint64) ([]AttendanceByEvent, error) {
	rows, err := r.queries.GetAttendanceByEventID(ctx, eventID)
	if err != nil {
		return nil, err
	}

	result := make([]AttendanceByEvent, 0, len(rows))
	for _, row := range rows {
		result = append(result, AttendanceByEvent{
			Attendance: Attendance{
				ID:            uint64(row.ID),
				OrderID:       row.OrderID,
				UserID:        row.UserID,
				CheckInTime:   nullTimeToPtr(row.CheckInTime),
				CheckInMethod: row.CheckInMethod.AttendancesCheckInMethod,
				Status:        row.Status.AttendancesStatus,
				CheckedBy:     NullInt64ToUint64Ptr(row.CheckedBy),
			},
			UserName:      row.UserName,
			UserEmail:     row.UserEmail,
			CheckedByName: nullStringToPtr(row.CheckedByName),
		})
	}
	return result, nil
}

func (r *repositoryAttendance) GetAttendanceByUserID(ctx context.Context, userID uint64) ([]AttendanceByUser, error) {
	rows, err := r.queries.GetAttendanceByUserID(ctx, userID)
	if err != nil {
		return nil, err
	}

	result := make([]AttendanceByUser, 0, len(rows))
	for _, row := range rows {
		result = append(result, AttendanceByUser{
			Attendance: Attendance{
				ID:            uint64(row.ID),
				OrderID:       row.OrderID,
				EventID:       row.EventID,
				UserID:        row.UserID,
				TicketTypeID:  row.TicketTypeID,
				CheckInTime:   nullTimeToPtr(row.CheckInTime),
				CheckInMethod: row.CheckInMethod.AttendancesCheckInMethod,
				Status:        row.Status.AttendancesStatus,
				CheckedBy:     NullInt64ToUint64Ptr(row.CheckedBy),
				CreatedAt:     row.CreatedAt.Time,
			},
			EventTitle: row.EventTitle,
			StartTime:  row.StartTime,
			EndTime:    row.EndTime,
		})
	}
	return result, nil
}

// ============================================
// STATISTICS
// ============================================

func (r *repositoryAttendance) CountAttendeesByEvent(ctx context.Context, eventID uint64) (int64, error) {
	return r.queries.CountAttendeesByEvent(ctx, eventID)
}

func (r *repositoryAttendance) CountAttendeesByOrganizer(ctx context.Context, organizerID uint64) (int64, error) {
	return r.queries.CountAttendeesByOrganizer(ctx, organizerID)
}

func (r *repositoryAttendance) CountAttendeesByTicketType(ctx context.Context, ticketTypeID uint64) (int64, error) {
	return r.queries.CountAttendeesByTicketType(ctx, ticketTypeID)
}

func (r *repositoryAttendance) GetEventAttendanceRate(ctx context.Context, organizerID uint64) ([]EventAttendanceRate, error) {
	rows, err := r.queries.GetEventAttendanceRate(ctx, organizerID)
	if err != nil {
		return nil, err
	}

	result := make([]EventAttendanceRate, 0, len(rows))
	for _, row := range rows {
		rate, err := toFloat64(row.AttendanceRate)
		if err != nil {
			return nil, err
		}
		result = append(result, EventAttendanceRate{
			EventID:        row.EventID,
			EventTitle:     row.EventTitle,
			TotalOrders:    row.TotalOrders,
			Attended:       row.Attended,
			AttendanceRate: rate,
		})
	}
	return result, nil
}

func (r *repositoryAttendance) GetCheckInTimeline(ctx context.Context, eventID uint64) ([]CheckInTimelineHour, error) {
	rows, err := r.queries.GetCheckInTimeline(ctx, eventID)
	if err != nil {
		return nil, err
	}

	result := make([]CheckInTimelineHour, 0, len(rows))
	for _, row := range rows {
		result = append(result, CheckInTimelineHour{
			Hour:  int(row.Hour),
			Count: row.Count,
		})
	}
	return result, nil
}

func (r *repositoryAttendance) GetCheckInTimelineByDate(ctx context.Context, eventID uint64) ([]CheckInTimelineDate, error) {
	rows, err := r.queries.GetCheckInTimelineByDate(ctx, eventID)
	if err != nil {
		return nil, err
	}

	result := make([]CheckInTimelineDate, 0, len(rows))
	for _, row := range rows {
		result = append(result, CheckInTimelineDate{
			Date:  row.Date,
			Count: row.Count,
		})
	}
	return result, nil
}

// ============================================
// UPDATE / DELETE
// ============================================

func (r *repositoryAttendance) CancelAttendance(ctx context.Context, orderID uint64, eventID uint64) error {
	return r.queries.CancelAttendance(ctx, database.CancelAttendanceParams{
		OrderID: orderID,
		EventID: eventID,
	})
}

func (r *repositoryAttendance) SoftDeleteAttendance(ctx context.Context, id uint64) error {
	return r.queries.SoftDeleteAttendance(ctx, int64(id))
}

// ============================================
// VALIDATION
// ============================================

func (r *repositoryAttendance) CheckAlreadyCheckedIn(ctx context.Context, orderID uint64, eventID uint64) (bool, error) {
	count, err := r.queries.CheckAlreadyCheckedIn(ctx, database.CheckAlreadyCheckedInParams{
		OrderID: orderID,
		EventID: eventID,
	})
	if err != nil {
		return false, err
	}
	return count > 0, nil
}

func (r *repositoryAttendance) GetOrderStatus(ctx context.Context, orderID uint64) (string, error) {
	status, err := r.queries.GetOrderStatus(ctx, orderID)
	if err != nil {
		return "", err
	}
	return string(status), nil
}

func (r *repositoryAttendance) GetEventTimeRange(ctx context.Context, eventID uint64) (time.Time, time.Time, error) {
	row, err := r.queries.GetEventTimeRange(ctx, eventID)
	if err != nil {
		return time.Time{}, time.Time{}, err
	}
	return row.StartTime, row.EndTime, nil
}
func (r *repositoryAttendance) GetOrderItemDetail(ctx context.Context, ticketID uint64) (*OrderItemDetail, error) {
	row, err := r.queries.GetOrderItemDetail(ctx, ticketID)
	if err != nil {
		return nil, err
	}

	return &OrderItemDetail{
		OrderID:      row.OrderID,
		EventID:      row.EventID,
		UserID:       row.UserID,
		TicketTypeID: row.TicketTypeID,
	}, nil
}
