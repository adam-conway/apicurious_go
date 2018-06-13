package handler

import (
  "net/http"

  "github.com/gorilla/mux"
  "github.com/jinzhu/gorm"
  _ "github.com/jinzhu/gorm/dialects/postgres"
  "github.com/atmavichara/apicurious_go/models"
)

func GetAllFoods(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
  foods := []models.Food{}
  db.Find(&foods)
  RespondJSON(w, http.StatusOK, foods)
}

func GetFood(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
  food := models.Food{}
  params := mux.Vars(r)

  if err := db.First(&food, params["id"]).Error; err != nil {
    RespondError(w, http.StatusNotFound, err.Error())
  } else {
    RespondJSON(w, http.StatusOK, food)
  }
}
