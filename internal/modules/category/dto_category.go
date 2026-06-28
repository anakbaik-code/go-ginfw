package category

import "time"

type CreateCategoryRequest struct {
	Name string `json:"name" binding:"required,min=3,max=100"`
	Slug string `json:"slug" binding:"required,min=3,max=100"`
}

type UpdateCategoryRequest struct {
	Name string `json:"name" binding:"required,min=3,max=100"`
	Slug string `json:"slug" binding:"required,min=3,max=100"`
}

type CategoryResponse struct {
	ID        uint32    `json:"id"`
	Name      string    `json:"name"`
	Slug      string    `json:"slug"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type ListCategoryRequest struct {
	Page  int32 `form:"page" binding:"gte=1"`
	Limit int32 `form:"limit" binding:"gte=1"`
}

type MetaResponse struct {
	Page  int   `json:"page"`
	Limit int   `json:"limit"`
	Total int64 `json:"total"`
}
type ListCategoriesResponse struct {
	Category []CategoryResponse `json:"data"`
	Meta     MetaResponse       `json:"meta"`
}
