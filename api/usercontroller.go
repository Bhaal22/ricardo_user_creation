package api

import (
    "net/http"
    "github.com/gin-gonic/gin"
    "github.com/jinzhu/gorm"

    "github.com/Bhaal22/ricardo_user_creation/entities"
    "github.com/Bhaal22/ricardo_user_creation/events"
    "github.com/Bhaal22/ricardo_user_creation/logic"
)

var route = "/user"

func SetupUserController(engine *gin.Engine) {
    engine.GET(route, getUsers)
    engine.GET(route+"/:id", getUser)
    engine.POST(route, postUser)
    engine.PATCH(route+"/:id", patchUser)
}

func getUsers(c *gin.Context) {
    db := c.MustGet("db").(*gorm.DB)

    var users []entities.User
    db.Find(&users)

    country, _ := logic.Country(c.ClientIP())
    c.JSON(http.StatusOK, gin.H{"data": logic.ValidCountry(country)})
}

func getUser(c *gin.Context) {
    db := c.MustGet("db").(*gorm.DB)

    // Get model if exist
    var user entities.User
    if err := db.Where("id = ?", c.Param("id")).First(&user).Error; err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "User not found!"})
        return
    }

    c.JSON(http.StatusOK, gin.H{"data": user})
}

type CreateUserInput struct {
    FirstName string `json:"first_name" binding:"required"`
    Email     string `json:"email" binding:"required"`
    Password  string `json:"password" binding:"required"`
}

func postUser(c *gin.Context) {
    db := c.MustGet("db").(*gorm.DB)
    rmq := c.MustGet("rmq").(*events.RMQ)

    event := &events.Event{Name: "CREATE_USER", Status: "UNKNOWN", Reason: ""}
    defer rmq.Publish(event)

    var input CreateUserInput
    if err := c.ShouldBindJSON(&input); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        event.Status = "FAILED"
        event.Reason = err.Error()

        return
    }

    user := entities.User{FirstName: input.FirstName, Email: input.Email, Password: input.Password}
    if err := db.Create(&user).Error; err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        event.Status = "FAILED"
        event.Reason = err.Error()

        return
    }

    c.JSON(http.StatusOK, gin.H{"data": user})

    event.Status = "SUCCESS"
    event.Reason = ""
}

type UpdateUserInput struct {
    FirstName string `json:"first_name"`
    Email     string `json:"email"`
    Password  string `json:"password"`
}

func patchUser(c *gin.Context) {
    db := c.MustGet("db").(*gorm.DB)
    rmq := c.MustGet("rmq").(*events.RMQ)

    event := &events.Event{Name: "UPDATE_USER", Status: "UNKNOWN", Reason: ""}
    defer rmq.Publish(event)

    var user entities.User
    if err := db.Where("id = ?", c.Param("id")).First(&user).Error; err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})

        event.Status = "FAILED"
        event.Reason = err.Error()
        return
    }

  
    var input UpdateUserInput
    if err := c.ShouldBindJSON(&input); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})

        event.Status = "FAILED"
        event.Reason = err.Error()
        return
    }

    db.Model(&user).Updates(input)


    c.JSON(http.StatusOK, gin.H{"data": user})

    event.Status = "SUCCESS"
    event.Reason = ""
}
