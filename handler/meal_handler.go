package handler

import (
  "net/http"

  "github.com/gorilla/mux"
  "github.com/jinzhu/gorm"
  _ "github.com/jinzhu/gorm/dialects/postgres"
  "github.com/atmavichara/apicurious_go/models"
)

func GetAllMeals(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
  meals := []models.Meal{}
  db.Find(&meals)
  RespondJSON(w, http.StatusOK, meals)
}

func GetMeal(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
  meal := models.Meal{}
  params := mux.Vars(r)

  if err := db.First(&meal, params["id"]).Error; err != nil {
    RespondError(w, http.StatusNotFound, err.Error())
  } else {
    RespondJSON(w, http.StatusOK, meal)
  }
}
