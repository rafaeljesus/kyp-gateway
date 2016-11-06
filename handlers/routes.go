package handlers

import (
	"encoding/json"
	"github.com/labstack/echo"
	"github.com/rafaeljesus/kyp-structs"
	"net/http"
	"os"
)

var KYP_USERS_ENDPOINT = os.Getenv("KYP_USERS_ENDPOINT")

type User structs.User

func UsersCreate(c echo.Context) error {
	res, err := http.Post(KYP_USERS_ENDPOINT, "application/json; charset=utf-8", c.Request().Body())
	if err != nil {
		return err
	}

	defer res.Body.Close()

	user := User{}
	json.NewDecoder(res.Body).Decode(&user)

	return c.JSON(http.StatusCreated, user)
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
