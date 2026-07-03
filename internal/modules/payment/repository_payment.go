package payment

import (
	"context"
	"database/sql"
	"go-fwgin/internal/database"
)

type RepositoryPayment interface {
	Count(ctx context.Context) (int64, error)
	Create(ctx context.Context, p Payment) (int64, error)
	Delete(ctx context.Context, id uint64) error
	GetByCode(ctx context.Context, payCode string) (*Payment, error)
	GetById(ctx context.Context, id uint64) (*Payment, error)
	GetByOrderID(ctx context.Context, orderID uint64) (*Payment, error)
	List(ctx context.Context, limit int32, offset int32) ([]Payment, error)
	ListByUserID(ctx context.Context, userID uint64, limit int32, offset int32) ([]Payment, error)
	MarkPaid(ctx context.Context, id uint64) error
	UpdateStatus(ctx context.Context, id uint64, p Payment) error
	UpdateProvTransID(ctx context.Context, id uint64, p Payment) error
}
type repositoryPayment struct {
	queries *database.Queries
}

func NewRepositoryPayment(q *database.Queries) RepositoryPayment {
	return &repositoryPayment{queries: q}
}
func (r *repositoryPayment) Count(ctx context.Context) (int64, error) {
	result, err := r.queries.CountPayments(ctx)
	if err != nil {
		return 0, err
	}
	return result, nil
}
func (r *repositoryPayment) Create(ctx context.Context, p Payment) (int64, error) {
	providerTransID := sql.NullString{
		String: p.ProviderTransactionID,
		Valid:  true,
	}
	paidAt := sql.NullTime{
		Time:  p.PaidAt,
		Valid: true,
	}
	expiredAt := sql.NullTime{
		Time:  p.ExpiredAt,
		Valid: true,
	}
	result, err := r.queries.CreatePayment(ctx, database.CreatePaymentParams{
		OrderID:               p.ID,
		PaymentCode:           p.PaymentCode,
		Amount:                p.Amount,
		PaymentMethod:         p.PaymentMethod,
		PaymentProvider:       p.PaymentProvider,
		ProviderTransactionID: providerTransID,
		Status:                p.Status,
		PaidAt:                paidAt,
		ExpiredAt:             expiredAt,
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
func (r *repositoryPayment) Delete(ctx context.Context, id uint64) error {
	err := r.queries.DeletePayment(ctx, id)
	if err != nil {
		return err
	}
	return nil
}
func (r *repositoryPayment) GetByCode(ctx context.Context, payCode string) (*Payment, error) {
	result, err := r.queries.GetPaymentByCode(ctx, payCode)
	if err != nil {
		return nil, err
	}
	payment := ToPayment(result)
	return payment, nil
}
func (r *repositoryPayment) GetById(ctx context.Context, id uint64) (*Payment, error) {
	result, err := r.queries.GetPaymentByID(ctx, id)
	if err != nil {
		return nil, err
	}
	payment := ToPayment(result)
	return payment, nil
}
func (r *repositoryPayment) GetByOrderID(ctx context.Context, orderID uint64) (*Payment, error) {
	result, err := r.queries.GetPaymentByOrderID(ctx, orderID)
	if err != nil {
		return nil, err
	}
	payment := ToPayment(result)
	return payment, nil
}
func (r *repositoryPayment) List(ctx context.Context, limit int32, offset int32) ([]Payment, error) {
	result, err := r.queries.ListPayments(ctx, database.ListPaymentsParams{
		Limit:  limit,
		Offset: offset,
	})
	if err != nil {
		return nil, err
	}
	payments := ToPayments(result)
	return payments, nil
}
func (r *repositoryPayment) ListByUserID(ctx context.Context, userID uint64, limit int32, offset int32) ([]Payment, error) {
	result, err := r.queries.ListPaymentsByUserID(ctx, database.ListPaymentsByUserIDParams{
		UserID: userID,
		Limit:  limit,
		Offset: offset,
	})
	if err != nil {
		return nil, err
	}
	payments := ToPayments(result)
	return payments, nil
}
func (r *repositoryPayment) MarkPaid(ctx context.Context, id uint64) error {
	err := r.queries.MarkPaymentAsPaid(ctx, id)
	if err != nil {
		return err
	}
	return nil
}
func (r *repositoryPayment) UpdateStatus(ctx context.Context, id uint64, p Payment) error {
	err := r.queries.UpdatePaymentStatus(ctx, database.UpdatePaymentStatusParams{
		ID:     id,
		Status: p.Status,
	})
	if err != nil {
		return err
	}
	return nil
}
func (r *repositoryPayment) UpdateProvTransID(ctx context.Context, id uint64, p Payment) error {
	provTransID := sql.NullString{
		String: p.ProviderTransactionID,
		Valid:  true,
	}
	err := r.queries.UpdateProviderTransactionID(ctx, database.UpdateProviderTransactionIDParams{
		ID:                    id,
		ProviderTransactionID: provTransID,
	})
	if err != nil {
		return err
	}
	return nil
}
