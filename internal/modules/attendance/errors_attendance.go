package attendance

import "errors"

var (
	ErrInvalidQRToken     = errors.New("QR code is invalid or expired")
	ErrOrderNotPaid       = errors.New("order is not paid, cannot check in")
	ErrAlreadyCheckedIn   = errors.New("this order has already checked in for this event")
	ErrEventNotStarted    = errors.New("event has not started yet")
	ErrEventEnded         = errors.New("event has already ended")
	ErrInvalidCheckInTime = errors.New("check-in is only allowed during the event window")
)
