package main

import (
    "github.com/gin-gonic/gin"
    "github.com/jinzhu/gorm"

    "github.com/Bhaal22/ricardo_user_creation/api"
    "github.com/Bhaal22/ricardo_user_creation/entities"
    "github.com/Bhaal22/ricardo_user_creation/events"
)

type server struct {
    engine *gin.Engine
    db     *gorm.DB
    rmq    *events.RMQ
}

func main() {
    s := server{engine: gin.Default(), db: entities.Setup(), rmq: &events.RMQ{ConnectionString: "amqp://guest:guest@rmq:5672/"}}

    s.engine.Use(func(c *gin.Context) {
        c.Set("db", s.db)
        c.Set("rmq", s.rmq)
        c.Next()
    })

    api.SetupUserController(s.engine)

    s.engine.Run()
}
