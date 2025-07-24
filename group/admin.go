package group

import (
	"github.com/kraxarn/website/data"
	"github.com/kraxarn/website/db"
	"github.com/kraxarn/website/repo"
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

func userIdFromContext(ctx echo.Context) (db.Id, error) {
	claims, err := data.ParseUserClaims(ctx)
	if err != nil {
		return 0, err
	}

	var userFlags data.UserFlags
	userFlags, err = claims.UserFlags()
	if err != nil || (userFlags&data.UserFlagsEditor) == 0 {
		return 0, echo.NewHTTPError(http.StatusForbidden, err)
	}

	var userId db.Id
	userId, err = claims.UserId()
	if err != nil {
		return 0, err
	}

	return userId, nil
}

func editorData(ctx echo.Context) error {
	var content editorContent
	if err := ctx.Bind(&content); err != nil {
		return err
	}

	userId, err := userIdFromContext(ctx)
	if err != nil {
		return err
	}

	conn, err := db.Acquire()
	if err != nil {
		return err
	}
	defer conn.Release()

	texts := repo.NewTexts(conn)

	var value string

	switch content.Type {
	case "Load":
		value, err = texts.Value(content.Key, userId)
	case "Save":
		var exists bool
		exists, err = texts.Exists(content.Key)
		if err != nil {
			break
		} else if exists {
			_, err = texts.Update(content.Key, content.Value, userId)
		} else {
			_, err = texts.Insert(content.Key, content.Value, userId)
		}
		value = content.Value
	case "Preview":
		value = content.Value
	default:
		err = echo.NewHTTPError(http.StatusNotFound)
	}

	if err != nil {
		return err
	}

	return ctx.Render(http.StatusOK, "editor.gohtml", map[string]string{
		"key":   content.Key,
		"value": value,
	})
}
