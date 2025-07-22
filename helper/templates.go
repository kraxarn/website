package helper

import (
	"github.com/kraxarn/website/config"
	"github.com/labstack/echo/v4"
	"html/template"
	"io"
)

type TemplateRenderer struct {
	templates *template.Template
}

func NewTemplateRenderer() (*TemplateRenderer, error) {
	funcMap := template.FuncMap{
		"currentVersion": func() string {
			return config.Version()
		},
	}

	templates := template.New("")
	templates.Funcs(funcMap)

	_, err := templates.ParseFiles(
		"html/icons/house.gohtml",
		"html/icons/info.gohtml",
		"html/icons/list_ul.gohtml",
		"html/icons/server.gohtml",
		"html/index.gohtml",
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
