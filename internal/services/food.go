package services

import (
	"context" 

	"github.com/google/uuid"

	"github.com/ocuprum/menu-constructor/internal/models"
	"github.com/ocuprum/menu-constructor/internal/repositories"
)

type FoodService struct {
	rep  repositories.FoodRepository
}

func NewFoodService(rep repositories.FoodRepository) *FoodService {
	return &FoodService{rep: rep}
}

func (s *FoodService) GetByID(ctx context.Context, id uuid.UUID) (models.Food, error) {
	return s.rep.GetByID(ctx, id)
}

func (s *FoodService) Paginate(ctx context.Context, limit, offset int) ([]models.Food, error) {
	return s.rep.Paginate(ctx, limit, offset)
}

func (s *FoodService) Create(ctx context.Context, food models.Food) error {
	return s.rep.Create(ctx, food)
}

func (s *FoodService) Change(ctx context.Context, food models.Food) error {
	return s.rep.Change(ctx, food)
}

func (s *FoodService) Delete(ctx context.Context, ids []uuid.UUID) error {
	return s.rep.Delete(ctx, ids)
}

func (s *FoodService) AddIngredient(ctx context.Context, fi models.IngredientFood) error {
	return s.rep.AddIngredient(ctx, fi)
}

func (s *FoodService) DeleteIngredient(ctx context.Context, fi models.IngredientFood) error {
	return s.rep.DeleteIngredient(ctx, fi)
}
