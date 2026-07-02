package eventmedia

import (
	"context"
	"database/sql"
	"go-fwgin/internal/database"
)

type RepositoryEventMedia interface {
	Create(ctx context.Context, e EventMedia) (uint64, error)
	Delete(ctx context.Context, id uint64) error
	GetByID(ctx context.Context, id uint64) (*EventMedia, error)
	GetPrimaryMedia(ctx context.Context, eventID uint64) (*EventMedia, error)
	List(ctx context.Context, eventID uint64) ([]EventMedia, error)
	ResetEventPrimary(ctx context.Context, eventID uint64) error
	SetPrimaryMedia(ctx context.Context, id uint64, eventID uint64) error
}
type repositoryEventMedia struct {
	queries *database.Queries
}

func NewRepositoryEventMedia(q *database.Queries) RepositoryEventMedia {
	return &repositoryEventMedia{queries: q}
}

func (r *repositoryEventMedia) Create(ctx context.Context, e EventMedia) (uint64, error) {
	primary := sql.NullBool{
		Bool:  e.IsPrimary,
		Valid: true,
	}
	result, err := r.queries.CreateEventMedia(ctx, database.CreateEventMediaParams{
		EventID:   e.EventId,
		ImagePath: e.ImagePath,
		IsPrimary: primary,
	})
	if err != nil {
		return 0, err
	}
	lastId, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}
	return uint64(lastId), nil
}

func (r *repositoryEventMedia) Delete(ctx context.Context, id uint64) error {
	if err := r.queries.DeleteEventMedia(ctx, id); err != nil {
		return err
	}
	return nil
}

func (r *repositoryEventMedia) GetByID(ctx context.Context, id uint64) (*EventMedia, error) {
	result, err := r.queries.GetEventMediaByID(ctx, id)
	if err != nil {
		return nil, err
	}
	eventMedia := &EventMedia{
		ID:        result.ID,
		EventId:   result.EventID,
		ImagePath: result.ImagePath,
		IsPrimary: result.IsPrimary.Bool,
		CreatedAt: result.CreatedAt.Time,
	}
	return eventMedia, nil
}
func (r *repositoryEventMedia) GetPrimaryMedia(ctx context.Context, eventID uint64) (*EventMedia, error) {
	result, err := r.queries.GetEventPrimaryMedia(ctx, eventID)
	if err != nil {
		return nil, err
	}
	eventMedia := &EventMedia{
		ID:        result.ID,
		EventId:   result.EventID,
		ImagePath: result.ImagePath,
		IsPrimary: result.IsPrimary.Bool,
		CreatedAt: result.CreatedAt.Time,
	}
	return eventMedia, nil
}

func (r *repositoryEventMedia) List(ctx context.Context, eventID uint64) ([]EventMedia, error) {
	rows, err := r.queries.ListEventMedia(ctx, eventID)
	if err != nil {
		return nil, err
	}
	eventsMedia := make([]EventMedia, 0, len(rows))
	for _, row := range rows {
		eventsMedia = append(eventsMedia, EventMedia{
			ID:        row.ID,
			EventId:   row.EventID,
			ImagePath: row.ImagePath,
			IsPrimary: row.IsPrimary.Bool,
			CreatedAt: row.CreatedAt.Time,
		})
	}
	return eventsMedia, nil
}

func (r *repositoryEventMedia) ResetEventPrimary(ctx context.Context, eventID uint64) error {
	if err := r.queries.ResetEventPrimaryMedia(ctx, eventID); err != nil {
		return err
	}
	return nil
}

func (r *repositoryEventMedia) SetPrimaryMedia(ctx context.Context, id uint64, eventID uint64) error {
	if err := r.queries.SetEventPrimaryMedia(ctx, database.SetEventPrimaryMediaParams{
		ID:      id,
		EventID: eventID,
	}); err != nil {
		return err
	}
	return nil
}
