package handler

import (
  "net/http"
  "strconv"
  "fmt"

  "github.com/gorilla/mux"
  "github.com/jinzhu/gorm"
  _ "github.com/jinzhu/gorm/dialects/postgres"
  "github.com/atmavichara/apicurious_go/models"
)

func GetAllMealFoods(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
  mealFoods := []models.MealFood{}
  db.Set("gorm:auto_preload", true).Find(&mealFoods)
  RespondJSON(w, http.StatusOK, mealFoods)
}

func CreateMealFood(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
  params := mux.Vars(r)
  meal := models.Meal{}
  food := models.Food{}

  fu64, err := strconv.ParseUint(params["id"], 10, 32)
    if err != nil {
        fmt.Println(err)
    }
  mu64, err := strconv.ParseUint(params["meal_id"], 10, 32)
    if err != nil {
        fmt.Println(err)
    }

  foodId := uint(fu64)
  mealId := uint(mu64)

  if err := db.First(&meal, params["meal_id"]).Error; err != nil {
    RespondError(w, http.StatusNotFound, err.Error())
  }
  if err := db.First(&food, params["id"]).Error; err != nil {
    RespondError(w, http.StatusNotFound, err.Error())
  }

  mealFood := &models.MealFood{MealID: mealId, FoodID: foodId}

  db.NewRecord(mealFood)
  if err := db.Create(&mealFood).Error; err != nil {
    RespondError(w, http.StatusInternalServerError, err.Error())
  } else {
    message := map[string]string{"message": "Successfully added " + food.Name + " to " + meal.Name}
    RespondJSON(w, http.StatusCreated, message)
  }
}

func DeleteMealFood(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
    params := mux.Vars(r)
    mealFood := models.MealFood{}
    meal := models.Meal{}
    food := models.Food{}

    if err := db.Where("meal_id = ? AND food_id = ?", params["meal_id"], params["id"]).First(&mealFood).Error; err != nil {
      RespondError(w, http.StatusNotFound, err.Error())
      return
    }
    if err := db.First(&meal, params["meal_id"]).Error; err != nil {
      RespondError(w, http.StatusNotFound, err.Error())
      return
    }
    if err := db.First(&food, params["id"]).Error; err != nil {
      RespondError(w, http.StatusNotFound, err.Error())
      return
    }

    if err := db.Delete(&mealFood).Error; err != nil {
      RespondError(w, http.StatusInternalServerError, err.Error())
      return
    }

    message := map[string]string{"message": "Successfully removed " + food.Name + " from " + meal.Name}
    RespondJSON(w, http.StatusOK, message)
}
