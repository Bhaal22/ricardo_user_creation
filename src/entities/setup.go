package entities

import (
    "github.com/jinzhu/gorm"
    _ "github.com/jinzhu/gorm/dialects/sqlite"
)

func Setup() *gorm.DB {
    db, err := gorm.Open("sqlite3", "test.db")
  
    if err != nil {
        panic("Failed to connect to database!")
    }
  
    db.AutoMigrate(&User{})
  
    return db
}