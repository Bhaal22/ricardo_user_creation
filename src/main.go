package main

import (
    "github.com/gin-gonic/gin"

    "drylm.org/ricardo/user_creation/api"
)

type server struct {
    context *gin.Engine 
}

func main() {
    s := server{context: gin.Default()}

    api.SetupUserController(s.context)

    s.context.Run()
}
