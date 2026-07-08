package admin

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"golang.org/x/sync/errgroup"
)



type HandlerDashboardAdmin struct {
	service ServiceDashboardAdmin
}

func NewDashboardHandler(service ServiceDashboardAdmin) *HandlerDashboardAdmin {
	return &HandlerDashboardAdmin{
		service: service,
	}
}

func (h *HandlerDashboardAdmin) GetDashboardAdmin(c *gin.Context) {
	ctx := c.Request.Context()

	// Ambil semua data secara paralel
	var(
		totalEvents           int64
		totalEventsActive     int64
		totalUsers            int64
		totalOrganizers       int64
		totalOrganizersActive int64
		revenueSummary        *RevenueSummary
    )


	g, gctx := errgroup.WithContext(ctx)

	g.Go(func() error {
		v, err := h.service.TotalEvent(gctx)
		totalEvents = v
		return err
	})

	g.Go(func() error {
		v, err := h.service.TotalEventActive(gctx)
		totalEventsActive = v
		return err
	})

	g.Go(func() error {
		v, err := h.service.TotalUsers(gctx)
		totalUsers = v
		return err
	})

	g.Go(func() error {
		v, err := h.service.TotalOrganizerAll(gctx)
		totalOrganizers = v
		return err
	})

	g.Go(func() error {
		v, err := h.service.TotalOrganizerActive(gctx)
		totalOrganizersActive = v
		return err
	})

	g.Go(func() error {
		year := time.Now().Year()
		yearTime := time.Date(year, time.January, 1, 0, 0, 0, 0, time.UTC)
		v, err := h.service.GetRevenueSummary(gctx, yearTime)
		revenueSummary = v
		return err
	})

	if err := g.Wait(); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  "error",
			"message": "Failed to fetch dashboard data",
			"error":   err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status": "success",
		"data": DashboardResponse{
			Stats: DashboardStats{
				TotalEvents:           totalEvents,
				TotalEventsActive:     totalEventsActive,
				TotalUsers:            totalUsers,
				TotalOrganizers:       totalOrganizers,
				TotalOrganizersActive: totalOrganizersActive,
			},
			Revenue: ToRevenueSummaryResponse(revenueSummary),
		},
	})
}
