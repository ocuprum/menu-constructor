package models 

import (
	"time"

	"github.com/google/uuid"
)

type Ingredient struct {
	ID     uuid.UUID
	Name   string
	Amount float64
}

func NewIngredient(name string, amount float64) Ingredient {
	return Ingredient{
		ID:     uuid.New(),
		Name:   name,
		Amount: amount,
	}
}

type Nutrients struct {
	Proteins float64
	Fats     float64
	Carbs    float64
}

type Food struct {
	ID              uuid.UUID
	Name            string
	Ingredients     []Ingredient
	Nutrients
	Calories        float64
	CookingDuration time.Duration
}   

func NewFood(name string, ingredients []Ingredient, nutrients Nutrients, 
			 calories float64, cookingDuration time.Duration) Food {
	return Food{
		ID:              uuid.New(),
		Name:            name,
		Ingredients:     ingredients,
		Nutrients:       nutrients,
		Calories:        calories,
		CookingDuration: cookingDuration,
	}
}

type Category struct {
	ID    uuid.UUID
	Name  string
	Foods []Food
}

func NewCategory(name string, foods []Food) Category {
	return Category{
		ID:    uuid.New(),
		Name:  name,
		Foods: foods,
	}
}

type Meal struct {
	ID          uuid.UUID
	Name        string
	Foods       []Food
	Ingredients []Ingredient
	Nutrients
	Calories    float64
	CookingDuration time.Duration
}

func NewMeal(name string, foods []Food, ingredients []Ingredient,
             nutrients Nutrients, calories float64, cookingDuration time.Duration) Meal {
	return Meal{
		ID:              uuid.New(),
		Name:            name,
		Foods:           foods,
		Ingredients:     ingredients,
		Nutrients:       nutrients,
		Calories:        calories,
		CookingDuration: cookingDuration,
	}	
}

type Day struct {
	ID          uuid.UUID
	Date        time.Time
	Meals       []Meal
	Ingredients []Ingredient
	Nutrients
	Calories    float64
	CookingDuration time.Duration
}

func NewDay(date time.Time, meals []Meal, ingredients []Ingredient,
			nutrients Nutrients, calories float64, cookingDuration time.Duration) Day {
	return Day{
		ID:              uuid.New(),
		Date:            date,
		Meals:           meals,
		Ingredients:     ingredients,
		Nutrients:       nutrients,
		Calories:        calories,
		CookingDuration: cookingDuration,
	}
}