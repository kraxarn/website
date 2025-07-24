package helper

import (
	"bytes"
	"github.com/kraxarn/website/db"
	"github.com/kraxarn/website/repo"
	"github.com/labstack/echo/v4"
	"github.com/yuin/goldmark"
	"html/template"
	"io"
	"net/http"
)

type TemplateRenderer struct {
	templates *template.Template
}

func NewTemplateRenderer() (*TemplateRenderer, error) {
	funcMap := template.FuncMap{}

	templates := template.New("")
	templates.Funcs(funcMap)

	_, err := templates.ParseFiles(
		"html/editor.gohtml",
		"html/icons/house.gohtml",
		"html/icons/info.gohtml",
		"html/icons/list_ul.gohtml",
		"html/icons/server.gohtml",
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

func (r *TemplateRenderer) Render(writer io.Writer, name string, data interface{}, _ echo.Context) error {
	return r.templates.ExecuteTemplate(writer, name, data)
}

func RenderPage(ctx echo.Context, key string) error {
	conn, err := db.Acquire()
	if err != nil {
		return err
	}
	defer conn.Release()

	texts := repo.NewTexts(conn)

	var val string
	val, err = texts.Value(key)
	if err != nil {
		return echo.NewHTTPError(http.StatusNotFound, err)
	}

	var content template.HTML
	content, err = RenderMarkdown(val)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err)
	}

	return ctx.Render(http.StatusOK, "page.gohtml", map[string]interface{}{
		"content": content,
	})
}

func RenderMarkdown(content string) (template.HTML, error) {
	var buf bytes.Buffer
	err := goldmark.Convert([]byte(content), &buf)
	if err != nil {
		return "", err
	}

	return template.HTML(buf.String()), nil
}
