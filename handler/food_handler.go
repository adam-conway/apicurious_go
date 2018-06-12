package handler

import (
  "net/http"

  "github.com/jinzhu/gorm"
  _ "github.com/jinzhu/gorm/dialects/postgres"
  "github.com/atmavichara/apicurious_go/models"
)

func GetAllFoods(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
  foods := []models.Food{}
  db.Find(&foods)
  respondJSON(w, http.StatusOK, r, foods)
}
