package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// getEggsCount - эндпоинт для получения текущего количества яиц (EggsCount)
func getEggsCount(c *gin.Context) {
	c.IndentedJSON(
		http.StatusOK,
		map[string]int{
			"eggsCount": EggsCount,
		},
	)
}
