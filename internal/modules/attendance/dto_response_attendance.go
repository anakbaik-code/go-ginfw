package attendance

import "time"

type AttendanceDetailResponse struct {
	ID             uint64     `json:"id"`
	OrderID        uint64     `json:"order_id"`
	EventID        uint64     `json:"event_id"`
	UserID         uint64     `json:"user_id"`
	TicketTypeID   uint64     `json:"ticket_type_id"`
	CheckInTime    *time.Time `json:"check_in_time,omitempty"`
	CheckInMethod  string     `json:"check_in_method"`
	Status         string     `json:"status"`
	CheckedBy      *uint64    `json:"checked_by,omitempty"`
	CreatedAt      time.Time  `json:"created_at"`
	UpdatedAt      time.Time  `json:"updated_at"`
	UserName       string     `json:"user_name"`
	UserEmail      string     `json:"user_email"`
	EventTitle     string     `json:"event_title"`
	TicketTypeName string     `json:"ticket_type_name"`
	CheckedByName  *string    `json:"checked_by_name,omitempty"`
}

type AttendanceByEventResponse struct {
	ID            uint64     `json:"id"`
	OrderID       uint64     `json:"order_id"`
	UserID        uint64     `json:"user_id"`
	UserName      string     `json:"user_name"`
	UserEmail     string     `json:"user_email"`
	CheckInTime   *time.Time `json:"check_in_time,omitempty"`
	CheckInMethod string     `json:"check_in_method"`
	Status        string     `json:"status"`
	CheckedBy     *uint64    `json:"checked_by,omitempty"`
	CheckedByName *string    `json:"checked_by_name,omitempty"`
}

type AttendanceByUserResponse struct {
	ID            uint64     `json:"id"`
	OrderID       uint64     `json:"order_id"`
	EventID       uint64     `json:"event_id"`
	UserID        uint64     `json:"user_id"`
	TicketTypeID  uint64     `json:"ticket_type_id"`
	CheckInTime   *time.Time `json:"check_in_time,omitempty"`
	CheckInMethod string     `json:"check_in_method"`
	Status        string     `json:"status"`
	CheckedBy     *uint64    `json:"checked_by,omitempty"`
	CreatedAt     time.Time  `json:"created_at"`
	EventTitle    string     `json:"event_title"`
	StartTime     time.Time  `json:"start_time"`
	EndTime       time.Time  `json:"end_time"`
}

type EventAttendanceRateResponse struct {
	EventID        uint64  `json:"event_id"`
	EventTitle     string  `json:"event_title"`
	TotalOrders    int64   `json:"total_orders"`
	Attended       int64   `json:"attended"`
	AttendanceRate float64 `json:"attendance_rate"`
}

type CheckInTimelineHourResponse struct {
	Hour  int   `json:"hour"`
	Count int64 `json:"count"`
}

type CheckInTimelineDateResponse struct {
	Date  time.Time `json:"date"`
	Count int64     `json:"count"`
}