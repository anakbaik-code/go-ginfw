package event

import "time"

type RequestCreateEvent struct {
	CategoryID  uint32    `json:"category_id" binding:"required"`
	Title       string    `json:"title" binding:"required"`
	Description string    `json:"description" binding:"required"`
	Location    string    `json:"location" binding:"required"`
	StartTime   time.Time `json:"start_time" binding:"required"`
	EndTime     time.Time `json:"end_time" binding:"required"`
	Price       uint32    `json:"price"`
	Quota       uint32    `json:"quota" binding:"required"`
	Status      string    `json:"status"`
}
type RequestUpdateEvent struct {
	CategoryID  uint32    `json:"category_id" binding:"required"`
	Title       string    `json:"title" binding:"required"`
	Description string    `json:"description" binding:"required"`
	Location    string    `json:"location" binding:"required"`
	StartTime   time.Time `json:"start_time" binding:"required"`
	EndTime     time.Time `json:"end_time" binding:"required"`
	Price       uint32    `json:"price"`
	Quota       uint32    `json:"quota" binding:"required"`
}
type EventResponse struct {
	ID             uint64    `json:"id"`
	CategoryName   string    `json:"category_name"`
	OrganizerName  string    `json:"organizer_name"`
	Title          string    `json:"title"`
	Description    string    `json:"description"`
	Location       string    `json:"location"`
	StartTime      time.Time `json:"start_time"`
	EndTime        time.Time `json:"end_time"`
	Price          uint32    `json:"price"`
	Quota          uint32    `json:"quota"`
	AvailableQuota uint32    `json:"available_quota"`
	Status         string    `json:"status"`
	CreatedAt      time.Time `json:"created_at"`
}
type RequestUpdateEventStatus struct {
	Status string `json:"status" binding:"required,oneof=active inactive cancelled"`
}
type RequestListEvent struct {
	Page  int32 `form:"page" binding:"gte=1"`
	Limit int32 `form:"limit" binding:"gte=1"`
}
