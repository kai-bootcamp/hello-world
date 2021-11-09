package controller

import (
	"github.com/labstack/echo/v4"
	"net/http"
)

func GetStr(c echo.Context) error {
	return c.JSON(http.StatusOK, "Hellooo")
}