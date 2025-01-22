package pgsql

import (
	"context"

	"gorm.io/gorm"

	"github.com/ocuprum/menu-constructor/internal/models"
)

type DayRepository struct {
	conn *gorm.DB
}

func NewDayRepository(conn *gorm.DB) *DayRepository {
	return &DayRepository{conn: conn}
}

func (r *DayRepository) AddMeal(ctx context.Context, md models.MealDay) error {
	return r.conn.WithContext(ctx).Create(&md).Error
}

func (r *DayRepository) DeleteMeal(ctx context.Context, md models.MealDay) error {
	return r.conn.WithContext(ctx).Delete(&md).Error
}