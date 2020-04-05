package main

import (
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"

	"github.com/ricardo_user_creation/api"
	"github.com/ricardo_user_creation/entities"
)

type server struct {
	engine *gin.Engine
	db     *gorm.DB
}

func main() {
	s := server{engine: gin.Default(), db: entities.Setup()}

	s.engine.Use(func(c *gin.Context) {
		c.Set("db", s.db)
		c.Next()
	})

	api.SetupUserController(s.engine)

	s.engine.Run()
}
