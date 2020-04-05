package api

import (
    "net/http"
    "github.com/gin-gonic/gin"

    "drylm.org/ricardo/user_creation/entities"
)

var route = "/user"

func SetupUserController(engine *gin.Engine) {
    engine.GET(route, getUser)
    engine.POST(route, postUser)
}

func getUser(c *gin.Context) {
    user := entities.User{FirstName: "john",Email: "muller.john@gmail.com"}

    c.JSON(http.StatusOK, gin.H{"user": user})
}

func postUser(c *gin.Context) {
    c.JSON(http.StatusOK, gin.H{})
}