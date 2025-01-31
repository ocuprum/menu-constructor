package services

import (
	"context" 

	"github.com/google/uuid"

	"github.com/ocuprum/menu-constructor/internal/models"
	"github.com/ocuprum/menu-constructor/internal/repositories"
)

type MealService struct {
	rep repositories.MealRepository
}

func NewMealService(rep repositories.MealRepository) *MealService {
	return &MealService{rep: rep}
}

func (s *MealService) GetByID(ctx context.Context, id uuid.UUID) (models.Meal, error) {
	return s.rep.GetByID(ctx, id)
}

func (s *MealService) Paginate(ctx context.Context, limit, offset int) ([]models.Meal, error) {
	return s.rep.Paginate(ctx, limit, offset)
}

func (s *MealService) Create(ctx context.Context, category models.Meal) error {
	return s.rep.Create(ctx, category)
}

func (s *MealService) Change(ctx context.Context, category models.Meal) error {
	return s.rep.Change(ctx, category)
}

func (s *MealService) Delete(ctx context.Context, ids []uuid.UUID) error {
	return s.rep.Delete(ctx, ids)
}

func (s *MealService) AddFood(ctx context.Context, fm models.FoodMeal) error {
	return s.rep.AddFood(ctx, fm)
}

func (s *MealService) DeleteFood(ctx context.Context, fm models.FoodMeal) error {
	return s.rep.DeleteFood(ctx, fm)
}