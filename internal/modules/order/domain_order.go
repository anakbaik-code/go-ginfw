package order

import (
	"database/sql"
	"go-fwgin/internal/database"
	"time"
)

type Order struct {
	ID          uint64
	UserID      uint64
	TotalAmount uint32
	Status      database.OrdersStatus
	CreatedAt   time.Time
	UpdatedAt   time.Time
	DeletedAt   sql.NullTime
}
