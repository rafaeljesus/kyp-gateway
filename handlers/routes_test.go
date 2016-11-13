package handlers

import (
	"github.com/labstack/echo"
	"github.com/labstack/echo/engine/standard"
	"github.com/stretchr/testify/assert"
	"gopkg.in/h2non/gock.v1"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestUsersCreate(t *testing.T) {
	defer gock.Off()
	gock.New(KYP_USERS_ENDPOINT).
		Post("/users").
		Reply(201).
		JSON(map[string]int{"id": 1})

	response := `{"id":"1"}`
	e := echo.New()
	req, _ := http.NewRequest(echo.POST, "/api/v1/users", strings.NewReader(response))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	ctx := e.NewContext(standard.NewRequest(req, e.Logger()), standard.NewResponse(rec, e.Logger()))

	if assert.NoError(t, UsersCreate(ctx)) {
		assert.Equal(t, http.StatusCreated, rec.Code)
	}
}

func TestUsersShow(t *testing.T) {
	defer gock.Off()
	gock.New(KYP_USERS_ENDPOINT).
		Get("/users").
		Reply(200).
		JSON(map[string]int{"id": 1})

	response := `{"id":"1"}`
	e := echo.New()
	req, _ := http.NewRequest(echo.GET, "/api/v1/users/1", strings.NewReader(response))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	ctx := e.NewContext(standard.NewRequest(req, e.Logger()), standard.NewResponse(rec, e.Logger()))

	if assert.NoError(t, UsersShow(ctx)) {
		assert.Equal(t, http.StatusOK, rec.Code)
	}
}
