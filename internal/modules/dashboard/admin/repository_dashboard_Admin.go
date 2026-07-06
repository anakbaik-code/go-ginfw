package admin

import (
	"context"
	"database/sql"
	"fmt"
	"go-fwgin/internal/database"
	"time"
)

type RepositoryDashboardAdmin interface {
	TotalRevenue(ctx context.Context) (uint64, error)
	RevenuePerWeek(ctx context.Context, year time.Time) ([]RevenuePerWeek, error)
	RevenuePerMonth(ctx context.Context, year time.Time) ([]RevenuePerMonth, error)
	RevenuePerYear(ctx context.Context) ([]RevenuePerYear, error)
	RevenuePerDateRange(ctx context.Context, startDate, endDate time.Time) ([]RevenuePerDate, error)
}
type repositoryDashboardAdmin struct {
	queries *database.Queries
}

func NewRepositoryDashboardAdmin(q *database.Queries) RepositoryDashboardAdmin {
	return &repositoryDashboardAdmin{queries: q}
}

func (r *repositoryDashboardAdmin) TotalRevenue(ctx context.Context) (uint64, error) {
	val, err := r.queries.SumRevenue(ctx)
	if err != nil {
		return 0, err
	}
	if val < 0 {
		return 0, fmt.Errorf("negative revenue value")
	}

	return uint64(val), nil
}

// 2. Get Revenue Per Week
func (r *repositoryDashboardAdmin) RevenuePerWeek(ctx context.Context, year time.Time) ([]RevenuePerWeek, error) {
	yearValid := sql.NullTime{
		Time:  year,
		Valid: true,
	}
	rows, err := r.queries.RevenuePerWeek(ctx, yearValid)
	if err != nil {
		return nil, err
	}

	result := make([]RevenuePerWeek, len(rows))
	for i, row := range rows {
		result[i] = RevenuePerWeek{
			Year:    int(row.Year),
			Week:    int(row.Week),
			Revenue: uint64(row.Revenue),
		}
	}
	return result, nil
}

// 3. Get Revenue Per Month
func (r *repositoryDashboardAdmin) RevenuePerMonth(ctx context.Context, year time.Time) ([]RevenuePerMonth, error) {
	yearValid := sql.NullTime{
		Time:  year,
		Valid: true,
	}
	rows, err := r.queries.RevenuePerMonth(ctx, yearValid)
	if err != nil {
		return nil, err
	}

	// Fix 1: Ubah key map menjadi int agar cocok dengan angka bulan
	monthNames := map[int]string{
		1: "January", 2: "February", 3: "March", 4: "April",
		5: "May", 6: "June", 7: "July", 8: "August",
		9: "September", 10: "October", 11: "November", 12: "December",
	}

	result := make([]RevenuePerMonth, len(rows))
	for i, row := range rows {
		monthInt := int(row.Month)

		result[i] = RevenuePerMonth{
			Year:      int(row.Year),
			Month:     monthInt,
			MonthName: monthNames[monthInt],
			Revenue:   uint64(row.Revenue),
		}
	}
	return result, nil
}

// 4. Get Revenue Per Year
func (r *repositoryDashboardAdmin) RevenuePerYear(ctx context.Context) ([]RevenuePerYear, error) {
	rows, err := r.queries.RevenuePerYear(ctx)
	if err != nil {
		return nil, err
	}

	result := make([]RevenuePerYear, len(rows))
	for i, row := range rows {
		result[i] = RevenuePerYear{
			Year:    int(row.Year),
			Revenue: uint64(row.Revenue),
		}
	}
	return result, nil
}

// 5. Get Revenue Per Date Range
func (r *repositoryDashboardAdmin) RevenuePerDateRange(ctx context.Context, startDate, endDate time.Time) ([]RevenuePerDate, error) {
	start := sql.NullTime{
		Time:  startDate,
		Valid: true,
	}
	// Tambahkan ini untuk menampung endDate
	end := sql.NullTime{
		Time:  endDate,
		Valid: true,
	}

	rows, err := r.queries.RevenuePerDateRange(ctx, database.RevenuePerDateRangeParams{
		FromCreatedAt: start,
		ToCreatedAt:   end, // Isi dengan variabel end
	})
	if err != nil {
		return nil, err
	}

	result := make([]RevenuePerDate, len(rows))
	for i, row := range rows {

		result[i] = RevenuePerDate{
			Date:    row.Date,
			Revenue: uint64(row.Revenue),
		}
	}
	return result, nil
}

// 6. Get Complete Revenue Summary (semua data dalam 1 panggilan)
