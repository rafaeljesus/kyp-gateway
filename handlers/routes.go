package handlers

import (
	"encoding/json"
	"github.com/labstack/echo"
	"github.com/rafaeljesus/kyp-structs"
	"net/http"
	"os"
)

var KYP_USERS_ENDPOINT = os.Getenv("KYP_USERS_ENDPOINT")
var KYP_AUTH_ENDPOINT = os.Getenv("KYP_AUTH_ENDPOINT")
var KYP_TODO_ENDPOINT = os.Getenv("KYP_TODO_ENDPOINT")

type User structs.User
type Todo structs.Todo

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
	id := c.Param("id")

	res, err := http.Get(KYP_USERS_ENDPOINT + "/" + id)
	if err != nil {
		return err
	}

	defer res.Body.Close()

	user := User{}
	json.NewDecoder(res.Body).Decode(&user)

	return c.JSON(http.StatusOK, user)
}

func TokenCreate(c echo.Context) error {
	res, err := http.Post(KYP_AUTH_ENDPOINT, "application/json; charset=utf-8", c.Request().Body())
	if err != nil {
		return err
	}

	defer res.Body.Close()

	var token struct{ Token string }
	json.NewDecoder(res.Body).Decode(&token)

	return c.JSON(http.StatusOK, token)
}

func TodosIndex(c echo.Context) error {
	res, err := http.Get(KYP_TODO_ENDPOINT)
	if err != nil {
		return err
	}

	defer res.Body.Close()

	todos := []Todo{}
	json.NewDecoder(res.Body).Decode(&todos)

	return c.JSON(http.StatusOK, todos)
}

func TodosCreate(c echo.Context) error {
	res, err := http.Post(KYP_TODO_ENDPOINT, "application/json; charset=utf-8", c.Request().Body())
	if err != nil {
		return err
	}

	defer res.Body.Close()

	todo := Todo{}
	json.NewDecoder(res.Body).Decode(&todo)

	return c.JSON(http.StatusCreated, todo)
}
