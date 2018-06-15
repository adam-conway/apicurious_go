package models

import (
  "github.com/jinzhu/gorm"
  _ "github.com/jinzhu/gorm/dialects/postgres"
)

type MealFood struct {
  gorm.Model

  Meal Meal
  MealID uint
  Food Food
  FoodID uint
}
