package handler

import (
	"github.com/labstack/echo"
	"log"
	"net/http"
)

//Login
func Login(c echo.Context) error {
	if err := c.String(http.StatusOK, "loginCheck!"); err != nil {
		return err
	}
	return nil
}

func Rendering(c echo.Context) error {
	log.Println("start rendering")
	//c.SetRequest()

	if err := c.Render(http.StatusOK, "views/index.html", nil); err != nil {
		return err
	}

	//if err := c.File("views/index.html"); err != nil {
	//	return err
	//}

	return nil
}
