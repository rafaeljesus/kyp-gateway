package handlers

import (
	"github.com/labstack/echo"
	"net/http"
)

func UsersCreate(c echo.Context) error {
	return c.JSON(http.StatusCreated, nil)
}

func UsersShow(c echo.Context) error {
	return c.JSON(http.StatusOK, nil)
}

func TokenCreate(c echo.Context) error {
	return c.JSON(http.StatusOK, nil)
}

func TodosIndex(c echo.Context) error {
	return c.JSON(http.StatusOK, nil)
}

func TodosCreate(c echo.Context) error {
	return c.JSON(http.StatusOK, nil)
}
