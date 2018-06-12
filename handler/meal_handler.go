package handler

import (
  "net/http"

  "github.com/jinzhu/gorm"
  _ "github.com/jinzhu/gorm/dialects/postgres"
  "github.com/atmavichara/apicurious_go/models"
)

func GetAllMeals(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
  meals := []models.Meal{}
  db.Find(&meals)
  respondJSON(w, http.StatusOK, r, meals)
}
