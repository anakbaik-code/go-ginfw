package attendance

import (
	"context"
	"time"
)

// AttendanceValidator holds pure business-rule validation logic,
// separated from gin binding validation (struct tags handle basic shape validation).
type AttendanceValidator struct {
	repo RepositoryAttendance
}

func NewAttendanceValidator(repo RepositoryAttendance) *AttendanceValidator {
	return &AttendanceValidator{repo: repo}
}


// ValidateCheckIn ensures the order is eligible to check in:
// 1. Order must be paid
// 2. Must not already be checked in
// 3. Must be within the event's time window
func (v *AttendanceValidator) ValidateCheckIn(ctx context.Context, orderID, eventID uint64) error {
	status, err := v.repo.GetOrderStatus(ctx, orderID)
	if err != nil {
		return err
	}
	if status != "paid" {
		return ErrOrderNotPaid
	}

	alreadyCheckedIn, err := v.repo.CheckAlreadyCheckedIn(ctx, orderID, eventID)
	if err != nil {
		return err
	}
	if alreadyCheckedIn {
		return ErrAlreadyCheckedIn
	}

	startTime, endTime, err := v.repo.GetEventTimeRange(ctx, eventID)
	if err != nil {
		return err
	}
	now := time.Now()
	if now.Before(startTime) {
		return ErrEventNotStarted
	}
	if now.After(endTime) {
		return ErrEventEnded
	}

	return nil
}
