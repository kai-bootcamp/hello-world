package controller

import (
    "backend/model"
    "backend/utils"
 "github.com/labstack/echo/v4"
 "net/http"
)

const stringLength = 125

func Hello(c echo.Context) error {
 return c.JSON(http.StatusOK, "hello world")
}

func Random(c echo.Context) error {
    var response model.Response
    str, err := utils.String(stringLength)
    if err != nil {
        response.Success = false
        return c.JSON(http.StatusBadRequest, response)
    }
    response.Success = true
    response.Data = str
   return c.JSON(http.StatusOK, response)
}