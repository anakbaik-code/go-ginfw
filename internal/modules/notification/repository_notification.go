package notification

import (
	"context"
	"database/sql"
	"go-fwgin/internal/database"
)

type RepositoryNotification interface {
	CountUnread(ctx context.Context, userID uint64) (uint64, error)
	Create(ctx context.Context, n Notification) (int64, error)
	Delete(ctx context.Context, id uint64) error
	DeleteByUserID(ctx context.Context, userID uint64) error
	GetByID(ctx context.Context, id uint64) (*Notification, error)
	ListByUserID(ctx context.Context, userID uint64, limit int32, offset int32) ([]Notification, error)
	MarkAllRead(ctx context.Context, userID uint64) error
	MarkRead(ctx context.Context, id uint64) error
}
type repositoryNotification struct {
	queries *database.Queries
}

func NewRepositoryNotification(q *database.Queries) RepositoryNotification {
	return &repositoryNotification{queries: q}
}

func (r *repositoryNotification) CountUnread(ctx context.Context, userID uint64) (uint64, error) {
	result, err := r.queries.CountUnreadNotifications(ctx, userID)
	if err != nil {
		return 0, err
	}
	return uint64(result), nil
}
func (r *repositoryNotification) Create(ctx context.Context, n Notification) (int64, error) {
	refType := sql.NullString{
		String: n.ReferenceType,
		Valid:  true,
	}
	refID := sql.NullInt64{
		Int64: n.ReferenceID,
		Valid: true,
	}
	result, err := r.queries.CreateNotification(ctx, database.CreateNotificationParams{
		UserID:        n.ID,
		Title:         n.Title,
		Message:       n.Message,
		Type:          n.Type,
		IsRead:        n.IsRead,
		ReferenceType: refType,
		ReferenceID:   refID,
	})
	if err != nil {
		return 0, err
	}
	lastID, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}
	return lastID, nil
}
func (r *repositoryNotification) Delete(ctx context.Context, id uint64) error {
	err := r.queries.DeleteNotification(ctx, id)
	if err != nil {
		return err
	}
	return nil
}
func (r *repositoryNotification) DeleteByUserID(ctx context.Context, userID uint64) error {
	err := r.queries.DeleteNotificationsByUserID(ctx, userID)
	if err != nil {
		return err
	}
	return nil
}
func (r *repositoryNotification) GetByID(ctx context.Context, id uint64) (*Notification, error) {
	result, err := r.queries.GetNotificationByID(ctx, id)
	if err != nil {
		return nil, err
	}
	notification := &Notification{
		ID:            result.ID,
		UserID:        result.UserID,
		Title:         result.Title,
		Message:       result.Message,
		Type:          result.Type,
		IsRead:        result.IsRead,
		ReferenceType: result.ReferenceType.String,
		ReferenceID:   result.ReferenceID.Int64,
		CreatedAt:     result.CreatedAt.Time,
		UpdatedAt:     result.UpdatedAt.Time,
		DeletedAt:     result.DeletedAt.Time,
	}
	return notification, nil
}

func (r *repositoryNotification) ListByUserID(ctx context.Context, userID uint64, limit int32, offset int32) ([]Notification, error) {
	rows, err := r.queries.ListNotificationsByUserID(ctx, database.ListNotificationsByUserIDParams{
		UserID: userID,
		Limit:  limit,
		Offset: offset,
	})
	if err != nil {
		return nil, err
	}
	notifications := make([]Notification, 0, len(rows))
	for _, row := range rows {
		notifications = append(notifications, Notification{
			ID:            row.ID,
			UserID:        row.UserID,
			Title:         row.Title,
			Message:       row.Message,
			Type:          row.Type,
			IsRead:        row.IsRead,
			ReferenceType: row.ReferenceType.String,
			ReferenceID:   row.ReferenceID.Int64,
			CreatedAt:     row.CreatedAt.Time,
			UpdatedAt:     row.UpdatedAt.Time,
			DeletedAt:     row.DeletedAt.Time,
		})
	}
	return notifications, nil
}
func (r *repositoryNotification) MarkAllRead(ctx context.Context, userID uint64) error {
	err := r.queries.MarkAllNotificationsAsRead(ctx, userID)
	if err != nil {
		return err
	}
	return nil
}
func (r *repositoryNotification) MarkRead(ctx context.Context, id uint64) error {
	err := r.queries.MarkNotificationAsRead(ctx, id)
	if err != nil {
		return err
	}
	return nil
}
