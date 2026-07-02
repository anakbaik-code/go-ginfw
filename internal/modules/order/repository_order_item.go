package order

import (
	"context"
	"go-fwgin/internal/database"
)

type RepositoryOrderItem interface {
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
func (r *repositoryOrder) DeleteByOrderID(ctx context.Context, orderID uint64) error {
	err := r.queries.DeleteOrderItemsByOrderID(ctx, orderID)
	if err != nil {
		return err
	}
	return nil
}
func (r *repositoryOrder) GetOrderItemByID(ctx context.Context, id uint64) (*OrderItem, error) {
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

