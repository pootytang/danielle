package handlers

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/pootytang/danielleapi/types"
)

func ShowPrivateRoutes(c echo.Context) error {
	route := types.Route{}
	return c.JSON(http.StatusOK, route.GetPrivateRoutes())
}
