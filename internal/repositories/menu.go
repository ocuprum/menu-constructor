package repositories

import (
	"context"

	"github.com/google/uuid"
	"github.com/ocuprum/menu-constructor/internal/models"
)

type IngredientRepository interface {
	GetByID(ctx context.Context, id uuid.UUID) (models.Ingredient, error)
	Paginate(ctx context.Context, limit, offset int) ([]models.Ingredient, error)
	Create(ctx context.Context, ingred models.Ingredient) error
	Change(ctx context.Context, ingred models.Ingredient) error
	Delete(ctx context.Context, ids []uuid.UUID) error
}

type FoodRepository interface {
	GetByID(ctx context.Context, id uuid.UUID) (models.Food, error)
	Paginate(ctx context.Context, limit, offset int) ([]models.Food, error)
	Create(ctx context.Context, food models.Food) error
	Change(ctx context.Context, food models.Food) error
	Delete(ctx context.Context, ids []uuid.UUID) error
	AddIngrediet(ctx context.Context, fi models.IngredientFood) error
	DeleteIngredient(ctx context.Context, fi models.IngredientFood) error
}

type CategoryRepository interface {
	GetByID(ctx context.Context, id uuid.UUID) (models.Category, error)
	Paginate(ctx context.Context, limit, offset int) ([]models.Category, error)
	Create(ctx context.Context, category models.Category) error
	Change(ctx context.Context, category models.Category) error
	Delete(ctx context.Context, ids []uuid.UUID) error
	AddFood(ctx context.Context, fc models.FoodCategory) error
	DeleteFood(ctx context.Context, fc models.FoodCategory) error
}

type MealRepository interface {
	GetByID(ctx context.Context, id uuid.UUID) (models.Meal, error)
	Paginate(ctx context.Context, limit, offset int) ([]models.Meal, error)
	Create(ctx context.Context, category models.Meal) error
	Change(ctx context.Context, id uuid.UUID, category models.Meal) error
	Delete(ctx context.Context, ids []uuid.UUID) error
	AddFood(ctx context.Context, fm models.FoodMeal) error
	DeleteFood(ctx context.Context, fm models.FoodMeal) error
}

type DayRepository interface {
	AddMeal(ctx context.Context, md models.MealDay) error
	DeleteMeal(ctx context.Context, md models.MealDay) error
}