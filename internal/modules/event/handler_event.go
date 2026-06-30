package event

import (
	"go-fwgin/internal/response"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type HandlerEvent struct {
	service ServiceEvent
}

func NewHandlerEvent(svc ServiceEvent) *HandlerEvent {
	return &HandlerEvent{service: svc}
}

func (h *HandlerEvent) Create(c *gin.Context) {
	var req RequestCreateEvent
	if err := c.ShouldBindJSON(&req); err != nil {
		response.Error(c, http.StatusBadRequest, "Invalid Request PAyload", err.Error())
		return
	}
	userId, ok := c.Get("user_id")
	if !ok {
		response.Error(c, http.StatusUnauthorized, "unauthorized", "user id not found")
	}
	id, err := h.service.CreateEvent(c.Request.Context(), userId.(uint64), req)
	if err != nil {
		response.Error(c, http.StatusInternalServerError, "failed to create event", err.Error())
		return
	}

	response.Success(c, http.StatusCreated, "event created successfully", gin.H{
		"id": id,
	})
}

func (h *HandlerEvent) Update(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 0, 64)
	if err != nil {
		response.Error(c, http.StatusBadRequest, "Invalid Id Request", err.Error())
		return
	}

	var req RequestUpdateEvent
	if err := c.ShouldBindJSON(&req); err != nil {
		response.Error(c, http.StatusBadRequest, "Invalid Request Payload", err.Error())
		return
	}
	// konversi any to uint pakai .()
	userId := c.MustGet("user_id").(uint64)
	err = h.service.UpdateEvent(c.Request.Context(), id, userId, req)
	if err != nil {
		response.Error(c, http.StatusInternalServerError, "failed to update event", err.Error())
		return
	}
	response.Success(c, http.StatusOK, "event updated successfully", nil)
}
func (h *HandlerEvent) Delete(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 0, 64)
	if err != nil {
		response.Error(c, http.StatusBadRequest, "Invalid Id Request", err.Error())
		return
	}
	err = h.service.DeleteEvent(c.Request.Context(), id)
	if err != nil {
		response.Error(c, http.StatusInternalServerError, "Failed To delete Event", err.Error())
		return
	}
	response.Success(c, http.StatusOK, "Event Deleted Succesfully", nil)
}
func (h *HandlerEvent) UpdateStatus(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 0, 64)
	if err != nil {
		response.Error(c, http.StatusBadRequest, "Invalid Id Request", err.Error())
		return
	}
	var req RequestUpdateEventStatus
	err = h.service.UpdateStatus(c.Request.Context(), req, id)
	if err != nil {
		response.Error(c, http.StatusInternalServerError, "Failed Update Status", err.Error())
		return
	}
	response.Success(c, http.StatusOK, "Event UpdateD Status Successfully", nil)
}
func (h *HandlerEvent) GetByID(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 0, 64)
	if err != nil {
		response.Error(c, http.StatusBadRequest, "Invalid Id Request", err.Error())
		return
	}
	event, err := h.service.GetByID(c.Request.Context(), id)
	if err != nil {
		response.Error(c, http.StatusInternalServerError, "Failed Get Event", err.Error())
		return
	}
	response.Success(c, http.StatusOK, "Succesfully fetch event", event)
}
func (h *HandlerEvent) ListActive(c *gin.Context) {
	var req RequestListEvent
	if err := c.ShouldBindQuery(&req); err != nil {
		response.Error(c, http.StatusBadRequest, "Invalid Id Request", err.Error())
		return
	}
	events, total, err := h.service.ListActive(c.Request.Context(), req.Page, req.Limit)
	if err != nil {
		response.Error(c, http.StatusInternalServerError, "Failed Get List Active Events", err.Error())
		return
	}
	responseData := gin.H{
		"events": events,
		"total":  total,
		"page":   req.Page,
		"limit":  req.Limit,
	}
	response.Success(c, http.StatusOK, "Success Fetch data list events active", responseData)
}

func (h *HandlerEvent) ListInactive(c *gin.Context) {
	var req RequestListEvent
	if err := c.ShouldBindQuery(&req); err != nil {
		response.Error(c, http.StatusBadRequest, "Invalid Id Request", err.Error())
		return
	}
	events, total, err := h.service.ListInActive(c.Request.Context(), req.Page, req.Limit)
	if err != nil {
		response.Error(c, http.StatusInternalServerError, "Failed Get List InActive Events", err.Error())
		return
	}
	responseData := gin.H{
		"events": events,
		"total":  total,
		"page":   req.Page,
		"limit":  req.Limit,
	}
	response.Success(c, http.StatusOK, "Success Fetch data list events InActive", responseData)
}
func (h *HandlerEvent) ListCancelled(c *gin.Context) {
	var req RequestListEvent
	if err := c.ShouldBindQuery(&req); err != nil {
		response.Error(c, http.StatusBadRequest, "Invalid Id Request", err.Error())
		return
	}
	events, total, err := h.service.ListCancelled(c.Request.Context(), req.Page, req.Limit)
	if err != nil {
		response.Error(c, http.StatusInternalServerError, "Failed Get List Cancelled Events", err.Error())
		return
	}
	responseData := gin.H{
		"events": events,
		"total":  total,
		"page":   req.Page,
		"limit":  req.Limit,
	}
	response.Success(c, http.StatusOK, "Success Fetch data list events Cancelled", responseData)
}
func (h *HandlerEvent) GetMyEventByID(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 0, 64)
	if err != nil {
		response.Error(c, http.StatusBadRequest, "Invalid Id Request", err.Error())
		return
	}
	userId := c.MustGet("user_id").(uint64)
	event, err := h.service.GetMyEventById(c.Request.Context(), userId, id)
	if err != nil {
		response.Error(c, http.StatusInternalServerError, "Failed Get Event", err.Error())
		return
	}
	response.Success(c, http.StatusOK, "Success Get Event ", event)
}
func (h *HandlerEvent) ListMyEvents(c *gin.Context) {
	var req RequestListEvent
	if err := c.ShouldBindQuery(&req); err != nil {
		response.Error(c, http.StatusBadRequest, "Invalid Id Request", err.Error())
		return
	}
	userId := c.MustGet("user_id").(uint64)
	events, total, err := h.service.ListMyEvents(c.Request.Context(), userId, req.Page, req.Limit)
	if err != nil {
		response.Error(c, http.StatusInternalServerError, "Failed Get List My Events", err.Error())
		return
	}
	responseData := gin.H{
		"events": events,
		"total":  total,
		"page":   req.Page,
		"limit":  req.Limit,
	}
	response.Success(c, http.StatusOK, "Success Fetch data list My Events ", responseData)
}
