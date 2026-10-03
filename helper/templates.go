package helper

import (
	"bytes"
	"html/template"
	"io"
	"net/http"

	"github.com/kraxarn/website/db"
	"github.com/kraxarn/website/repo"
	"github.com/labstack/echo/v5"
	"github.com/yuin/goldmark"
	"github.com/yuin/goldmark/extension"
)

type TemplateRenderer struct {
	templates *template.Template
}

func NewTemplateRenderer() (*TemplateRenderer, error) {
	templates := template.New("")

	funcMap := template.FuncMap{
		"static": staticFileVersion,
		"icon": func(name string) (template.HTML, error) {
			return icon(templates, name)
		},
	}

	templates.Funcs(funcMap)

	_, err := templates.ParseFiles(
		"html/editor.gohtml",
		"html/error.gohtml",
		"html/icons/house.gohtml",
		"html/icons/info.gohtml",
		"html/icons/list_ul.gohtml",
		"html/icons/server.gohtml",
		"html/items.gohtml",
		"html/login.gohtml",
		"html/page.gohtml",
		"html/partials/footer.gohtml",
		"html/partials/header.gohtml",
		"html/partials/layout_begin.gohtml",
		"html/partials/layout_end.gohtml",
	)

	if err != nil {
		return nil, err
	}

	return &TemplateRenderer{
		templates: templates,
	}, nil
}

func (r *TemplateRenderer) Render(_ *echo.Context, writer io.Writer, name string, data any) error {
	return r.templates.ExecuteTemplate(writer, name, data)
}

func Render(ctx *echo.Context, code int, name string, data map[string]any) error {
	conn, err := db.Acquire()
	if err != nil {
		return err
	}
	defer conn.Release()

	itemsRepo := repo.NewItemsFromPool(conn)

	var items []repo.Item
	items, err = itemsRepo.SelectAll()
	if err != nil {
		return err
	}

	if data == nil {
		data = map[string]any{}
	}
	data["items"] = items

	return ctx.Render(code, name, data)
}

func RenderPage(ctx *echo.Context, key string, data map[string]any) error {
	conn, err := db.Acquire()
	if err != nil {
		return err
	}
	defer conn.Release()

	texts := repo.NewTexts(conn)

	var val string
	val, err = texts.Value(key)
	if err != nil {
		return echo.NewHTTPError(http.StatusNotFound, err.Error())
	}

	var content template.HTML
	content, err = RenderMarkdown(val)
	if err != nil {
		return err
	}

	if data == nil {
		data = map[string]any{}
	}
	data["content"] = content

	return Render(ctx, http.StatusOK, "page.gohtml", data)
}

func RenderMarkdown(content string) (template.HTML, error) {
	markdown := goldmark.New(
		goldmark.WithExtensions(extension.Table),
	)

	var buf bytes.Buffer
	err := markdown.Convert([]byte(content), &buf)
	if err != nil {
		return "", err
	}

	return template.HTML(buf.String()), nil
}
