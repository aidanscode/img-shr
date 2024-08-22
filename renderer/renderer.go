package renderer

import (
	"fmt"
	"html/template"
	"io"
	"strings"

	"github.com/labstack/echo/v4"
)

type Renderer struct {
	templates map[string]*template.Template
}

func (r *Renderer) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	split := strings.Split(name, ".")
	if len(split) == 1 {
		return r.templates[name].ExecuteTemplate(w, "layout", data)
	} else {
		return r.templates[split[0]].ExecuteTemplate(w, split[1], data)
	}
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
