package router

import (
	"github.com/labstack/echo"
	"net/http"
)

func New(e *echo.Echo)  {
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!\n")
	})
}