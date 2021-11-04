package app

import (
	"gostagram/service"

	"github.com/labstack/echo/v4"
)

func Routes() *echo.Echo {
	e := echo.New()

	instaHandler := InstaHandler{Service: service.DefaultInstaService{}}
	e.GET("/api/v1/instagram/posts", instaHandler.GetPosts)
	e.GET("/api/v1/instagram/child", instaHandler.GetChild)
	return e
}
