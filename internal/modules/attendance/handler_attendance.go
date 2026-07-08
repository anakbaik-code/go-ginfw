package attendance

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type HandlerAttendance struct {
	service ServiceAttendance
}

func NewHandlerAttendance(service ServiceAttendance) *HandlerAttendance {
	return &HandlerAttendance{service: service}
}

// ============================================
// CREATE / CHECK-IN
// ============================================

// POST /admin/attendances/check-in
func (h *HandlerAttendance) CheckIn(c *gin.Context) {
	var req CheckInRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": "error", "message": err.Error()})
		return
	}

	err := h.service.CheckIn(c.Request.Context(), req)
	if err != nil {
		h.handleCheckInError(c, err)
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": "success", "message": "Check-in successful"})
}

// POST /admin/attendances/check-in/qr
func (h *HandlerAttendance) CheckInByQR(c *gin.Context) {
	var req CheckInByQRRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": "error", "message": err.Error()})
		return
	}

	err := h.service.CheckInByQR(c.Request.Context(), req)
	if err != nil {
		h.handleCheckInError(c, err)
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": "success", "message": "Check-in via QR successful"})
}

// POST /admin/attendances/check-in/manual
func (h *HandlerAttendance) ManualCheckIn(c *gin.Context) {
	var req ManualCheckInRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": "error", "message": err.Error()})
		return
	}

	err := h.service.ManualCheckIn(c.Request.Context(), req)
	if err != nil {
		h.handleCheckInError(c, err)
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": "success", "message": "Manual check-in successful"})
}

// ============================================
// READ
// ============================================

// GET /admin/attendances/order/:order_id
func (h *HandlerAttendance) GetAttendanceByOrderID(c *gin.Context) {
	orderID, err := strconv.ParseUint(c.Param("order_id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": "error", "message": "invalid order_id"})
		return
	}

	result, err := h.service.GetAttendanceByOrderID(c.Request.Context(), orderID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"status": "error", "message": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": "success", "data": result})
}

// GET /admin/attendances/event/:event_id
func (h *HandlerAttendance) GetAttendanceByEventID(c *gin.Context) {
	eventID, err := strconv.ParseUint(c.Param("event_id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": "error", "message": "invalid event_id"})
		return
	}

	result, err := h.service.GetAttendanceByEventID(c.Request.Context(), eventID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"status": "error", "message": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": "success", "data": result})
}

// GET /admin/attendances/user/:user_id
func (h *HandlerAttendance) GetAttendanceByUserID(c *gin.Context) {
	userID, err := strconv.ParseUint(c.Param("user_id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": "error", "message": "invalid user_id"})
		return
	}

	result, err := h.service.GetAttendanceByUserID(c.Request.Context(), userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"status": "error", "message": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": "success", "data": result})
}

// ============================================
// STATISTICS
// ============================================

// GET /admin/attendances/event/:event_id/count
func (h *HandlerAttendance) CountAttendeesByEvent(c *gin.Context) {
	eventID, err := strconv.ParseUint(c.Param("event_id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": "error", "message": "invalid event_id"})
		return
	}

	count, err := h.service.CountAttendeesByEvent(c.Request.Context(), eventID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"status": "error", "message": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": "success", "data": gin.H{"count": count}})
}

// GET /admin/attendances/organizer/:organizer_id/count
func (h *HandlerAttendance) CountAttendeesByOrganizer(c *gin.Context) {
	organizerID, err := strconv.ParseUint(c.Param("organizer_id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": "error", "message": "invalid organizer_id"})
		return
	}

	count, err := h.service.CountAttendeesByOrganizer(c.Request.Context(), organizerID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"status": "error", "message": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": "success", "data": gin.H{"count": count}})
}

// GET /admin/attendances/ticket-type/:ticket_type_id/count
func (h *HandlerAttendance) CountAttendeesByTicketType(c *gin.Context) {
	ticketTypeID, err := strconv.ParseUint(c.Param("ticket_type_id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": "error", "message": "invalid ticket_type_id"})
		return
	}

	count, err := h.service.CountAttendeesByTicketType(c.Request.Context(), ticketTypeID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"status": "error", "message": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": "success", "data": gin.H{"count": count}})
}

// GET /admin/attendances/organizer/:organizer_id/rate
func (h *HandlerAttendance) GetEventAttendanceRate(c *gin.Context) {
	organizerID, err := strconv.ParseUint(c.Param("organizer_id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": "error", "message": "invalid organizer_id"})
		return
	}

	result, err := h.service.GetEventAttendanceRate(c.Request.Context(), organizerID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"status": "error", "message": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": "success", "data": result})
}

// GET /admin/attendances/event/:event_id/timeline/hourly
func (h *HandlerAttendance) GetCheckInTimeline(c *gin.Context) {
	eventID, err := strconv.ParseUint(c.Param("event_id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": "error", "message": "invalid event_id"})
		return
	}

	result, err := h.service.GetCheckInTimeline(c.Request.Context(), eventID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"status": "error", "message": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": "success", "data": result})
}

// GET /admin/attendances/event/:event_id/timeline/daily
func (h *HandlerAttendance) GetCheckInTimelineByDate(c *gin.Context) {
	eventID, err := strconv.ParseUint(c.Param("event_id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": "error", "message": "invalid event_id"})
		return
	}

	result, err := h.service.GetCheckInTimelineByDate(c.Request.Context(), eventID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"status": "error", "message": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": "success", "data": result})
}

// ============================================
// UPDATE / DELETE
// ============================================

// POST /admin/attendances/cancel
func (h *HandlerAttendance) CancelAttendance(c *gin.Context) {
	var req CancelAttendanceRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": "error", "message": err.Error()})
		return
	}

	if err := h.service.CancelAttendance(c.Request.Context(), req); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"status": "error", "message": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": "success", "message": "Attendance cancelled"})
}

// DELETE /admin/attendances/:id
func (h *HandlerAttendance) SoftDeleteAttendance(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": "error", "message": "invalid id"})
		return
	}

	if err := h.service.SoftDeleteAttendance(c.Request.Context(), id); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"status": "error", "message": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": "success", "message": "Attendance deleted"})
}

// ============================================
// ERROR MAPPING
// ============================================

// handleCheckInError maps domain errors to proper HTTP status codes,
// instead of always returning 500 for business-rule violations.
func (h *HandlerAttendance) handleCheckInError(c *gin.Context, err error) {
	switch {
	case errors.Is(err, ErrOrderNotPaid):
		c.JSON(http.StatusUnprocessableEntity, gin.H{"status": "error", "message": err.Error()})
	case errors.Is(err, ErrAlreadyCheckedIn):
		c.JSON(http.StatusConflict, gin.H{"status": "error", "message": err.Error()})
	case errors.Is(err, ErrEventNotStarted):
		c.JSON(http.StatusUnprocessableEntity, gin.H{"status": "error", "message": err.Error()})
	case errors.Is(err, ErrEventEnded):
		c.JSON(http.StatusUnprocessableEntity, gin.H{"status": "error", "message": err.Error()})
	default:
		c.JSON(http.StatusInternalServerError, gin.H{"status": "error", "message": err.Error()})
	}
}