package event

import (
	"go-fwgin/internal/database"
	"time"
)

type Event struct {
	ID          uint64
	CategoryID  uint32
	UserID      uint64
	Title       string
	Description string
	Location    string
	StartTime   time.Time
	EndTime     time.Time
	Status      database.EventsStatus
	CreatedAt   time.Time
	UpdatedAt   time.Time
	DeletedAt   *time.Time
}
type EventWithDetails struct {
	Event
	CategoryName   string
	OrganizerName  string
	AvailableQuota int32
}

type EventStatus struct {
	Status database.EventsStatus
}
