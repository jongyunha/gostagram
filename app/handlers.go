package app

import (
	"gostagram/dto"
	"gostagram/service"
	"net/http"

	"github.com/labstack/echo/v4"
)

type InstaHandler struct {
	Service service.InstaService
}

// @Summary Instagram get posts by user
// @Description Instagram get posts by user
// @Accept  string
// @Success 200 {object} dto.InstaResponse
// @Router  / [get]
func (sh InstaHandler) GetPosts(c echo.Context) error {
	var res dto.InstaResponse

	return c.JSON(http.StatusOK, res)
}
