package main

import (
	"github.com/gin-gonic/gin"
)

// EggsCount - Глобальный счётчик яиц
var EggsCount = 0

// SetupRouter - получения gin-роутера
func SetupRouter() *gin.Engine {
	router := gin.Default()
	router.GET("/eggs", getEggsCount)
	return router
}

// main - точка входа
func main() {
	router := SetupRouter()
	config := InitConfig()

	go startChickenTask(config)
	go startFarmerTask(config)

	router.Run(":8000")
}
