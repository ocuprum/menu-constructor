package models 

import (
	"time"

	"github.com/google/uuid"
)

type Ingredient struct {
	ID     uuid.UUID `json:"id"`
	Name   string    `json:"name"`
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
	Proteins float64 `json:"proteins"`
	Fats     float64 `json:"fats"`
	Carbs    float64 `json:"carbs"`
	Calories float64 `json:"calories"`
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
	ID              uuid.UUID     `json:"id"`
	Name            string        `json:"name"`
	Ingredients     []Ingredient  `json:"ingredients" gorm:"-"`
	Nutrients
	CookingDuration time.Duration `json:"cooking_duration"`
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
	ID    uuid.UUID `json:"id"`
	Name  string    `json:"name"`
	Foods []Food    `json:"foods" gorm:"-"`
}

func NewCategory(name string, foods []Food) Category {
	return Category{
		ID:    uuid.New(),
		Name:  name,
		Foods: foods,
	}
}

type Meal struct {
	ID          uuid.UUID `json:"id"`
	Name        string    `json:"name"`
	Foods       []Food    `json:"foods" gorm:"-"`
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

func NewMeal(name string, foods []Food) Meal {
	return Meal{
		ID:              uuid.New(),
		Name:            name,
		Foods:           foods,
	}	
}

type Day struct {
	ID          uuid.UUID `json:"id"`
	Date        time.Time `json:"date"`
	Meals       []Meal    `json:"meals" gorm:"-"`
}

func NewDay(date time.Time, meals []Meal) Day {
	return Day{
		ID:              uuid.New(),
		Date:            date,
		Meals:           meals,
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
	IngredientID uuid.UUID `json:"ingredient_id"`
	FoodID       uuid.UUID `json:"food_id"`
}

func (IngredientFood) TableName() string {
	return "ingredient_food"
}

type FoodCategory struct {
	FoodID     uuid.UUID `json:"food_id"`
	CategoryID uuid.UUID `json:"category_id"`
}

func (FoodCategory) TableName() string {
	return "food_category"
}

type FoodMeal struct {
	FoodID uuid.UUID `json:"food_id"`
	MealID uuid.UUID `json:"meal_id"`
}

func (FoodMeal) TableName() string {
	return "food_meal"
}

type MealDay struct {
	MealID uuid.UUID `json:"meal_id"`
	DayID  uuid.UUID `json:"day_id"`
}

func (MealDay) TableName() string {
	return "meal_day"
}