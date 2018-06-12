package db

import (
  "fmt"
  "log"

  "github.com/jinzhu/gorm"
  _ "github.com/jinzhu/gorm/dialects/postgres"
  "github.com/atmavichara/apicurious_go/models"
  "github.com/atmavichara/apicurious_go/config"
)

type DB struct {
  DB *gorm.DB
}

func Init() {
  dbParams = fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s",
    config.DB.Host,
    config.DB.Port,
    config.DB.User,
    config.DB.DBName,
    config.DB.Password)
  var err error
  db, err = gorm.Open("postgres", dbParams)
  if err != nil {
    log.Fatal("Error in connecting to database")
  }

  db.AutoMigrate(&models.Food{})
  db.AutoMigrate(&models.Meal{})
}

func (db *DB) CloseDb() {
  db.DB.Close()
}
