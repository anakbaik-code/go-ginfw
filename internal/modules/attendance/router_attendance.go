package attendance

import (
	"go-fwgin/internal/config"
	"go-fwgin/internal/middleware"
	"go-fwgin/internal/pkg/jwt"

	"github.com/gin-gonic/gin"
)

func (h *HandlerAttendance) RegisterAttendanceRoutes(rg *gin.RouterGroup, cfg *config.Config) {
	jwtService := jwt.New(cfg.JwtSecret)
	attendance := rg.Group("/attendances")
	attendance.Use(middleware.AuthJwtMiddleware(jwtService))
	{
		// Check-in
		attendance.POST("/check-in", h.CheckIn)
		attendance.POST("/check-in/qr", h.CheckInByQR)
		attendance.POST("/check-in/manual", h.ManualCheckIn)

		// Read
		attendance.GET("/order/:order_id", h.GetAttendanceByOrderID)
		attendance.GET("/event/:event_id", h.GetAttendanceByEventID)
		attendance.GET("/user/:user_id", h.GetAttendanceByUserID)

		// Statistics
		attendance.GET("/event/:event_id/count", h.CountAttendeesByEvent)
		attendance.GET("/organizer/:organizer_id/count", h.CountAttendeesByOrganizer)
		attendance.GET("/ticket-type/:ticket_type_id/count", h.CountAttendeesByTicketType)
		attendance.GET("/organizer/:organizer_id/rate", h.GetEventAttendanceRate)
		attendance.GET("/event/:event_id/timeline/hourly", h.GetCheckInTimeline)
		attendance.GET("/event/:event_id/timeline/daily", h.GetCheckInTimelineByDate)

		// Update / Delete
		attendance.POST("/cancel", h.CancelAttendance)
		attendance.DELETE("/:id", h.SoftDeleteAttendance)
	}
}
