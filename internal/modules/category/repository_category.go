package category

import (
	"context"
	"go-fwgin/internal/database"
)

type CategoryRepository interface {
	Create(ctx context.Context, c *Category) (uint32, error)
	Delete(ctx context.Context, id uint32) error
	GetById(ctx context.Context, id uint32) (*Category, error)
	List(ctx context.Context) ([]Category, error)
	Update(ctx context.Context, id uint32, c *Category) error
	GetByName(ctx context.Context, name string) (*Category, error)
}
type repositoryCategory struct {
	queries *database.Queries
}

func NewRepositoryCategory(q *database.Queries) CategoryRepository {
	return &repositoryCategory{
		queries: q,
	}
}

func (r *repositoryCategory) Create(ctx context.Context, c *Category) (uint32, error) {
	result, err := r.queries.CreateCategory(ctx, database.CreateCategoryParams{
		Name: c.Name,
		Slug: c.Slug,
	})
	if err != nil {
		return 0, err
	}
	id, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}
	return uint32(id), nil

}

func (r *repositoryCategory) Delete(ctx context.Context, id uint32) error {
	err := r.queries.DeleteCategory(ctx, id)
	return err
}

func (r *repositoryCategory) GetById(ctx context.Context, id uint32) (*Category, error) {
	result, err := r.queries.GetCategoryByID(ctx, id)
	if err != nil {
		return nil, err
	}
	category := &Category{
		Id:        result.ID,
		Name:      result.Name,
		Slug:      result.Slug,
		CreatedAt: result.CreatedAt,
		UpdatedAt: result.UpdatedAt,
	}
	return category, nil
}

func (r *repositoryCategory) List(ctx context.Context) ([]Category, error) {
	rows, err := r.queries.ListCategories(ctx)
	if err != nil {
		return nil, err
	}
	categories := make([]Category, 0, len(rows))
	for _, row := range rows {
		categories = append(categories, Category{
			Id:        row.ID,
			Name:      row.Name,
			Slug:      row.Slug,
			CreatedAt: row.CreatedAt,
			UpdatedAt: row.UpdatedAt,
		})
	}
	return categories, nil
}

func (r *repositoryCategory) Update(ctx context.Context, id uint32, c *Category) error {
	err := r.queries.UpdateCategory(ctx, database.UpdateCategoryParams{
		ID:   id,
		Name: c.Name,
		Slug: c.Slug,
	})
	return err
}
func (r *repositoryCategory) GetByName(ctx context.Context, name string) (*Category, error) {
	row, err := r.queries.GetCategoryByName(ctx, name)
	if err != nil {
		return nil, err
	}
	category := &Category{
		Id:        row.ID,
		Name:      row.Name,
		Slug:      row.Slug,
		CreatedAt: row.CreatedAt,
		UpdatedAt: row.UpdatedAt,
	}
	return category, nil
}
