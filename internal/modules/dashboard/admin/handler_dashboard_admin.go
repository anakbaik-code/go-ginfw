package admin

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
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
	type dashboardResult struct {
		totalEvents           int64
		totalEventsActive     int64
		totalUsers            int64
		totalOrganizers       int64
		totalOrganizersActive int64
		revenueSummary        *RevenueSummary
		err                   error
	}

	ch := make(chan dashboardResult, 1)

	go func() {
		var res dashboardResult

		// Parallel calls
		totalEvents, err := h.service.TotalEvent(ctx)
		if err != nil {
			res.err = err
			ch <- res
			return
		}
		res.totalEvents = totalEvents

		totalEventsActive, err := h.service.TotalEventActive(ctx)
		if err != nil {
			res.err = err
			ch <- res
			return
		}
		res.totalEventsActive = totalEventsActive

		totalUsers, err := h.service.TotalUsers(ctx)
		if err != nil {
			res.err = err
			ch <- res
			return
		}
		res.totalUsers = totalUsers

		totalOrganizers, err := h.service.TotalOrganizerAll(ctx)
		if err != nil {
			res.err = err
			ch <- res
			return
		}
		res.totalOrganizers = totalOrganizers

		totalOrganizersActive, err := h.service.TotalOrganizerActive(ctx)
		if err != nil {
			res.err = err
			ch <- res
			return
		}
		res.totalOrganizersActive = totalOrganizersActive

		// Revenue summary (opsional, bisa pake year from query)
		year := time.Now().Year()
		yearTime := time.Date(year, time.January, 1, 0, 0, 0, 0, time.UTC)
		revenue, err := h.service.GetRevenueSummary(ctx, yearTime)
		if err != nil {
			res.err = err
			ch <- res
			return
		}
		res.revenueSummary = revenue

		ch <- res
	}()

	res := <-ch
	if res.err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  "error",
			"message": "Failed to fetch dashboard data",
			"error":   res.err.Error(),
		})
		return
	}

	// Response ONE ENDPOINT dengan semua data
	c.JSON(http.StatusOK, gin.H{
		"status": "success",
		"data":    DashboardResponse{
			Stats: DashboardStats{
				TotalEvents:           res.totalEvents,
				TotalEventsActive:     res.totalEventsActive,
				TotalUsers:            res.totalUsers,
				TotalOrganizers:       res.totalOrganizers,
				TotalOrganizersActive: res.totalOrganizersActive,
			},
			Revenue: ToRevenueSummaryResponse(res.revenueSummary),
		},
	})
}
