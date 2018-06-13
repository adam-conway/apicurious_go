package main

import (
  "net/http"
  "log"
  "fmt"

  "github.com/jinzhu/gorm"
  _ "github.com/jinzhu/gorm/dialects/postgres"
  "github.com/gorilla/mux"
  "github.com/atmavichara/apicurious_go/models"
  "github.com/atmavichara/apicurious_go/handler"
  "github.com/atmavichara/apicurious_go/config"
)

type App struct {
  DB *gorm.DB
}

func main() {
  config := config.GetConfig()
  a := &App{}
  a.Init(config)
  fmt.Println("Server Running on :3000")

  router := mux.NewRouter()
  sub := router.PathPrefix("/api/v1").Subrouter()
  sub.HandleFunc("/foods", a.GetAllFoods).Methods("GET")
  sub.HandleFunc("/meals", a.GetAllMeals).Methods("GET")
  sub.HandleFunc("/meals/{id}/foods", a.GetMeal).Methods("GET")
  sub.HandleFunc("/foods/{id}", a.GetFood).Methods("GET")
  log.Fatal(http.ListenAndServe(":3000", sub))
}

func (a *App) Init(config *config.Config) {
  dbParams := fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=%s",
    config.DB.Host,
    config.DB.Port,
    config.DB.User,
    config.DB.DBName,
    config.DB.Password,
    config.DB.SSLMode)

  var err error
  a.DB, err = gorm.Open("postgres", dbParams)
  if err != nil {
    fmt.Println("ERROR IN CONNECTING TO DATABASE\n")
    log.Fatal(err.Error())
  }

  fmt.Println("Connected to Database")

  a.DB.AutoMigrate(&models.Food{})
  a.DB.AutoMigrate(&models.Meal{})
}

func (a *App) GetAllFoods(w http.ResponseWriter, r *http.Request) {
  handler.GetAllFoods(a.DB, w, r)
}

func (a *App) GetAllMeals(w http.ResponseWriter, r *http.Request) {
  handler.GetAllMeals(a.DB, w, r)
}

func (a *App) GetMeal(w http.ResponseWriter, r *http.Request) {
  handler.GetMeal(a.DB, w, r)
}

func (a *App) GetFood(w http.ResponseWriter, r *http.Request) {
  handler.GetFood(a.DB, w, r)
}
