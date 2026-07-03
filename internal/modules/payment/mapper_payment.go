package payment

import "go-fwgin/internal/database"

// Mapper Pymet Repository
func ToPayment(p database.Payment) *Payment {
	return &Payment{
		ID:                    p.ID,
		OrderID:               p.OrderID,
		PaymentCode:           p.PaymentCode,
		Amount:                p.Amount,
		PaymentMethod:         p.PaymentMethod,
		PaymentProvider:       p.PaymentProvider,
		ProviderTransactionID: p.ProviderTransactionID.String,
		Status:                p.Status,
		PaidAt:                p.PaidAt.Time,
		ExpiredAt:             p.ExpiredAt.Time,
		CreatedAt:             p.CreatedAt.Time,
		UpdatedAt:             p.UpdatedAt.Time,
		DeletedAt:             p.DeletedAt.Time,
	}
}
func ToPayments(payments []database.Payment) []Payment {
	result := make([]Payment, 0, len(payments))

	for _, p := range payments {
		result = append(result, *ToPayment(p))
	}

	return result
}