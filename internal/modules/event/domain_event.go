package event

import "time"

type Event struct {
	ID          uint64
	CategoryID  uint32
	UserID      uint64
	Title       string
	Description string
	Location    string
	StartTime   time.Time
	EndTime     time.Time
	Price       float64
	Quota       int32
	Status      string
	CreatedAt   time.Time
	UpdatedAt   time.Time
	DeletedAt   *time.Time
}
