package models

import (
  "github.com/jinzhu/gorm"
  _ "github.com/jinzhu/gorm/dialects/postgres"
)

type Food struct {
  gorm.Model

  Name string `gorm: "type:varchar(100)" json: "name"`
  Calories uint `json: "calories"`
  Meals []Meal `gorm: "many2many:meal_foods"`
}
