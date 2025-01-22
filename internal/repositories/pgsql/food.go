package pgsql

import (
	"context"

	"gorm.io/gorm"

	"github.com/google/uuid"
	"github.com/ocuprum/menu-constructor/internal/models"
)

type FoodRepository struct {
	conn *gorm.DB
}

func NewFoodRepository(conn *gorm.DB) *FoodRepository {
	return &FoodRepository{conn: conn}
}

func (r *FoodRepository) GetByID(ctx context.Context, id uuid.UUID) (food models.Food, err error) {
	err = r.conn.WithContext(ctx).Where("id = ?", id).First(&food).Error
	if err != nil {
		return models.Food{}, err
	}

	return food, nil
}

func (r *FoodRepository) Paginate(ctx context.Context, limit, offset int) (foods []models.Food, err error) {
	err = r.conn.WithContext(ctx).Limit(limit).Offset(offset).Find(&foods).Error
	if err != nil {
		return []models.Food{}, err
	}
	
	return foods, nil
}

func (r *FoodRepository) Create(ctx context.Context, food models.Food) error {
	return r.conn.WithContext(ctx).Create(&food).Error
}

func (r *FoodRepository) Change(ctx context.Context, food models.Food) error {
	return r.conn.WithContext(ctx).Save(&food).Error
}

func (r *FoodRepository) Delete(ctx context.Context, ids []uuid.UUID) error {
	if len(ids) == 0 {
		return nil
	}

	return r.conn.WithContext(ctx).Where("id IN (?)", ids).Delete(&models.Food{}).Error
}

func (r *FoodRepository) AddIngredient(ctx context.Context, fi models.IngredientFood) error {
	return r.conn.WithContext(ctx).Create(&fi).Error
}

func (r *FoodRepository) DeleteIngredient(ctx context.Context, fi models.IngredientFood) error {
	return r.conn.WithContext(ctx).Delete(&fi).Error
}
