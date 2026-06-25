package category

import (
	"context"
	"errors"
)

type CategoryService interface {
}
type categoryService struct {
	repo CategoryRepository
}

func NewCategoryService(repo CategoryRepository) CategoryService {
	return &categoryService{
		repo: repo,
	}
}

func (s *categoryService) CreateCategory(ctx context.Context, req *CreateCategoryRequest) (*Category, error) {
	exist, err := s.repo.GetByName(ctx, req.Name)
	if err == nil && exist != nil {
		return nil, errors.New("category already exists")
	}
	category := &Category{
		Name: req.Name,
		Slug: req.Slug,
	}
	id, err := s.repo.Create(ctx, category)
	if err != nil {
		return nil, err
	}
	category.Id = id
	return category, nil
}

func (s *categoryService) DeleteCategory(ctx context.Context, id uint32) error {
	err := s.repo.Delete(ctx, id)
	return err
}

func (s *categoryService) GetById(ctx context.Context, id uint32) (*Category, error) {
	row, err := s.repo.GetById(ctx, id)
	if err != nil {
		return nil, err
	}
	category := &Category{
		Id:   row.Id,
		Name: row.Name,
		Slug: row.Slug,
	}
	return category, nil
}
