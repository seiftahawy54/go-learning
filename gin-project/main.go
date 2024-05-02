package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	ginInstance := gin.Default()

	ginInstance.GET("/events", getAllEvents)

	ginInstance.Run(":1234")
}

func getAllEvents(context * gin.Context) {
		context.JSON(http.StatusOK, gin.H{
			"welcomeMessage": "Welcome from my first golang application",
		})
}

