package app

import (
	"gostagram/service"

	"github.com/labstack/echo/v4"
)

func Routes() *echo.Echo {
	e := echo.New()

	instaHandler := InstaHandler{Service: service.DefaultInstaService{}}
	e.GET("/", instaHandler.GetPosts)
	return e
}
