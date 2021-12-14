package utils

import (
	"io"
	"text/template"

	"github.com/phc-dm/server-poisson/config"
)

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
func (t *TemplateRenderer) Render(w io.Writer, name string, data interface{}) error {
	tmpl := t.templateMap[name]

	if config.Mode == "development" || tmpl == nil {
		tmpl = template.Must(t.baseTemplate.Clone())

		if config.Mode == "development" {
			tmpl = template.Must(template.ParseFiles("./views/" + t.baseFile))
		} else {
			tmpl = template.Must(t.baseTemplate.Clone())
		}

		tmpl.ParseFiles("./views/" + name)
		t.templateMap[name] = tmpl
	}

	return tmpl.ExecuteTemplate(w, "base", data)
}
