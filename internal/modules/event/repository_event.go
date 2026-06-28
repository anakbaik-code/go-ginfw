package event

import (
	// "context"
	// "database/sql"
	"go-fwgin/internal/database"
)

type RepositoryEvent interface {
}

type repositoryEvent struct {
	queries *database.Queries
}

func NewRepositoryCategory(query *database.Queries) RepositoryEvent {
	return &repositoryEvent{queries: query}
}

// func (r *repositoryEvent)Create (ctx context.Context,e Event) error{
// 	var price sql.NullFloat64
// 	if e.Price > 0 {
// 		price = sql.NullFloat64{
// 			Float64: e.Price,
// 			Valid: true,
// 		}
// 	}
// 	result, err := r.queries.CreateEvent(ctx,database.CreateEventParams{
// 		CategoryID: e.CategoryID,
// 		UserID: e.UserID,
// 		Title: e.Title,
// 		Description: e.Description,
// 		Location: e.Location,
// 		StartTime: e.StartTime,
// 		EndTime: e.EndTime,
// 		Price: price,
		
// 	})
// }
