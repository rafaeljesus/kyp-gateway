package main

import (
	"github.com/labstack/echo"
	"github.com/labstack/echo/engine/fasthttp"
	"github.com/labstack/echo/middleware"
	"github.com/rafaeljesus/kyp-gateway/handlers"
	"log"
	"os"
)

var KYP_GATEWAY_PORT = os.Getenv("KYP_GATEWAY_PORT")

func main() {
	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.CORS())
	e.Use(middleware.Secure())
	e.Use(middleware.Gzip())

	v1 := e.Group("/api/v1")
	v1.GET("/healthz", handlers.HealthzIndex)
	v1.POST("/users", handlers.UsersCreate)
	v1.POST("/token", handlers.TokenCreate)

	v1.Use(middleware.JWT([]byte(os.Getenv("KYP_SECRET_KEY"))))

	v1.GET("/users/:id", handlers.UsersShow)
	v1.GET("/todos", handlers.TodosIndex)
	v1.POST("/todos", handlers.TodosCreate)

	log.Print("Starting Kyp Gateway Service...", KYP_GATEWAY_PORT)

	e.Run(fasthttp.New(":" + KYP_GATEWAY_PORT))
}
