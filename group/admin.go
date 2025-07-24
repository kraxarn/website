package group

import (
	"github.com/labstack/echo/v4"
	"net/http"
)

func RegisterAdmin(app *echo.Echo) {
	group := app.Group("/admin")

	group.GET("/editor", editor)
}

func editor(ctx echo.Context) error {
	return ctx.Render(http.StatusOK, "editor.gohtml", nil)
}
