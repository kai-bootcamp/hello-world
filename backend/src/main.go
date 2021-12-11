package main

import (
    "github.com/gin-gonic/gin"
    "k8s.io/apimachinery/pkg/util/rand"
)

func main() {
    router := gin.Default()

    router.GET("/api/strings/random", func(c *gin.Context) {
        c.JSON(200, gin.H{
            "string": rand.String(20),
        })
    })

    router.Run()
}
