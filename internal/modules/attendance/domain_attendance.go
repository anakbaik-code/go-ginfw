package attendance

import (
	"go-fwgin/internal/database"
	"time"
)

// DOMAIN ENTITY
type Attendance struct {
	ID            uint64
	OrderID       uint64
	EventID       uint64
	UserID        uint64
	TicketTypeID  uint64
	CheckInTime   *time.Time
	CheckInMethod database.AttendancesCheckInMethod
	Status        database.AttendancesStatus
	CheckedBy     *uint64
	CreatedAt     time.Time
	UpdatedAt     time.Time
	DeletedAt     *time.Time
}

// DOMAIN VALUE OBJECTS
type AttendanceDetail struct {
	Attendance
	UserName       string
	UserEmail      string
	EventTitle     string
	TicketTypeName string
	CheckedByName  *string
}

type AttendanceByEvent struct {
	Attendance
	UserName      string
	UserEmail     string
	CheckedByName *string
}

type AttendanceByUser struct {
	Attendance
	EventTitle string
	StartTime  time.Time
	EndTime    time.Time
}

type EventAttendanceRate struct {
	EventID        uint64
	EventTitle     string
	TotalOrders    int64
	Attended       int64
	AttendanceRate float64
}

type CheckInTimelineHour struct {
	Hour  int
	Count int64
}

type CheckInTimelineDate struct {
	Date  time.Time
	Count int64
}
type OrderItemDetail struct {
	OrderID      uint64
	EventID      uint64
	UserID       uint64
	TicketTypeID uint64
}