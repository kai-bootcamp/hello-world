package main

import (
    "github.com/gin-contrib/cors"
    "github.com/gin-gonic/gin"
    "k8s.io/apimachinery/pkg/util/rand"
)

func main() {
    router := gin.Default()

    router.Use(cors.New(cors.Config{AllowAllOrigins: true}))

    router.GET("/api/strings/random", func(c *gin.Context) {
        c.JSON(200, gin.H{
            "string": rand.String(20),
        })
    })

    router.Run()
}
