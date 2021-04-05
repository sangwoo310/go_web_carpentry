package router

import (
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"go_web_carpentry/handler"
	"net/http"
)

type StructRouter struct {
	option bool
}

func Router() *echo.Echo {
	// ehco.New() 를 통해 *Echo 리턴 및 객체 생성
	e := echo.New()

	// Debug 모드
	e.Debug = true

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{ //CORS Middleware
		AllowOrigins: []string{"*"},
		AllowMethods: []string{http.MethodGet, http.MethodPut, http.MethodPost, http.MethodDelete},
	}))

	// Health check
	e.GET("/health", func(c echo.Context) error {
		return c.String(http.StatusOK, "Server is healthy!")
	})

	e.GET("/*", func(c echo.Context) error {
		return c.String(http.StatusOK, "This is Index Page")
	})

	test := e.Group("/test")
	{
		test.GET("/test", handler.Rendering)
	}

	login := e.Group("/login")
	{
		login.GET("/*", handler.Login)
		//login.GET("/*", handler.Login)
	}

	return e
}
