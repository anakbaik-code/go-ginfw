package organizer

import "go-fwgin/internal/database"

type RepositoryDashboardOrganizer interface {
	
}

type repositoryDashboardOrganizer struct {
	queries *database.Queries
}

func NewRepositoryDashboardOrganizer(q *database.Queries)RepositoryDashboardOrganizer{
	return  &repositoryDashboardOrganizer{queries: q}
}

