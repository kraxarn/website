package helper

import (
	"bytes"
	"github.com/kraxarn/website/config"
	"github.com/labstack/echo/v4"
	"github.com/yuin/goldmark"
	"html/template"
	"io"
)

type TemplateRenderer struct {
	templates *template.Template
}

func NewTemplateRenderer() (*TemplateRenderer, error) {
	funcMap := template.FuncMap{
		"renderMarkdown": renderMarkdown,
	}

	templates := template.New("")
	templates.Funcs(funcMap)

	_, err := templates.ParseFiles(
		"html/editor.gohtml",
		"html/icons/house.gohtml",
		"html/icons/info.gohtml",
		"html/icons/list_ul.gohtml",
		"html/icons/server.gohtml",
		"html/index.gohtml",
		"html/login.gohtml",
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

func renderMarkdown(content string) template.HTML {
	var buf bytes.Buffer
	err := goldmark.Convert([]byte(content), &buf)
	if err != nil {
		if config.Dev() {
			return template.HTML(err.Error())
		}
		return "error: invalid template (this shouldn't happen!)"
	}

	return template.HTML(buf.String())
}
