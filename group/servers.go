package group

import (
	"github.com/kraxarn/website/helper"
	"github.com/labstack/echo/v4"
)

func RegisterServers(app *echo.Echo) {
	group := app.Group("/servers")

	group.GET("", servers)
}

func servers(ctx echo.Context) error {
	return helper.RenderPage(ctx, "servers", nil)
}
