package main

import (
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"go_web_carpentry/router"
)

func main() {
	// Echo instance
	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// Router
	router.New(e)

	// Start server
	e.Logger.Fatal(e.Start(":3100"))

}