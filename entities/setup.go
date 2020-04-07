package entities

import (
    "fmt"
    "github.com/jinzhu/gorm"
    _ "github.com/jinzhu/gorm/dialects/postgres"
)

func Setup() *gorm.DB {
    db, err := gorm.Open("postgres", "host=pgsql user=postgres dbname=ricardo password=ricardo sslmode=disable")
  
    if err != nil {
        panic(fmt.Sprintf("Failed to connect to database! %s", err))
    }
  
    db.AutoMigrate(&User{})
  
    return db
}