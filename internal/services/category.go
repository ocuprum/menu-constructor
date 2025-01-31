package services

import (
	"context" 

	"github.com/google/uuid"

	"github.com/ocuprum/menu-constructor/internal/models"
	"github.com/ocuprum/menu-constructor/internal/repositories"
)

type CategoryService struct {
	rep repositories.CategoryRepository
}

func NewCategoryService(rep repositories.CategoryRepository) *CategoryService {
	return &CategoryService{rep: rep}
}

func (s *CategoryService) GetByID(ctx context.Context, id uuid.UUID) (models.Category, error) {
	return s.rep.GetByID(ctx, id)
}

func (s *CategoryService) Paginate(ctx context.Context, limit, offset int) ([]models.Category, error) {
	return s.rep.Paginate(ctx, limit, offset)
}

func (s *CategoryService) Create(ctx context.Context, category models.Category) error {
	return s.rep.Create(ctx, category)
}

func (s *CategoryService) Change(ctx context.Context, category models.Category) error {
	return s.rep.Change(ctx, category)
}

func (s *CategoryService) Delete(ctx context.Context, ids []uuid.UUID) error {
	return s.rep.Delete(ctx, ids)
}

func (s *CategoryService) AddFood(ctx context.Context, fc models.FoodCategory) error {
	return s.rep.AddFood(ctx, fc)
}

func (s *CategoryService) DeleteFood(ctx context.Context, fc models.FoodCategory) error {
	return s.rep.DeleteFood(ctx, fc)
}