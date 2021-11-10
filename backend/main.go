package main

import (
 "backend/controller"
 "github.com/labstack/echo/v4"
 "github.com/labstack/echo/v4/middleware"
 "os"
)
func main() {
 port := os.Getenv("PORT")
 server := echo.New()
 server.Use(middleware.CORSWithConfig(middleware.CORSConfig{
  AllowOrigins: []string{"*"},
  AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept},
 }))

 server.GET("/", controller.Hello)
 server.GET("/random", controller.Random)

 server.Logger.Fatal(server.Start(":"+port))
}
