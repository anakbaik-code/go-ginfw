package ticket

import "go-fwgin/internal/database"

// 1. Mapper khusus untuk query GetAvailableTicketTypeByID
func ToTicketTypeFromAvailableRow(result database.GetAvailableTicketTypeByIDRow) *TicketType {
	return &TicketType{
		ID:                result.ID,
		EventID:           result.EventID,
		Name:              result.Name,
		Price:             result.Price,
		Quota:             result.Quota,
		MaxPerTransaction: uint32(result.MaxPerTransaction.Int32),
		StartSaleAt:       result.StartSaleAt.Time,
		EndSaleAt:         result.EndSaleAt.Time,
		CreatedAt:         result.CreatedAt.Time,
		UpdatedAt:         result.UpdatedAt.Time,
	}
}

// 2. Mapper khusus untuk query GetTicketTypeByID (Ganti nama biar gak bentrok)
func ToTicketTypeFromDetailRow(result database.GetTicketTypeByIDRow) *TicketType {
	return &TicketType{
		ID:                result.ID,
		EventID:           result.EventID,
		Name:              result.Name,
		Price:             result.Price,
		Quota:             result.Quota,
		MaxPerTransaction: uint32(result.MaxPerTransaction.Int32),
		StartSaleAt:       result.StartSaleAt.Time,
		EndSaleAt:         result.EndSaleAt.Time,
		CreatedAt:         result.CreatedAt.Time,
		UpdatedAt:         result.UpdatedAt.Time,
	}
}

func ToTicketType(result database.ListActiveTicketTypesRow) *TicketType {
	return &TicketType{
		ID:                result.ID,
		EventID:           result.EventID,
		Name:              result.Name,
		Price:             result.Price,
		Quota:             result.Quota,
		MaxPerTransaction: uint32(result.MaxPerTransaction.Int32),
		StartSaleAt:       result.StartSaleAt.Time,
		EndSaleAt:         result.EndSaleAt.Time,
		CreatedAt:         result.CreatedAt.Time,
		UpdatedAt:         result.UpdatedAt.Time,
	}
}

// 3. Mapper untuk slice/list database.TicketType murni
func ToTicketTypes(result []database.ListActiveTicketTypesRow) []TicketType {
	tickets := make([]TicketType, 0, len(result))
	for _, t := range result {
		tickets = append(tickets, *ToTicketType(t))
	}
	return tickets
}

func ToTicketTypeFromEventIDRow(result database.ListTicketTypesByEventIDRow)*TicketType{
	return &TicketType{
		ID:                result.ID,
		EventID:           result.EventID,
		Name:              result.Name,
		Price:             result.Price,
		Quota:             result.Quota,
		MaxPerTransaction: uint32(result.MaxPerTransaction.Int32),
		StartSaleAt:       result.StartSaleAt.Time,
		EndSaleAt:         result.EndSaleAt.Time,
		CreatedAt:         result.CreatedAt.Time,
		UpdatedAt:         result.UpdatedAt.Time,
	}
}
func ToTicketTypeFromEventList(result []database.ListTicketTypesByEventIDRow) []TicketType {
	tickets := make([]TicketType, 0, len(result))
	for _, t := range result {
		tickets = append(tickets, *ToTicketTypeFromEventIDRow(t))
	}
	return tickets
}