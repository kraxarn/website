package group

import (
	"github.com/kraxarn/website/helper"
	"github.com/labstack/echo/v4"
)

func RegisterIndex(app *echo.Echo) {
	group := app.Group("")

	group.GET("/", index)
	group.GET("/about", about)
	group.GET("/projects", projects)
}

func index(ctx echo.Context) error {
	return helper.RenderPage(ctx, "home", nil)
}

func about(ctx echo.Context) error {
	return helper.RenderPage(ctx, "about", nil)
}

func projects(ctx echo.Context) error {
	return helper.RenderPage(ctx, "projects", nil)
}
