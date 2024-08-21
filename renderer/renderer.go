package renderer

import (
	"fmt"
	"html/template"
	"io"

	"github.com/labstack/echo/v4"
)

type Renderer struct {
	templates map[string]*template.Template
}

func (r *Renderer) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	return r.templates[name].ExecuteTemplate(w, "layout", data)
}

func NewRenderer() *Renderer {
	templates := make(map[string]*template.Template)
	templates["index"] = tryLoadTemplateWithLayout("index")
	templates["upload"] = tryLoadTemplateWithLayout("upload")

	renderer := &Renderer{
		templates: templates,
	}
	return renderer
}

func tryLoadTemplateWithLayout(templateName string) *template.Template {
	template, err := template.New("").ParseFiles(fmt.Sprintf("views/%s.html", templateName), "views/navbar.html", "views/layout.html")
	if err != nil {
		panic(err)
	}

	return template
}