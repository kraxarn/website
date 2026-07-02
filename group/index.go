package group

import (
	"github.com/kraxarn/website/helper"
	"github.com/labstack/echo/v4"
)

func RegisterIndex(app *echo.Echo) {
	group := app.Group("")

	group.GET("/", index)
	group.GET("/:page", page)
}

func index(ctx echo.Context) error {
	return helper.RenderPage(ctx, "home", nil)
}

func page(ctx echo.Context) error {
	key := ctx.Param("page")
	return helper.RenderPage(ctx, key, nil)
}
