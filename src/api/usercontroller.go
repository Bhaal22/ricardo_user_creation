package api

import (
	"net/http"
    "github.com/gin-gonic/gin"
)

func Setup(engine *gin.Engine) {
	engine.GET("/user", func(c *gin.Context) {
        c.JSON(http.StatusOK, gin.H{"data": "hello world"})
    })
}