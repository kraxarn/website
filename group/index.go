package group

import (
	"github.com/labstack/echo/v4"
	"net/http"
)

func RegisterIndex(app *echo.Echo) {
	group := app.Group("/")

	group.GET("", index)
}

func index(ctx echo.Context) error {
	return ctx.Render(http.StatusOK, "index.gohtml", nil)
}
