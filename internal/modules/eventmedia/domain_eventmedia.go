package eventmedia

import "time"

type EventMedia struct {
	ID        uint64
	EventId   uint64
	ImagePath string
	IsPrimary bool
	CreatedAt time.Time
}
