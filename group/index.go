package group

import (
	"github.com/kraxarn/website/helper"
	"github.com/labstack/echo/v4"
)

func RegisterIndex(app *echo.Echo) {
	group := app.Group("/")

	group.GET("", index)
}

func index(ctx echo.Context) error {
	return helper.RenderPage(ctx, "home")
}
