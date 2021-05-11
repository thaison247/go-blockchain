package app

import (
	"errors"
	"fmt"
	"html/template"
	"io"

	"github.com/thaison247/go-blockchain/utils"

	"github.com/labstack/echo/v4"
)

type TemplateRegistry struct {
	templates map[string]*template.Template
}

func (t *TemplateRegistry) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	tmpl, ok := t.templates[name]
	if !ok {
		err := errors.New("Template not found -> " + name)
		return err
	}
	return tmpl.ExecuteTemplate(w, "base.html", data)
}

func WebController(e *echo.Echo) {
	e.Static("/", "public")

	commonFile := []string{
		"view/common/base.html", "view/common/header.html", "view/common/main_content.html", "view/common/topbar.html",
		"view/common/footer.html", "view/common/leftbar.html", "view/common/navbar.html"}

	templates := make(map[string]*template.Template)
	for _, value := range utils.ARR_TEMPLATES {
		viewName := fmt.Sprintf("view/%s", value)
		templates[value] = template.Must(template.ParseFiles(append(commonFile, viewName)...))
	}

	renderer := &TemplateRegistry{
		templates: templates,
	}
	e.Renderer = renderer
}

