package ticket

import "go-fwgin/internal/database"

func ToTicket(t database.Ticket) *Ticket {
	return &Ticket{
		ID:          t.ID,
		OrderItemID: t.OrderItemID,
		TicketCode:  t.TicketCode,
		QrCode:      t.QrCode,
		Status:      t.Status,
		CheckInAt:   t.CheckedInAt.Time,
	}
}
func ToTickets(tickets []database.Ticket) []Ticket {
	result := make([]Ticket, 0, len(tickets))
	for _, t := range tickets {
		result = append(result, *ToTicket(t))
	}
	return result 
}
