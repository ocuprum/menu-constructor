package services

import (
	"context" 

	"github.com/ocuprum/menu-constructor/internal/models"
	"github.com/ocuprum/menu-constructor/internal/repositories"
)

type DayService struct {
	rep repositories.DayRepository
}

func NewDayService(rep repositories.DayRepository) *DayService {
	return &DayService{rep: rep}
}

func (s *DayService) AddMeal(ctx context.Context, md models.MealDay) error {
	return s.rep.AddMeal(ctx, md)
}

func (s *DayService) DeleteMeal(ctx context.Context, md models.MealDay) error {
	return s.rep.DeleteMeal(ctx, md)
}