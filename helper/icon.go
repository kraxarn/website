package helper

import (
	"bytes"
	"fmt"
	"html/template"
)

func icon(templates *template.Template, name string) (template.HTML, error) {
	var buffer bytes.Buffer

	templateName := fmt.Sprintf("icons/%s", name)
	err := templates.ExecuteTemplate(&buffer, templateName, nil)
	if err != nil {
		return "", err
	}

	return template.HTML(buffer.String()), nil
}
