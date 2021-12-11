package main

import (
    "github.com/gin-gonic/gin"
    "k8s.io/apimachinery/pkg/util/rand"
    docs "github.com/huuduc2312/bootcamp-go/src/docs"
    ginSwagger "github.com/swaggo/gin-swagger"
    swaggerFiles "github.com/swaggo/files"
)

// @BasePath /api

// @Summary Get a random string
// @Description Get a random string
// @Tags String
// @Accept json
// @Produce json
// @Success 200 {string} a-random-string
// @Router /strings/random [get]
func getRandomStringfunc(c *gin.Context) {
    c.JSON(200, gin.H{
        "string": rand.String(20),
    })
}

func main() {
    router := gin.Default()
    docs.SwaggerInfo.BasePath = "/api"
    
    router.GET("/api/strings/random", getRandomStringfunc)
    
    router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
    router.Run()
}
