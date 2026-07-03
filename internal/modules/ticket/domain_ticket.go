package ticket

import (
	"go-fwgin/internal/database"
	"time"
)

type Ticket struct {
	ID          uint64
	OrderItemID uint64
	TicketCode  string
	QrCode      string
	Status      database.TicketsStatus
	CheckInAt   time.Time
}
