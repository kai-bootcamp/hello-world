package main

import (
 "backend/controller"
 "github.com/labstack/echo/v4"
 "github.com/labstack/echo/v4/middleware"
)
func main() {
 server := echo.New()
 server.Use(middleware.CORSWithConfig(middleware.CORSConfig{
  AllowOrigins: []string{"*"},
  AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept},
 }))
 general := server.Group("/general")
 {
  general.GET("/hello", controller.Hello)
  general.GET("/random", controller.Random)
 }

 server.Logger.Fatal(server.Start(":8081"))
}
