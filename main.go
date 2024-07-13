package main

import (
	"github.com/SiwakornEDZ/awesomeProject/api/handlers"
	"github.com/SiwakornEDZ/awesomeProject/db/db"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

func main() {
	router := gin.Default()

	client, err := ConnectToDB()
	if err != nil {
		log.Fatal("Failed to connect to MongoDB: ", err)
	}

	router.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	router.GET("/coupons", func(c *gin.Context) {
		getUsers(c, client)
	})

	router.Run(":8080")
}
