package group

import (
	"github.com/labstack/echo/v4"
	"net/http"
)

type editorContent struct {
	Key   string `form:"key"`
	Value string `form:"value"`
	Type  string `form:"type"`
}

func RegisterAdmin(app *echo.Echo) {
	group := app.Group("/admin")

	group.GET("/editor", editor)
	group.POST("/editor", editorData)
}

func editor(ctx echo.Context) error {
	return ctx.Render(http.StatusOK, "editor.gohtml", nil)
}

func editorData(ctx echo.Context) error {
	var content editorContent
	if err := ctx.Bind(&content); err != nil {
		return err
	}

	return ctx.Render(http.StatusOK, "editor.gohtml", nil)
}
