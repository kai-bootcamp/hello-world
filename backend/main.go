package main

import (
	// "fmt"
	"math/rand"
	"net/http"
	"os"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

var letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")
var randString string

func randSeq(n int) string {
	b := make([]rune, n)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	return string(b)
}

func hw1(c echo.Context) error {
	randString = "Hello World"
	return c.JSON(http.StatusOK, randString)
}

func hw2(c echo.Context) error {
	// http.HandleFunc("/random", handlerRandomStr)
	rand.Seed(time.Now().UnixNano())
	randString = randSeq(10)
	return c.JSON(http.StatusOK, randString)
}

func main() {

	port := os.Getenv("PORT")
	server := echo.New()
	server.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept},
	}))

	server.GET("/hw1", hw1)
	server.GET("/hw2", hw2)
	
	// server.Logger.Fatal(server.Start(":8081"))
	server.Logger.Fatal(server.Start(":"+port))
	
}
