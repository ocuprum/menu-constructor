package pgsql

import (
	"context"

	"gorm.io/gorm"

	"github.com/google/uuid"
	"github.com/ocuprum/menu-constructor/internal/models"
)

type IngredientRepository struct {
	conn *gorm.DB
}

func NewIngredientRepository(conn *gorm.DB) *IngredientRepository {
	return &IngredientRepository{conn: conn}
}

func (r *IngredientRepository) GetByID(ctx context.Context, id uuid.UUID) (ingred models.Ingredient, err error) {
	err = r.conn.WithContext(ctx).Where("id = ?", id).First(&ingred).Error
	if err != nil {
		return models.Ingredient{}, err
	}

	return ingred, nil
}

func (r *IngredientRepository) Paginate(ctx context.Context, limit, offset int) (ingreds []models.Ingredient, err error) {
	err = r.conn.WithContext(ctx).Limit(limit).Offset(offset).Find(&ingreds).Error
	if err != nil {
		return []models.Ingredient{}, err
	}
	
	return ingreds, nil
}

func (r *IngredientRepository) Create(ctx context.Context, ingred models.Ingredient) error {
	return r.conn.WithContext(ctx).Create(&ingred).Error
}

func (r *IngredientRepository) Change(ctx context.Context, ingred models.Ingredient) error {
	return r.conn.WithContext(ctx).Save(&ingred).Error
}

func (r *IngredientRepository) Delete(ctx context.Context, ids []uuid.UUID) error {
	if len(ids) == 0 {
		return nil
	}

	return r.conn.WithContext(ctx).Where("id IN (?)", ids).Delete(&models.Ingredient{}).Error
}