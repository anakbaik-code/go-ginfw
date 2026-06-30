package category

import (
	"context"
	"errors"
)

type CategoryService interface {
	CreateCategory(ctx context.Context, req CreateCategoryRequest) (*CategoryResponse, error)
	DeleteCategory(ctx context.Context, id uint32) error
	GetById(ctx context.Context, id uint32) (*CategoryResponse, error)
	List(ctx context.Context, page int32, limit int32) ([]CategoryResponse, int64, error)
	Update(ctx context.Context, cat Category) error
	GetByName(ctx context.Context, name string) (*CategoryResponse, error)
}
type categoryService struct {
	repo CategoryRepository
}

func NewServiceCategory(repo CategoryRepository) CategoryService {
	return &categoryService{
		repo: repo,
	}
}

func (s *categoryService) CreateCategory(ctx context.Context, req CreateCategoryRequest) (*CategoryResponse, error) {
	exist, err := s.repo.GetByName(ctx, req.Name)
	if err == nil && exist != nil {
		return nil, errors.New("category already exists")
	}
	category := Category{
		Name: req.Name,
		Slug: req.Slug,
	}
	id, err := s.repo.Create(ctx, category)
	if err != nil {
		return nil, err
	}

	categoryResponse := &CategoryResponse{
		ID:   id,
		Name: req.Name,
		Slug: req.Slug,
	}

	return categoryResponse, nil
}

func (s *categoryService) DeleteCategory(ctx context.Context, id uint32) error {
	if err := s.repo.Delete(ctx, id); err != nil {
		return err
	}
	return nil
}

func (s *categoryService) GetById(ctx context.Context, id uint32) (*CategoryResponse, error) {
	row, err := s.repo.GetById(ctx, id)
	if err != nil {
		return nil, err
	}
	categoryResponse := &CategoryResponse{
		ID:        row.Id,
		Name:      row.Name,
		Slug:      row.Slug,
		CreatedAt: row.CreatedAt,
		UpdatedAt: row.UpdatedAt,
	}
	return categoryResponse, nil
}

func (s *categoryService) List(ctx context.Context, page int32, limit int32) ([]CategoryResponse, int64, error) {
	if page <= 0 {
		page = 1
	}

	if limit <= 0 {
		limit = 10
	}

	offset := (page - 1) * limit
	categories, err := s.repo.List(ctx, limit, offset)
	if err != nil {
		return nil, 0, err
	}

	categoryResponse := make([]CategoryResponse, 0, len(categories))
	for _, category := range categories {
		categoryResponse = append(categoryResponse, CategoryResponse{
			ID:        category.Id,
			Name:      category.Name,
			Slug:      category.Slug,
			CreatedAt: category.CreatedAt,
			UpdatedAt: category.UpdatedAt,
		})
	}

	total, err := s.repo.CountCategory(ctx)
	if err != nil {
		return nil, 0, err
	}
	return categoryResponse, total, nil
}

func (s *categoryService) Update(ctx context.Context, cat Category) error {
	if err := s.repo.Update(ctx, cat); err != nil {
		return err
	}
	return nil
}
func (s *categoryService) GetByName(ctx context.Context, name string) (*CategoryResponse, error) {
	row, err := s.repo.GetByName(ctx, name)
	if err != nil {
		return nil, err
	}
	category := &CategoryResponse{
		ID:        row.Id,
		Name:      row.Name,
		Slug:      row.Slug,
		CreatedAt: row.CreatedAt,
		UpdatedAt: row.UpdatedAt,
	}

	return category, nil
}
