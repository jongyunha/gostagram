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
// @Produce  json
// @Param   user_id     query    string     true        "user id"
// @Param   access_token     query    string     true        "access token"
// @Success 200 {object} dto.InstaResponse
// @Router  /api/v1/instagram/posts [get]
func (sh InstaHandler) GetPosts(c echo.Context) error {
	userId := c.QueryParam("user_id")
	accessToken := c.QueryParam("access_token")

	var req dto.InstaRequest
	req.UserId = userId
	req.AccessToken = accessToken

	if err := req.Validate(); err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	res, err := sh.Service.GetPosts(req)

	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	return c.JSON(http.StatusOK, res)
}

// func (sh InstaHandler) GetChildren(c echo.Context) error {
//
// }
