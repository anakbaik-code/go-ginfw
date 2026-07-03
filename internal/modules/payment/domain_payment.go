package payment

import (
	"go-fwgin/internal/database"
	"time"
)

type Payment struct {
	ID                    uint64
	OrderID               uint64
	PaymentCode           string
	Amount                uint32
	PaymentMethod         string
	PaymentProvider       string
	ProviderTransactionID string
	Status                database.PaymentsStatus
	PaidAt                time.Time
	ExpiredAt             time.Time
	CreatedAt             time.Time
	UpdatedAt             time.Time
	DeletedAt             time.Time
}
