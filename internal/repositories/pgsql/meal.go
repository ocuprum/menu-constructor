package pgsql

import (
	"context"

	"gorm.io/gorm"

	"github.com/google/uuid"
	"github.com/ocuprum/menu-constructor/internal/models"
)

type MealRepository struct {
	conn *gorm.DB
}

func NewMealRepository(conn *gorm.DB) *MealRepository {
	return &MealRepository{conn: conn}
}

func (r *MealRepository) GetByID(ctx context.Context, id uuid.UUID) (meal models.Meal, err error) {
	err = r.conn.WithContext(ctx).Where("id = ?", id).First(&meal).Error
	if err != nil {
		return models.Meal{}, err
	}

	return meal, nil
}

func (r *MealRepository) Paginate(ctx context.Context, limit, offset int) (meals []models.Meal, err error) {
	err = r.conn.WithContext(ctx).Limit(limit).Offset(offset).Find(&meals).Error
	if err != nil {
		return []models.Meal{}, err
	}
	
	return meals, nil
}

func (r *MealRepository) Create(ctx context.Context, meal models.Meal) error {
	return r.conn.WithContext(ctx).Create(&meal).Error
}

func (r *MealRepository) Change(ctx context.Context, meal models.Meal) error {
	return r.conn.WithContext(ctx).Save(&meal).Error
}

func (r *MealRepository) Delete(ctx context.Context, ids []uuid.UUID) error {
	if len(ids) == 0 {
		return nil
	}

	return r.conn.WithContext(ctx).Where("id IN (?)", ids).Delete(&models.Meal{}).Error
}

func (r *MealRepository) AddFood(ctx context.Context, fm models.FoodMeal) error {
	return r.conn.WithContext(ctx).Create(&fm).Error
}

func (r *MealRepository) DeleteFood(ctx context.Context, fm models.FoodMeal) error {
	return r.conn.WithContext(ctx).Delete(&fm).Error
}