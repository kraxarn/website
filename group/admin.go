package group

import (
	"context"
	"fmt"
	"html/template"
	"net/http"
	"strconv"

	"github.com/kraxarn/website/data"
	"github.com/kraxarn/website/db"
	"github.com/kraxarn/website/helper"
	"github.com/kraxarn/website/repo"
	"github.com/labstack/echo/v5"
)

type editorContent struct {
	Key   string `form:"key"`
	Value string `form:"value"`
	Type  string `form:"type"`
}

type itemsContent struct {
	Items []string `form:"items"`
}

func RegisterAdmin(app *echo.Echo) {
	group := app.Group("/admin")

	group.GET("/editor", editor)
	group.POST("/editor", editorData)

	group.GET("/items", items)
	group.POST("/items", itemsData)
}

func editor(ctx *echo.Context) error {
	return helper.Render(ctx, http.StatusOK, "editor.gohtml", nil)
}

func userIdFromContext(ctx *echo.Context) (db.Id, error) {
	claims, err := data.ParseUserClaims(ctx)
	if err != nil {
		return 0, err
	}

	var userFlags data.UserFlags
	userFlags, err = claims.UserFlags()
	if err != nil || (userFlags&data.UserFlagsEditor) == 0 {
		var message string
		if err != nil {
			message = err.Error()
		} else {
			message = "invalid flag"
		}
		return 0, echo.NewHTTPError(http.StatusForbidden, message)
	}

	var userId db.Id
	userId, err = claims.UserId()
	if err != nil {
		return 0, err
	}

	return userId, nil
}

func editorData(ctx *echo.Context) error {
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
		value, err = texts.Value(content.Key)
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
		err = echo.NewHTTPError(http.StatusNotFound, "invalid type")
	}

	if err != nil {
		return err
	}

	var preview template.HTML
	preview, err = helper.RenderMarkdown(value)
	if err != nil {
		preview = template.HTML(fmt.Sprintf("<pre>%s</pre>", value))
	}

	return helper.Render(ctx, http.StatusOK, "editor.gohtml", map[string]interface{}{
		"key":     content.Key,
		"value":   value,
		"preview": preview,
	})
}

func items(ctx *echo.Context) error {
	conn, err := db.Acquire()
	if err != nil {
		return err
	}
	defer conn.Release()

	itemRepo := repo.NewItemsFromPool(conn)

	var rows []repo.EditItem
	rows, err = itemRepo.SelectAllForEdit()
	if err != nil {
		return err
	}

	return helper.Render(ctx, http.StatusOK, "items.gohtml", map[string]any{
		"editItems": rows,
	})
}

func itemsData(ctx *echo.Context) error {
	var content itemsContent
	if err := ctx.Bind(&content); err != nil {
		return err
	}

	conn, err := db.Acquire()
	if err != nil {
		return err
	}
	defer conn.Release()

	tx, err := conn.Begin(context.Background())
	if err != nil {
		return err
	}

	itemsRepo := repo.NewItemsFromTx(tx)

	rollback := func() error {
		if txErr := tx.Rollback(context.Background()); txErr != nil {
			return txErr
		}
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	if err = itemsRepo.DeleteAll(); err != nil {
		return rollback()
	}

	for _, item := range content.Items {
		priorityValue := ctx.FormValue(fmt.Sprintf("%s/priority", item))
		var priority int64
		priority, err = strconv.ParseInt(priorityValue, 10, 32)
		if err != nil {
			return rollback()
		}

		value := ctx.FormValue(fmt.Sprintf("%s/value", item))
		icon := ctx.FormValue(fmt.Sprintf("%s/icon", item))
		if len(value) == 0 || len(icon) == 0 {
			return rollback()
		}

		_, err = itemsRepo.Insert(item, value, icon, int(priority))
		if err != nil {
			return rollback()
		}
	}

	if err = tx.Commit(context.Background()); err != nil {
		return err
	}

	return items(ctx)
}
