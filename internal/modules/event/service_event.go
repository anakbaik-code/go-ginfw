package event

import (
	"context"
	"errors"
	"go-fwgin/internal/database"
)

type ServiceEvent interface {
	// Organizer
	CreateEvent(ctx context.Context, userId uint64, req RequestCreateEvent) (uint64, error)
	DeleteEvent(ctx context.Context, id uint64) error
	UpdateEvent(ctx context.Context, id uint64, userID uint64, req RequestUpdateEvent) error
	UpdateStatus(ctx context.Context, req RequestUpdateEventStatus, id uint64) error

	// Detail
	GetByID(ctx context.Context, id uint64) (*EventResponse, error)
	GetMyEventById(ctx context.Context, userId uint64, id uint64) (*EventResponse, error)

	// List
	ListActive(ctx context.Context, page int32, limit int32) ([]EventResponse, int64, error)
	ListInActive(ctx context.Context, page int32, limit int32) ([]EventResponse, int64, error)
	ListCancelled(ctx context.Context, page int32, limit int32) ([]EventResponse, int64, error)

	// Organizer Dashboard
	ListMyEvents(ctx context.Context, userId uint64, page int32, limit int32) ([]EventResponse, int64, error)
}
type serviceEvent struct {
	repo RepositoryEvent
}

func NewServiceEvent(repo RepositoryEvent) ServiceEvent {
	return &serviceEvent{
		repo: repo,
	}
}

func (s *serviceEvent) CreateEvent(ctx context.Context, userId uint64, req RequestCreateEvent) (uint64, error) {
	event := ToEvent(req, userId)
	if event.Status == "" {
		event.Status = database.EventsStatusActive
	}

	id, err := s.repo.Create(ctx, event)
	if err != nil {
		return 0, err
	}
	return id, nil
}

func (s *serviceEvent) DeleteEvent(ctx context.Context, id uint64) error {
	if err := s.repo.Delete(ctx, id); err != nil {
		return err
	}
	return nil
}

func (s *serviceEvent) GetByID(ctx context.Context, id uint64) (*EventResponse, error) {
	result, err := s.repo.GetById(ctx, id)
	if err != nil {
		return nil, err
	}

	if result.Status != database.EventsStatusActive {
		return nil, errors.New("Event Not Found")
	}
	event := ToEventResponse(*result)

	return &event, nil
}


func (s *serviceEvent) ListActive(ctx context.Context, page int32, limit int32) ([]EventResponse, int64, error) {
	if page <= 0 {
		page = 1
	}

	if limit <= 0 {
		limit = 10
	}

	offset := (page - 1) * limit
	events, err := s.repo.ListByStatus(ctx, EventStatus{
		Status: database.EventsStatusActive,
	}, limit, offset)
	if err != nil {
		return nil, 0, err
	}

	total, err := s.repo.CountByStatus(ctx, EventStatus{
		Status: database.EventsStatusActive,
	})
	if err != nil {
		return nil, 0, err
	}

	eventsResponse := ToEventResponses(events)
	return eventsResponse, total, nil
}

func (s *serviceEvent) ListInActive(ctx context.Context, page int32, limit int32) ([]EventResponse, int64, error) {
	if page <= 0 {
		page = 1
	}

	if limit <= 0 {
		limit = 10
	}

	offset := (page - 1) * limit
	events, err := s.repo.ListByStatus(ctx, EventStatus{
		Status: database.EventsStatusInactive,
	}, limit, offset)
	if err != nil {
		return nil, 0, err
	}

	total, err := s.repo.CountByStatus(ctx, EventStatus{
		Status: database.EventsStatusInactive,
	})
	if err != nil {
		return nil, 0, err
	}

	eventsResponse := ToEventResponses(events)
	return eventsResponse, total, nil
}

func (s *serviceEvent) ListCancelled(ctx context.Context, page int32, limit int32) ([]EventResponse, int64, error) {
	if page <= 0 {
		page = 1
	}

	if limit <= 0 {
		limit = 10
	}

	offset := (page - 1) * limit
	events, err := s.repo.ListByStatus(ctx, EventStatus{
		Status: database.EventsStatusCancelled,
	}, limit, offset)
	if err != nil {
		return nil, 0, err
	}

	total, err := s.repo.CountByStatus(ctx, EventStatus{
		Status: database.EventsStatusCancelled,
	})
	if err != nil {
		return nil, 0, err
	}

	eventsResponse := ToEventResponses(events)
	return eventsResponse, total, nil
}

func (s *serviceEvent) UpdateEvent(ctx context.Context, id uint64, userID uint64, req RequestUpdateEvent) error {
	event, err := s.repo.GetById(ctx, id)
	if err != nil {
		return err
	}
	if event.UserID != userID {
		return errors.New("Error Forbidden")
	}
	result := ToUpdateEvent(id, userID, req)
	if err := s.repo.Update(ctx, result); err != nil {
		return err
	}
	return nil
}

func (s *serviceEvent) UpdateStatus(ctx context.Context, req RequestUpdateEventStatus, id uint64) error {
	status := ToEventStatus(req)
	err := s.repo.UpdateStatus(ctx, status, id)
	if err != nil {
		return err
	}
	return nil
}

func (s *serviceEvent) ListMyEvents(ctx context.Context, userId uint64, page int32, limit int32) ([]EventResponse, int64, error) {
	if page <= 0 {
		page = 1
	}

	if limit <= 0 {
		limit = 10
	}

	offset := (page - 1) * limit
	events, err := s.repo.ListMyEvents(ctx, userId, limit, offset)
	if err != nil {
		return nil, 0, err
	}

	total, err := s.repo.CountMyEvent(ctx, userId)
	if err != nil {
		return nil, 0, err
	}

	eventsResponse := ToEventResponses(events)
	return eventsResponse, total, nil
}

func (s *serviceEvent) GetMyEventById(ctx context.Context, userId uint64, id uint64) (*EventResponse, error) {
	result, err := s.repo.GetMyEventId(ctx, userId, id)
	if err != nil {
		return nil, err
	}
	if result.UserID != userId {
		return nil,errors.New("Error Forbidden")
	}
	event := ToEventResponse(*result)
	return &event, nil
}
