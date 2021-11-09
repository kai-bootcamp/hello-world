package controller

import (
 "github.com/labstack/echo/v4"
 "net/http"
)

func Hello(c echo.Context) error {
 return c.JSON(http.StatusOK, "hello world")
}