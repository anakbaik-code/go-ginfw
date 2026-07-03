package order

import (
	"context"
	"go-fwgin/internal/database"
)

type RepositoryOrderItem interface {
	CountByOrderID(ctx context.Context, orderID uint64) (int64, error)
	Create(ctx context.Context, oi OrderItem) (uint64, error)
	Delete(ctx context.Context, id uint64) error
	DeleteByOrderID(ctx context.Context, orderID uint64) error
	GetOrderItemByID(ctx context.Context, id uint64) (*OrderItem, error)
	ListByOrderID(ctx context.Context, orderID uint64) ([]OrderItem, error)
}
type repositoryOrderItem struct {
	queries *database.Queries
}

func NewRepositoryOrderItem(q *database.Queries) RepositoryOrderItem {
	return &repositoryOrderItem{queries: q}
}

func (r *repositoryOrderItem) CountByOrderID(ctx context.Context, orderID uint64) (int64, error) {
	result, err := r.queries.CountOrderItemsByOrderID(ctx, orderID)
	if err != nil {
		return 0, err
	}
	return result, nil
}
func (r *repositoryOrderItem) Create(ctx context.Context, oi OrderItem) (uint64, error) {
	result, err := r.queries.CreateOrderItem(ctx, database.CreateOrderItemParams{
		OrderID:      oi.OrderID,
		EventID:      oi.EventID,
		TicketTypeID: oi.TicketTypeID,
		Quantity:     oi.Quantity,
		Price:        oi.Price,
	})
	if err != nil {
		return 0, err
	}
	lastID, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}
	return uint64(lastID), nil
}
func (r *repositoryOrderItem) Delete(ctx context.Context, id uint64) error {
	if err := r.queries.DeleteOrderItem(ctx, id); err != nil {
		return err
	}
	return nil
}
func (r *repositoryOrderItem) DeleteByOrderID(ctx context.Context, orderID uint64) error {
	err := r.queries.DeleteOrderItemsByOrderID(ctx, orderID)
	if err != nil {
		return err
	}
	return nil
}
func (r *repositoryOrderItem) GetOrderItemByID(ctx context.Context, id uint64) (*OrderItem, error) {
	result, err := r.queries.GetOrderItemByID(ctx, id)
	if err != nil {
		return nil, err
	}
	orderItem := &OrderItem{
		ID:           result.ID,
		OrderID:      result.OrderID,
		EventID:      result.EventID,
		TicketTypeID: result.TicketTypeID,
		Quantity:     result.Quantity,
		Price:        result.Price,
	}
	return orderItem, nil
}

func (r *repositoryOrderItem) ListByOrderID(ctx context.Context, orderID uint64) ([]OrderItem, error) {
	rows, err := r.queries.ListOrderItemsByOrderID(ctx, orderID)
	if err != nil {
		return nil, err
	}
	orderItems := make([]OrderItem, 0, len(rows))
	for _, row := range rows {
		orderItems = append(orderItems, OrderItem{
			ID:           row.ID,
			OrderID:      row.OrderID,
			EventID:      row.EventID,
			TicketTypeID: row.TicketTypeID,
			Quantity:     row.Quantity,
			Price:        row.Price,
		})
	}
	return orderItems, nil
}
