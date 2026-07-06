package admin

import "time"
type RevenuePerWeek struct {
	Year    int     `json:"year"`
	Week    int     `json:"week"`
	Revenue uint64 `json:"revenue"`
}

type RevenuePerMonth struct {
	Year      int     `json:"year"`
	Month     int     `json:"month"`
	MonthName string  `json:"month_name"`
	Revenue   uint64 `json:"revenue"`
}

type RevenuePerYear struct {
	Year    int     `json:"year"`
	Revenue uint64 `json:"revenue"`
}

type RevenuePerDate struct {
	Date    time.Time `json:"date"`
	Revenue uint64   `json:"revenue"`
}

type RevenueSummary struct {
	TotalRevenue    uint64           `json:"total_revenue"`
	RevenuePerWeek  []RevenuePerWeek  `json:"revenue_per_week"`
	RevenuePerMonth []RevenuePerMonth `json:"revenue_per_month"`
	RevenuePerYear  []RevenuePerYear  `json:"revenue_per_year"`
}

type RevenueSummaryResponse struct {
	TotalRevenue    uint64           `json:"total_revenue"`
	RevenuePerWeek  []RevenuePerWeek  `json:"revenue_per_week"`
	RevenuePerMonth []RevenuePerMonth `json:"revenue_per_month"`
	RevenuePerYear  []RevenuePerYear  `json:"revenue_per_year"`
}

type DashboardResponse struct {
    Stats   DashboardStats          `json:"stats"`
    Revenue RevenueSummaryResponse  `json:"revenue"`
}

type DashboardStats struct {
    TotalEvents            int64 `json:"total_events"`
    TotalEventsActive      int64 `json:"total_events_active"`
    TotalUsers             int64 `json:"total_users"`
    TotalOrganizers        int64 `json:"total_organizers"`
    TotalOrganizersActive  int64 `json:"total_organizers_active"`
}