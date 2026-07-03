package notification

import "time"

type Notification struct {
	ID            uint64
	UserID        uint64
	Title         string
	Message       string
	Type          string
	IsRead        bool
	ReferenceType string
	ReferenceID   int64
	CreatedAt     time.Time
	UpdatedAt     time.Time
	DeletedAt     time.Time
}
