package event

import (
	"go-fwgin/internal/config"
	"go-fwgin/internal/middleware"

	"github.com/gin-gonic/gin"
)

func (h *HandlerEvent) RoutesEvent(r *gin.RouterGroup, cfg *config.Config) {
	// Public routes
	events := r.Group("/events")
	{
		events.GET("", h.ListActive) // GET /events?page=1&limit=10
		events.GET("/:id", h.GetByID)
	}

	// Organizer routes
	organizer := r.Group("/events")
	organizer.Use(
		middleware.AuthJwtMiddleware(cfg),
		// middleware.RBAC("organizer"),
	)
	{
		organizer.POST("", h.Create)
		organizer.PUT("/:id", h.Update)
		organizer.DELETE("/:id", h.Delete)
		organizer.PATCH("/:id/status", h.UpdateStatus)

		organizer.GET("/inactive", h.ListInactive)
		organizer.GET("/cancelled", h.ListCancelled)
	}

	// My Events
	myEvents := r.Group("/my/events")
	myEvents.Use(
		middleware.AuthJwtMiddleware(cfg),
		// middleware.RBAC("organizer"),
	)
	{
		myEvents.GET("", h.ListMyEvents)
		myEvents.GET("/:id", h.GetMyEventByID)
	}
}
