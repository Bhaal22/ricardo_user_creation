package main

import (
    "github.com/gin-gonic/gin"
    "github.com/jinzhu/gorm"

    "drylm.org/ricardo/user_creation/api"
    "drylm.org/ricardo/user_creation/entities"
)

type server struct {
    engine *gin.Engine 
    db  *gorm.DB
}



func main() {
    s := server{engine: gin.Default(),db: entities.Setup()}   
    
    s.engine.Use(func(c *gin.Context) {
        c.Set("db", s.db)
        c.Next()
      })

    api.SetupUserController(s.engine)

    s.engine.Run()
}
