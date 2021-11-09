package main

import (
 "backend/controller"
 "github.com/labstack/echo/v4"
)
func main() {
 server := echo.New()

 general := server.Group("/general")
 {
  general.GET("/hello", controller.Hello)
 }

 server.Logger.Fatal(server.Start(":8081"))
}
