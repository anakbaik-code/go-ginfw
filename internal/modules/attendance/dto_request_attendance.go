package attendance

// Request
// ============================================
// REQUEST DTO
// ============================================

type CheckInRequest struct {
	OrderID      uint64 `json:"order_id" binding:"required"`
	EventID      uint64 `json:"event_id" binding:"required"`
	UserID       uint64 `json:"user_id" binding:"required"`
	TicketTypeID uint64 `json:"ticket_type_id" binding:"required"`
	CheckedBy    uint64 `json:"checked_by" binding:"required"`
}

type CheckInByQRRequest struct {
	QRToken   string `json:"qr_token" binding:"required"`
	CheckedBy uint64 `json:"checked_by" binding:"required"`
}

type ManualCheckInRequest struct {
	OrderID      uint64 `json:"order_id" binding:"required"`
	EventID      uint64 `json:"event_id" binding:"required"`
	UserID       uint64 `json:"user_id" binding:"required"`
	TicketTypeID uint64 `json:"ticket_type_id" binding:"required"`
	CheckedBy    uint64 `json:"checked_by" binding:"required"`
}

type CancelAttendanceRequest struct {
	OrderID uint64 `json:"order_id" binding:"required"`
	EventID uint64 `json:"event_id" binding:"required"`
}
