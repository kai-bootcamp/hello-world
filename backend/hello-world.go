package main

import (
	"fmt"

	"github.com/gin-gonic/gin"

	"os"

	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load()
	port := ":" + os.Getenv("PORT")
	fmt.Println("Hello World!")
	router := gin.Default()

	router.GET("/hello-world", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Hello Khoa, welcome to new world! Let's make something fantastic ;)",
		})
	})

	router.Run(port)
}
