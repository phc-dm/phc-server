package utils

import (
	"io"
	"os"
	"text/template"

	"github.com/labstack/echo/v4"
)

var isDevMode bool

func init() {
	isDevMode = os.Getenv("MODE") == "development"
}

// TemplateRenderer holds cached templates for rendering
type TemplateRenderer struct {
	baseFile     string
	baseTemplate *template.Template
	templateMap  map[string]*template.Template
}

// NewTemplateRenderer constructs a template renderer with a base file
func NewTemplateRenderer(baseFile string) *TemplateRenderer {
	return &TemplateRenderer{
		baseFile:     baseFile,
		baseTemplate: template.Must(template.ParseFiles("./views/" + baseFile)),
		templateMap:  make(map[string]*template.Template),
	}
}

// Render the template
func (t *TemplateRenderer) Render(w io.Writer, name string, data interface{}, c echo.Context) error {

	tmpl := t.templateMap[name]
	if isDevMode || tmpl == nil {

		if isDevMode {
			tmpl = template.Must(template.ParseFiles("./views/" + t.baseFile))
		} else {
			tmpl = template.Must(t.baseTemplate.Clone())
		}

		tmpl.ParseFiles("./views/" + name)
		t.templateMap[name] = tmpl
	}

	// // Add global methods if data is a map
	// if viewContext, isMap := data.(map[string]interface{}); isMap {
	// 	viewContext["reverse"] = c.Echo().Reverse
	// }

	return tmpl.ExecuteTemplate(w, "base", data)
}
