package main

import (
	"awesomeProject/api"
	"awesomeProject/db"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

func main() {
	router := gin.Default()

	client, err := db.ConnectToDB()
	if err != nil {
		log.Fatal("Failed to connect to MongoDB: ", err)
	}

	router.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	router.GET("/coupons", func(c *gin.Context) {
		handlers.GetUsers(c, client)
	})

	router.Run(":8080")
}
