package pgsql

import (
	"context"

	"gorm.io/gorm"

	"github.com/google/uuid"
	"github.com/ocuprum/menu-constructor/internal/models"
)

type CategoryRepository struct {
	conn *gorm.DB
}

func NewCategoryRepository(conn *gorm.DB) *CategoryRepository {
	return &CategoryRepository{conn: conn}
}

func (r *CategoryRepository) GetByID(ctx context.Context, id uuid.UUID) (category models.Category, err error) {
	err = r.conn.WithContext(ctx).Where("id = ?", id).First(&category).Error
	if err != nil {
		return models.Category{}, err
	}

	return category, nil
}

func (r *CategoryRepository) Paginate(ctx context.Context, limit, offset int) (categories []models.Category, err error) {
	err = r.conn.WithContext(ctx).Limit(limit).Offset(offset).Find(&categories).Error
	if err != nil {
		return []models.Category{}, err
	}
	
	return categories, nil
}

func (r *CategoryRepository) Create(ctx context.Context, category models.Category) error {
	return r.conn.WithContext(ctx).Create(&category).Error
}

func (r *CategoryRepository) Change(ctx context.Context, category models.Category) error {
	return r.conn.WithContext(ctx).Save(&category).Error
}

func (r *CategoryRepository) Delete(ctx context.Context, ids []uuid.UUID) error {
	if len(ids) == 0 {
		return nil
	}

	return r.conn.WithContext(ctx).Where("id IN (?)", ids).Delete(&models.Category{}).Error
}

func (r *CategoryRepository) AddFood(ctx context.Context, fc models.FoodCategory) error {
	return r.conn.WithContext(ctx).Create(&fc).Error
}

func (r *CategoryRepository) DeleteFood(ctx context.Context, fc models.FoodCategory) error {
	return r.conn.WithContext(ctx).
			Where("category_id = ? AND food_id = ?", fc.CategoryID, fc.FoodID).
			Delete(&models.FoodCategory{}).Error
}