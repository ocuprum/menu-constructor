package services

import (
	"context"

	"github.com/google/uuid"
	
	"github.com/ocuprum/menu-constructor/internal/models"
	"github.com/ocuprum/menu-constructor/internal/repositories"
)

type IngredientService struct {
	rep repositories.IngredientRepository
}

func NewIngredientService(rep repositories.IngredientRepository) *IngredientService {
	return &IngredientService{rep: rep}
}

func (s *IngredientService) GetByID(ctx context.Context, id uuid.UUID) (models.Ingredient, error) {
	return s.rep.GetByID(ctx, id)
}

func (s *IngredientService) Paginate(ctx context.Context, limit, offset int) ([]models.Ingredient, error) {
	return s.rep.Paginate(ctx, limit, offset)
}

func (s *IngredientService) Create(ctx context.Context, ingred models.Ingredient) error {
	return s.rep.Create(ctx, ingred)
}

func (s *IngredientService) Change(ctx context.Context, ingred models.Ingredient) error {
	return s.rep.Change(ctx, ingred)
}

func (s *IngredientService) Delete(ctx context.Context, ids []uuid.UUID) error {
	return s.rep.Delete(ctx, ids)
}