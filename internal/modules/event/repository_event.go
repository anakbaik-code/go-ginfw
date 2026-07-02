package event

import (
	"context"
	"go-fwgin/internal/database"
	"time"
)

type RepositoryEvent interface {
	Create(ctx context.Context, e Event) (uint64, error)
	Delete(ctx context.Context, id uint64) error
	GetById(ctx context.Context, id uint64) (*EventWithDetails, error)
	ListByStatus(ctx context.Context, es EventStatus, limit int32, offset int32) ([]EventWithDetails, error)
	Update(ctx context.Context, e Event) error
	UpdateStatus(ctx context.Context, e EventStatus, id uint64) error
	CountByStatus(ctx context.Context, es EventStatus) (int64, error)
	ListMyEvents(ctx context.Context, userID uint64, limit int32, offset int32) ([]EventWithDetails, error)
	CountMyEvent(ctx context.Context, userId uint64) (int64, error)
	GetMyEventId(ctx context.Context, userId uint64, id uint64) (*EventWithDetails, error)
}

type repositoryEvent struct {
	queries *database.Queries
}

func NewRepositoryEvent(query *database.Queries) RepositoryEvent {
	return &repositoryEvent{queries: query}
}

func (r *repositoryEvent) Create(ctx context.Context, e Event) (uint64, error) {
	res, err := r.queries.CreateEvent(ctx, database.CreateEventParams{
		CategoryID:  e.CategoryID,
		UserID:      e.UserID,
		Title:       e.Title,
		Description: e.Description,
		Location:    e.Location,
		StartTime:   e.StartTime,
		EndTime:     e.EndTime,
		Status: database.NullEventsStatus{
			EventsStatus: e.Status,
			Valid:        e.Status != "",
		},
	})
	if err != nil {
		return 0, err
	}
	id, err := res.LastInsertId()
	if err != nil {
		return 0, err
	}
	return uint64(id), nil
}

func (r *repositoryEvent) Delete(ctx context.Context, id uint64) error {
	if err := r.queries.DeleteEvent(ctx, id); err != nil {
		return err
	}
	return nil
}

func (r *repositoryEvent) GetById(ctx context.Context, id uint64) (*EventWithDetails, error) {
	row, err := r.queries.GetEventByID(ctx, id)
	if err != nil {
		return nil, err
	}
	var statusDomain database.EventsStatus
	if row.Status.Valid {
		statusDomain = row.Status.EventsStatus
	}

	var createdAtDomain time.Time
	if row.CreatedAt.Valid {
		createdAtDomain = row.CreatedAt.Time
	}

	event := &EventWithDetails{
		Event: Event{
			ID:          row.ID,
			Title:       row.Title,
			Description: row.Description,
			Location:    row.Location,
			StartTime:   row.StartTime,
			EndTime:     row.EndTime,
			Status:      statusDomain,
			CreatedAt:   createdAtDomain,
		},
		CategoryName:   row.CategoryName,
		OrganizerName:  row.OrganizerName,
		AvailableQuota: row.AvailableQuota,
	}
	return event, nil
}

func (r *repositoryEvent) ListByStatus(ctx context.Context, es EventStatus, limit int32, offset int32) ([]EventWithDetails, error) {
	rows, err := r.queries.ListEventsByStatus(ctx, database.ListEventsByStatusParams{
		Status: database.NullEventsStatus{
			EventsStatus: database.EventsStatus(es.Status),
		},
		Limit:  limit,
		Offset: offset,
	})
	if err != nil {
		return nil, err
	}

	events := make([]EventWithDetails, 0, len(rows))
	for _, row := range rows {
		var statusDomain database.EventsStatus
		if row.Status.Valid {
			statusDomain = row.Status.EventsStatus
		}

		var createdAtDomain time.Time
		if row.CreatedAt.Valid {
			createdAtDomain = row.CreatedAt.Time
		}
		events = append(events, EventWithDetails{
			Event: Event{
				ID:          row.ID,
				Title:       row.Title,
				Description: row.Description,
				Location:    row.Location,
				StartTime:   row.StartTime,
				EndTime:     row.EndTime,
				Status:      statusDomain,
				CreatedAt:   createdAtDomain,
			},
			CategoryName:   row.CategoryName,
			OrganizerName:  row.OrganizerName,
			AvailableQuota: row.AvailableQuota,
		})
	}
	return events, nil
}

func (r *repositoryEvent) Update(ctx context.Context, e Event) error {
	err := r.queries.UpdateEvent(ctx, database.UpdateEventParams{
		ID:          e.ID,
		CategoryID:  e.CategoryID,
		Title:       e.Title,
		Description: e.Description,
		Location:    e.Location,
		StartTime:   e.StartTime,
		EndTime:     e.EndTime,
		Status: database.NullEventsStatus{
			EventsStatus: e.Status,
			Valid:        e.Status != "",
		},
	})
	if err != nil {
		return err
	}
	return nil
}

func (r *repositoryEvent) UpdateStatus(ctx context.Context, e EventStatus, id uint64) error {
	err := r.queries.UpdateEventStatus(ctx, database.UpdateEventStatusParams{
		Status: database.NullEventsStatus{
			EventsStatus: e.Status,
			Valid:        e.Status != "",
		},
		ID: id,
	})
	if err != nil {
		return err
	}
	return nil
}
func (r *repositoryEvent) CountByStatus(ctx context.Context, es EventStatus) (int64, error) {
	res, err := r.queries.CountEventsByStatus(ctx, database.NullEventsStatus{
		EventsStatus: es.Status,
		Valid:        true,
	})
	if err != nil {
		return 0, err
	}
	return res, nil
}

func (r *repositoryEvent) ListMyEvents(ctx context.Context, userID uint64, limit int32, offset int32) ([]EventWithDetails, error) {
	rows, err := r.queries.ListMyEvents(ctx, database.ListMyEventsParams{
		UserID: userID,
		Limit:  limit,
		Offset: offset,
	})
	if err != nil {
		return nil, err
	}

	events := make([]EventWithDetails, 0, len(rows))
	for _, row := range rows {
		var statusDomain database.EventsStatus
		if row.Status.Valid {
			statusDomain = row.Status.EventsStatus
		}

		var createdAtDomain time.Time
		if row.CreatedAt.Valid {
			createdAtDomain = row.CreatedAt.Time
		}
		events = append(events, EventWithDetails{
			Event: Event{
				ID:          row.ID,
				Title:       row.Title,
				Description: row.Description,
				Location:    row.Location,
				StartTime:   row.StartTime,
				EndTime:     row.EndTime,
				Status:      statusDomain,
				CreatedAt:   createdAtDomain,
			},
			CategoryName:   row.CategoryName,
			OrganizerName:  row.OrganizerName,
			AvailableQuota: row.AvailableQuota,
		})
	}

	return events, nil
}

func (r *repositoryEvent) CountMyEvent(ctx context.Context, userId uint64) (int64, error) {
	result, err := r.queries.CountMyEvents(ctx, userId)
	if err != nil {
		return 0, err
	}
	return result, nil
}
func (r *repositoryEvent) GetMyEventId(ctx context.Context, userId uint64, id uint64) (*EventWithDetails, error) {
	row, err := r.queries.GetMyEventByID(ctx, database.GetMyEventByIDParams{
		ID:     id,
		UserID: userId,
	})
	if err != nil {
		return nil, err
	}

	var statusDomain database.EventsStatus
	if row.Status.Valid {
		statusDomain = row.Status.EventsStatus
	}

	var createdAtDomain time.Time
	if row.CreatedAt.Valid {
		createdAtDomain = row.CreatedAt.Time
	}

	event := &EventWithDetails{
		Event: Event{
			ID:          row.ID,
			Title:       row.Title,
			Description: row.Description,
			Location:    row.Location,
			StartTime:   row.StartTime,
			EndTime:     row.EndTime,
			Status:      statusDomain,
			CreatedAt:   createdAtDomain,
		},
		CategoryName:   row.CategoryName,
		OrganizerName:  row.OrganizerName,
		AvailableQuota: row.AvailableQuota,
	}
	return event, nil
}
