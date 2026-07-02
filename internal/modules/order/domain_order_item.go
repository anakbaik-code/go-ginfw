package order

type OrderItem struct {
	ID           uint64
	OrderID      uint64
	EventID      uint64
	TicketTypeID uint64
	Quantity     uint32
	Price        uint32
}
