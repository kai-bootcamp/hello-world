package main

import (
	"backend/controller"

	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()

	// Routes
	e.GET("/reactjs", controller.GetStr)

	// Start server
	e.Logger.Fatal(e.Start(":1323"))
}
