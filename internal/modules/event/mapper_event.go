package event

import "go-fwgin/internal/database"

func ToEventResponse(e EventWithDetails) EventResponse {
	return EventResponse{
		ID:             e.ID,
		CategoryName:   e.CategoryName,
		OrganizerName:  e.OrganizerName,
		Title:          e.Title,
		Description:    e.Description,
		Location:       e.Location,
		StartTime:      e.StartTime,
		EndTime:        e.EndTime,
		Price:          e.Price,
		Quota:          e.Quota,
		AvailableQuota: uint32(e.AvailableQuota),
		Status:         string(e.Status),
		CreatedAt:      e.CreatedAt,
	}
}

func ToEventResponses(events []EventWithDetails) []EventResponse {
	responses := make([]EventResponse, 0, len(events))
	for _, e := range events {
		responses = append(responses, ToEventResponse(e))
	}
	return responses
}

func ToEvent(req RequestCreateEvent, userID uint64) Event {
	return Event{
		CategoryID:  req.CategoryID,
		UserID:      userID,
		Title:       req.Title,
		Description: req.Description,
		Location:    req.Location,
		StartTime:   req.StartTime,
		EndTime:     req.EndTime,
		Price:       req.Price,
		Quota:       req.Quota,
		Status:      database.EventsStatus(req.Status),
	}
}
func ToUpdateEvent(id uint64, userID uint64, req RequestUpdateEvent) Event {
	return Event{
		ID:          id,
		UserID:      userID,
		CategoryID:  req.CategoryID,
		Title:       req.Title,
		Description: req.Description,
		Location:    req.Location,
		StartTime:   req.StartTime,
		EndTime:     req.EndTime,
		Price:       req.Price,
		Quota:       req.Quota,
	}
}

func ToEventStatus(req RequestUpdateEventStatus) EventStatus {
	return EventStatus{
		Status: database.EventsStatus(req.Status),
	}
}
