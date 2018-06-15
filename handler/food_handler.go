package handler

import (
  "net/http"
  "encoding/json"
  "fmt"

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

func CreateFood(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
  food := models.Food{}
  decoder := json.NewDecoder(r.Body)
  fmt.Println(decoder)
  if err := decoder.Decode(&food); err != nil {
    RespondError(w, http.StatusBadRequest, err.Error())
    return
  }

  defer r.Body.Close()
  db.NewRecord(food)
  if err := db.Create(&food).Error; err != nil {
    RespondError(w, http.StatusInternalServerError, err.Error())
    return
  } else {
    RespondJSON(w, http.StatusCreated, food)
  }
}
