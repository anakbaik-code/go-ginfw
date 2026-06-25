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