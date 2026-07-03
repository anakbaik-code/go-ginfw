package ticket

import (
	"context"
	"database/sql"
	"go-fwgin/internal/database"
)

type RepositoryTicket interface {
	CheckIn(ctx context.Context, id uint64) error
	CountByOrderItemID(ctx context.Context, orderItemID uint64) (int64, error)
	Create(ctx context.Context, t Ticket) (int64, error)
	Delete(ctx context.Context, id uint64) error
	GetByCode(ctx context.Context, ticketCode string) (*Ticket, error)
	GetByID(ctx context.Context, id uint64) (*Ticket, error)
	ListByOrderItem(ctx context.Context, orderItemID uint64) ([]Ticket, error)
	UpdateQrCode(ctx context.Context, id uint64, qrCode string) error
	UpdateStatus(ctx context.Context, id uint64, status database.TicketsStatus) error
}
type repositoryTicket struct {
	queries *database.Queries
}

func NewRepositoryTicket(q *database.Queries) RepositoryTicket {
	return &repositoryTicket{queries: q}
}
func (r *repositoryTicket) CheckIn(ctx context.Context, id uint64) error {
	err := r.queries.CheckInTicket(ctx, id)
	if err != nil {
		return err
	}
	return nil
}
func (r *repositoryTicket) CountByOrderItemID(ctx context.Context, orderItemID uint64) (int64, error) {
	result, err := r.queries.CountTicketsByOrderItemID(ctx, orderItemID)
	if err != nil {
		return 0, err
	}
	return result, nil
}
func (r *repositoryTicket) Create(ctx context.Context, t Ticket) (int64, error) {
	checkInAt := sql.NullTime{
		Time:  t.CheckInAt,
		Valid: true,
	}
	result, err := r.queries.CreateTicket(ctx, database.CreateTicketParams{
		OrderItemID: t.OrderItemID,
		TicketCode:  t.TicketCode,
		QrCode:      t.QrCode,
		Status:      t.Status,
		CheckedInAt: checkInAt,
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
func (r *repositoryTicket) Delete(ctx context.Context, id uint64) error {
	err := r.queries.DeleteTicket(ctx, id)
	if err != nil {
		return err
	}
	return nil
}
func (r *repositoryTicket) GetByCode(ctx context.Context, ticketCode string) (*Ticket, error) {
	result, err := r.queries.GetTicketByCode(ctx, ticketCode)
	if err != nil {
		return nil, err
	}
	ticket := ToTicket(result)
	return ticket, nil
}
func (r *repositoryTicket) GetByID(ctx context.Context, id uint64) (*Ticket, error) {
	result, err := r.queries.GetTicketByID(ctx, id)
	if err != nil {
		return nil, err
	}
	ticket := ToTicket(result)
	return ticket, nil
}
func (r *repositoryTicket) ListByOrderItem(ctx context.Context, orderItemID uint64) ([]Ticket, error) {
	rows, err := r.queries.ListTicketsByOrderItemID(ctx, orderItemID)
	if err != nil {
		return nil, err
	}
	tickets := ToTickets(rows)
	return tickets, nil
}
func (r *repositoryTicket) UpdateQrCode(ctx context.Context, id uint64, qrCode string) error {
	err := r.queries.UpdateTicketQRCode(ctx, database.UpdateTicketQRCodeParams{
		ID:     id,
		QrCode: qrCode,
	})
	if err != nil {
		return err
	}
	return nil
}
func (r *repositoryTicket) UpdateStatus(ctx context.Context, id uint64, status database.TicketsStatus) error {
	err := r.queries.UpdateTicketStatus(ctx, database.UpdateTicketStatusParams{
		ID:     id,
		Status: status,
	})
	if err != nil {
		return err
	}
	return nil
}
