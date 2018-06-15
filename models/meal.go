package models

import (
  "github.com/jinzhu/gorm"
  _ "github.com/jinzhu/gorm/dialects/postgres"
)

type Meal struct {
  gorm.Model

  Name string `gorm: "type:varchar(100)" json: "name"`
  Foods []*Food `gorm:"many2many:meal_foods;"`
}
