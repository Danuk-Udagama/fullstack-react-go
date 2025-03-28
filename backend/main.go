package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	router := gin.Default()

	router.GET("/ping", func(requestBody *gin.Context) {
		requestBody.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})
	router.Run("0.0.0.0:8080")
}
