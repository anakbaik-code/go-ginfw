package ticket

import "time"

type TicketType struct {
	ID                uint64
	EventID           uint64
	Name              string
	Price             uint64
	Quota             uint32
	MaxPerTransaction uint32
	StartSaleAt       time.Time
	EndSaleAt         time.Time
	CreatedAt         time.Time
	UpdatedAt         time.Time
	DeletedAt         time.Time
}
