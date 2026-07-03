package ticket

import (
	"context"
	"database/sql"
	"go-fwgin/internal/database"
	"time"
)

type RepositoryTicketType interface {
	CountByEventID(ctx context.Context, eventID uint64) (int64, error)
	CreateTicketType(ctx context.Context, tt TicketType) (int64, error)
	DeleteType(ctx context.Context, id uint64) error
	GetAvailableByID(ctx context.Context, id uint64, start time.Time, end time.Time) (*TicketType, error)
	GetTicketTypeByID(ctx context.Context, id uint64) (*TicketType, error)
	ListActive(ctx context.Context, evenID uint64, start time.Time, end time.Time) ([]TicketType, error)
	ListByEventID(ctx context.Context, eventID uint64, limit int32, offset int32) ([]TicketType, error)
	UpdateTicketType(ctx context.Context, tt TicketType) error
}
type repositoryTicketType struct {
	queries *database.Queries
}

func NewRepositoryTicketType(q *database.Queries) RepositoryTicketType {
	return &repositoryTicketType{queries: q}
}
func (r *repositoryTicketType) CountByEventID(ctx context.Context, eventID uint64) (int64, error) {
	result, err := r.queries.CountTicketTypesByEventID(ctx, eventID)
	if err != nil {
		return 0, err
	}
	return result, nil
}
func (r *repositoryTicketType) CreateTicketType(ctx context.Context, tt TicketType) (int64, error) {
	transaction := sql.NullInt32{
		Int32: int32(tt.MaxPerTransaction),
		Valid: true,
	}
	startSale := sql.NullTime{
		Time:  tt.StartSaleAt,
		Valid: true,
	}
	endSale := sql.NullTime{
		Time:  tt.EndSaleAt,
		Valid: true,
	}
	result, err := r.queries.CreateTicketType(ctx, database.CreateTicketTypeParams{
		EventID:           tt.EventID,
		Name:              tt.Name,
		Price:             tt.Price,
		Quota:             tt.Quota,
		MaxPerTransaction: transaction,
		StartSaleAt:       startSale,
		EndSaleAt:         endSale,
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
func (r *repositoryTicketType) DeleteType(ctx context.Context, id uint64) error {
	err := r.queries.DeleteTicketType(ctx, id)
	if err != nil {
		return err
	}
	return nil
}
func (r *repositoryTicketType) GetAvailableByID(ctx context.Context, id uint64, start time.Time, end time.Time) (*TicketType, error) {
	startSale := sql.NullTime{
		Time:  start,
		Valid: true,
	}
	endSale := sql.NullTime{
		Time:  end,
		Valid: true,
	}
	result, err := r.queries.GetAvailableTicketTypeByID(ctx, database.GetAvailableTicketTypeByIDParams{
		ID:          id,
		StartSaleAt: startSale,
		EndSaleAt:   endSale,
	})
	if err != nil {
		return nil, err
	}
	ticketType := ToTicketTypeFromAvailableRow(result)
	return ticketType, nil
}
func (r *repositoryTicketType) GetTicketTypeByID(ctx context.Context, id uint64) (*TicketType, error) {
	result, err := r.queries.GetTicketTypeByID(ctx, id)
	if err != nil {
		return nil, err
	}
	ticket := ToTicketTypeFromDetailRow(result)
	return ticket, nil
}
func (r *repositoryTicketType) ListActive(ctx context.Context, evenID uint64, start time.Time, end time.Time) ([]TicketType, error) {
	startSale := sql.NullTime{
		Time:  start,
		Valid: true,
	}
	endSale := sql.NullTime{
		Time:  end,
		Valid: true,
	}
	result, err := r.queries.ListActiveTicketTypes(ctx, database.ListActiveTicketTypesParams{
		EventID:     evenID,
		StartSaleAt: startSale,
		EndSaleAt:   endSale,
	})
	if err != nil {
		return nil, err
	}
	ticketTypes := ToTicketTypes(result)
	return ticketTypes, nil
}

func (r *repositoryTicketType) ListByEventID(ctx context.Context, eventID uint64, limit int32, offset int32) ([]TicketType, error) {
	result, err := r.queries.ListTicketTypesByEventID(ctx, database.ListTicketTypesByEventIDParams{
		EventID: eventID,
		Limit:   limit,
		Offset:  offset,
	})
	if err != nil {
		return nil, err
	}
	ticketTypes := ToTicketTypeFromEventList(result)
	return ticketTypes, nil
}
func (r *repositoryTicketType) UpdateTicketType(ctx context.Context, tt TicketType) error {
	transaction := sql.NullInt32{
		Int32: int32(tt.MaxPerTransaction),
		Valid: true,
	}
	startSale := sql.NullTime{
		Time:  tt.StartSaleAt,
		Valid: true,
	}
	endSale := sql.NullTime{
		Time:  tt.EndSaleAt,
		Valid: true,
	}
	err := r.queries.UpdateTicketType(ctx, database.UpdateTicketTypeParams{
		ID:                tt.ID,
		Name:              tt.Name,
		Price:             tt.Price,
		Quota:             tt.Quota,
		MaxPerTransaction: transaction,
		StartSaleAt:       startSale,
		EndSaleAt:         endSale,
	})
	if err != nil {
		return err
	}
	return nil
}
