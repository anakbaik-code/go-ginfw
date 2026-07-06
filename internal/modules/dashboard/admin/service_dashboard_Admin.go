package admin

import (
	"context"
	"go-fwgin/internal/database"
	"go-fwgin/internal/modules/event"
	"go-fwgin/internal/modules/order"
	"go-fwgin/internal/modules/payment"
	"go-fwgin/internal/modules/user"
	"time"
)

type ServiceDashboardAdmin interface {
	TotalEvent(ctx context.Context) (int64, error)
	TotalEventActive(ctx context.Context) (int64, error)
	TotalUsers(ctx context.Context) (int64, error)
	TotalOrganizerActive(ctx context.Context) (int64, error)
	TotalOrganizerAll(ctx context.Context) (int64, error)
	ListOrganizer(ctx context.Context, limit int32, page int32) ([]user.UserResponse, int64, error)
	GetRevenueSummary(ctx context.Context, year time.Time) (*RevenueSummary, error)
}
type serviceDashboardAdmin struct {
	adminRepo   RepositoryDashboardAdmin
	userRepo    user.UserRepository
	eventRepo   event.RepositoryEvent
	orderRepo   order.RepositoryOrder
	paymentRepo payment.RepositoryPayment
}

func NewServiceDashboardAdmin(
	ar RepositoryDashboardAdmin,
	ur user.UserRepository,
	er event.RepositoryEvent,
	or order.RepositoryOrder,
	pr payment.RepositoryPayment) ServiceDashboardAdmin {
	return &serviceDashboardAdmin{
		adminRepo:   ar,
		userRepo:    ur,
		eventRepo:   er,
		orderRepo:   or,
		paymentRepo: pr,
	}
}

// Event Dashboard
func (s *serviceDashboardAdmin) TotalEvent(ctx context.Context) (int64, error) {
	result, err := s.eventRepo.CountEvent(ctx)
	if err != nil {
		return 0, err
	}
	return result, nil
}
func (s *serviceDashboardAdmin) TotalEventActive(ctx context.Context) (int64, error) {
	status := database.EventsStatusActive
	result, err := s.eventRepo.CountByStatus(ctx, event.EventStatus{
		Status: status,
	})
	if err != nil {
		return 0, err
	}
	return result, nil
}

// User Dashboard
func (s *serviceDashboardAdmin) TotalUsers(ctx context.Context) (int64, error) {
	result, err := s.userRepo.CountUsersActive(ctx)
	if err != nil {
		return 0, err
	}
	return result, nil
}

// Organizer
func (s *serviceDashboardAdmin) TotalOrganizerActive(ctx context.Context) (int64, error) {
	result, err := s.userRepo.CountOrganizerActive(ctx)
	if err != nil {
		return 0, err
	}
	return result, nil
}
func (s *serviceDashboardAdmin) TotalOrganizerAll(ctx context.Context) (int64, error) {
	result, err := s.userRepo.CountOrganizerAll(ctx)
	if err != nil {
		return 0, err
	}
	return result, nil
}
func (s *serviceDashboardAdmin) ListOrganizer(ctx context.Context, limit int32, page int32) ([]user.UserResponse, int64, error) {
	if limit <= 0 {
		limit = 10
	}
	if page <= 0 {
		page = 1
	}
	if limit > 100 {
		limit = 100
	}
	offset := (page - 1) * limit

	users, err := s.userRepo.ListOrganizer(ctx, limit, offset)
	if err != nil {
		return nil, 0, err
	}

	// 4. Ambil total count (buat pagination)
	total, err := s.userRepo.CountOrganizerAll(ctx)
	if err != nil {
		return nil, 0, err
	}

	resUser := user.ToUserResponses(users)
	return resUser, total, nil
}
func (s *serviceDashboardAdmin) GetRevenueSummary(ctx context.Context, year time.Time) (*RevenueSummary, error) {
	type result struct {
		total    uint64
		perWeek  []RevenuePerWeek
		perMonth []RevenuePerMonth
		perYear  []RevenuePerYear
		err      error
	}

	ch := make(chan result, 1)

	go func() {
		var res result

		// Panggil method-method dari layer repository (s.adminRepo)
		total, err := s.adminRepo.TotalRevenue(ctx)
		if err != nil {
			res.err = err
			ch <- res
			return
		}
		res.total = total

		perWeek, err := s.adminRepo.RevenuePerWeek(ctx, year)
		if err != nil {
			res.err = err
			ch <- res
			return
		}
		res.perWeek = perWeek

		perMonth, err := s.adminRepo.RevenuePerMonth(ctx, year)
		if err != nil {
			res.err = err
			ch <- res
			return
		}
		res.perMonth = perMonth

		perYear, err := s.adminRepo.RevenuePerYear(ctx)
		if err != nil {
			res.err = err
			ch <- res
			return
		}
		res.perYear = perYear

		ch <- res
	}()

	res := <-ch
	if res.err != nil {
		return nil, res.err
	}

	return &RevenueSummary{
		TotalRevenue:    res.total,
		RevenuePerWeek:  res.perWeek,
		RevenuePerMonth: res.perMonth,
		RevenuePerYear:  res.perYear,
	}, nil
}
