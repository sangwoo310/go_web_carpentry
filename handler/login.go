package handler

import (
	"github.com/labstack/echo"
	"net/http"
)

//Login
func Login(c echo.Context) error {
	if err:= c.String(http.StatusOK, "loginCheck!"); err != nil {
		return err
	}
	return nil
}

func Landering(c echo.Context) error {
	c.Render()
}