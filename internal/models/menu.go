package models 

import (
	"time"

	"github.com/google/uuid"
)

type Ingredient struct {
	ID     uuid.UUID
	Name   string
}

func NewIngredient(name string, amount float64) Ingredient {
	return Ingredient{
		ID:     uuid.New(),
		Name:   name,
	}
}

type Cart struct {
	Ingredients []Ingredient
}

type Nutrients struct {
	Proteins float64
	Fats     float64
	Carbs    float64
	Calories float64
}

func (n Nutrients) Add(addition Nutrients) Nutrients {
	n.Proteins += addition.Proteins
	n.Fats += addition.Fats
	n.Carbs += addition.Carbs
	n.Calories += addition.Calories

	return n
}

type Nutritional interface {
	CalcNutrients() Nutrients
}

type CookDuration interface {
	CalcCookingDuration() time.Duration
}

type Food struct {
	ID              uuid.UUID
	Name            string
	Ingredients     []Ingredient
	Nutrients
	CookingDuration time.Duration
}   

func (f *Food) CalcNutrients() Nutrients {
	return f.Nutrients
}

func (f *Food) CalcCookingDuration() time.Duration {
	return f.CookingDuration
}

func NewFood(name string, ingredients []Ingredient, 
	         nutrients Nutrients, cookingDuration time.Duration) Food {
	return Food{
		ID:              uuid.New(),
		Name:            name,
		Ingredients:     ingredients,
		Nutrients:       nutrients,
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
}

func (m *Meal) CalcNutrients() Nutrients {
	n := Nutrients{}
	for _, food := range m.Foods {
		n = n.Add(food.CalcNutrients())
	}
	
	return n
}

func (m *Meal) CalcCookingDuration() time.Duration {
	var cd time.Duration
	for _, food := range m.Foods {
		cd += food.CalcCookingDuration()
	}

	return cd
}

func NewMeal(name string, foods []Food, ingredients []Ingredient) Meal {
	return Meal{
		ID:              uuid.New(),
		Name:            name,
		Foods:           foods,
		Ingredients:     ingredients,
	}	
}

type Day struct {
	ID          uuid.UUID
	Date        time.Time
	Meals       []Meal
	Ingredients []Ingredient
}

func NewDay(date time.Time, meals []Meal, ingredients []Ingredient) Day {
	return Day{
		ID:              uuid.New(),
		Date:            date,
		Meals:           meals,
		Ingredients:     ingredients,
	}
}

func (d *Day) CalcNutrients() Nutrients {
	n := Nutrients{}
	for _, meal := range d.Meals {
		n = n.Add(meal.CalcNutrients())
	}

	return n
}

func (d *Day) CalcCookingDuration() time.Duration {
	var cd time.Duration
	for _, meal := range d.Meals {
		cd += meal.CalcCookingDuration()
	}

	return cd
}

type IngredientFood struct {
	IngredientID uuid.UUID
	FoodID       uuid.UUID
}

type FoodCategory struct {
	FoodID     uuid.UUID
	CategoryID uuid.UUID
}

type FoodMeal struct {
	FoodID uuid.UUID
	MealID uuid.UUID
}

type MealDay struct {
	MealID uuid.UUID
	DayID  uuid.UUID
}