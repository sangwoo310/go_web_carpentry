package router

import (
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"net/http"

)

type TStructRouter struct {
	handlers map[string]map[string]http.HandlerFunc
}

func (r *TStructRouter)TRouter() *echo.Echo {
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

	return nil

}