package category

import (
	"go-fwgin/internal/response"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type HandlerCategory struct {
	service CategoryService
}

func NewHandlerCategory(service CategoryService) *HandlerCategory {
	return &HandlerCategory{service: service}
}

func (h *HandlerCategory) Create(c *gin.Context) {
	var req CreateCategoryRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.Error(c, http.StatusBadRequest, "invalid request payload", err.Error())
		return
	}

	category, err := h.service.CreateCategory(c.Request.Context(), req)
	if err != nil {
		response.Error(c, http.StatusInternalServerError, "Invalid Server Error", err.Error())
		return
	}

	response.Success(c, http.StatusOK, "Category created successfully", category)
}

// 2. GET BY ID
func (h *HandlerCategory) GetByID(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		response.Error(c, http.StatusBadRequest, "Id invalid request", err.Error())
		return
	}

	category, err := h.service.GetById(c.Request.Context(), uint32(id))
	if err != nil {
		response.Error(c, http.StatusInternalServerError, "Internal Server Error", err.Error())
		return
	}

	if category == nil {
		response.Error(c, http.StatusNotFound, "Category Not Found", nil)
		return
	}

	response.Success(c, http.StatusOK, "Success", category)
}

// 3. LIST WITH PAGINATION
func (h *HandlerCategory) ListCategory(c *gin.Context) {
	var req ListCategoryRequest
	if err := c.ShouldBindQuery(&req); err != nil {
		response.Error(c, http.StatusBadRequest, "invalid query parameter", err.Error())
		return
	}

	categories, total, err := h.service.List(c.Request.Context(), req.Page, req.Limit)
	if err != nil {
		response.Error(c, http.StatusInternalServerError, "Internal Server Error", err.Error())
		return
	}

	response.Success(c, http.StatusOK, "Categories retrieved successfully", ListCategoriesResponse{
		Category: categories,
		Meta: MetaResponse{
			Page:  int(req.Page),
			Limit: int(req.Limit),
			Total: total,
		},
	})

}

func (h *HandlerCategory) Update(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		response.Error(c, http.StatusBadRequest, "Invalid Id request", err.Error())
		return
	}
	var req UpdateCategoryRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request payload"})
		return
	}

	category := Category{
		Id:   uint32(id),
		Name: req.Name,
		Slug: req.Slug,
	}

	err = h.service.Update(c.Request.Context(), category)
	if err != nil {
		response.Error(c, http.StatusInternalServerError, "Internal Server Error", err.Error())
		return
	}
	response.Success(c, http.StatusOK, "Updated Successfully", category)
}

func (h *HandlerCategory) Delete(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		response.Error(c, http.StatusBadRequest, "Invalid Bad Request", err.Error())
		return
	}
	if err := h.service.DeleteCategory(c.Request.Context(), uint32(id)); err != nil {
		response.Error(c, http.StatusInternalServerError, "Internal Server Error", err.Error())
		return
	}
	response.Success(c, http.StatusOK, "deleted successfully", nil)
}
