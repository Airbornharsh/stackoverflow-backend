package main

import (
	"os"

	"github.com/airbornharsh/stackoverflow-backend/internal/database"
	"github.com/gin-gonic/gin"
)

func main() {
	database.INIT()

	Port := os.Getenv("PORT")
	if Port == "" {
		Port = "8080"
	}

	r := gin.New()

	api := r.Group("/api")

	api.GET("", func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "Hello World"})
	})

	r.Run() //
}
