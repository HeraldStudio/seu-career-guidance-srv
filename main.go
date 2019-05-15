package main

import (
	"github.com/gin-gonic/gin"
	"go-common/database"
	"go-common/handlers"
)

const ServerPort = "3001"

func main() {

	database.Connect()
	engine := gin.Default()
	handlers.Handler(engine)
	engine.Run(":" + ServerPort)

}
