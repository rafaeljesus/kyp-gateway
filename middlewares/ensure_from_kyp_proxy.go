package middlewares

import (
	"github.com/labstack/echo"
	"net/http"
)

func EnsurKypProxyHeader(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		header := c.Request().Header().Get("KypProxy")

		if header == "X-Kyp-Proxy" {
			return next(c)
		}

		return c.JSON(http.StatusForbidden, "Forbidden")
	}
}
