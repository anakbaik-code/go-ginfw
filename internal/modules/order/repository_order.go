package order

import (
	"context"
	"go-fwgin/internal/database"
)

type RepositoryOrder interface {
	Count(ctx context.Context) (int64, error)
	CountByUserID(ctx context.Context, userID uint64) (int64, error)
	Create(ctx context.Context, o Order) (uint64, error)
	Delete(ctx context.Context, id uint64) error
	GetByID(ctx context.Context, id uint64) (*Order, error)
	ListByUserID(ctx context.Context, userID uint64, limit int32, offset int32) ([]Order, error)
	UpdateStatus(ctx context.Context, o Order) error
}
type repositoryOrder struct {
	queries *database.Queries
}

func NewRepositoryOrder(q *database.Queries) RepositoryOrder {
	return &repositoryOrder{
		queries: q,
	}
}

func (r *repositoryOrder) Count(ctx context.Context) (int64, error) {
	result, err := r.queries.CountOrders(ctx)
	if err != nil {
		return 0, err
	}
	return result, nil
}

func (r *repositoryOrder) CountByUserID(ctx context.Context, userID uint64) (int64, error) {
	result, err := r.queries.CountOrdersByUserID(ctx, userID)
	if err != nil {
		return 0, err
	}
	return result, nil
}
func (r *repositoryOrder) Create(ctx context.Context, o Order) (uint64, error) {
	result, err := r.queries.CreateOrder(ctx, database.CreateOrderParams{
		UserID:      o.ID,
		TotalAmount: o.TotalAmount,
		Status:      o.Status,
	})
	if err != nil {
		return 0, err
	}
	idLast, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}
	return uint64(idLast), nil
}
func (r *repositoryOrder) Delete(ctx context.Context, id uint64) error {
	if err := r.queries.DeleteOrder(ctx, id); err != nil {
		return err
	}
	return nil
}
func (r *repositoryOrder) GetByID(ctx context.Context, id uint64) (*Order, error) {
	result, err := r.queries.GetOrderByID(ctx, id)
	if err != nil {
		return nil, err
	}
	order := &Order{
		ID:          result.ID,
		UserID:      result.UserID,
		TotalAmount: result.TotalAmount,
		Status:      result.Status,
		CreatedAt:   result.CreatedAt.Time,
		UpdatedAt:   result.UpdatedAt.Time,
	}
	return order, nil
}
func (r *repositoryOrder) ListByUserID(ctx context.Context, userID uint64, limit int32, offset int32) ([]Order, error) {
	rows, err := r.queries.ListOrdersByUserID(ctx, database.ListOrdersByUserIDParams{
		UserID: userID,
		Limit:  limit,
		Offset: offset,
	})
	if err != nil {
		return nil, err
	}
	orders := make([]Order, 0, len(rows))
	for _, row := range rows {
		orders = append(orders, Order{
			ID:          row.ID,
			UserID:      row.UserID,
			TotalAmount: row.TotalAmount,
			Status:      row.Status,
			CreatedAt:   row.CreatedAt.Time,
			UpdatedAt:   row.UpdatedAt.Time,
		})
	}
	return orders, nil
}

func (r *repositoryOrder) UpdateStatus(ctx context.Context, o Order) error {
	err := r.queries.UpdateOrderStatus(ctx, database.UpdateOrderStatusParams{
		ID:     o.ID,
		Status: o.Status,
	})
	if err != nil {
		return err
	}
	return nil
}
