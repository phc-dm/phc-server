package templates

import (
	"html/template"
	"io"
	"path"
	"strings"

	"github.com/phc-dm/phc-server/config"
	"github.com/phc-dm/phc-server/util"
)

type LoadTemplate interface {
	Load(t *template.Template) (*template.Template, error)
}

type File string

func (file File) Load(t *template.Template) (*template.Template, error) {
	return t.ParseFiles(string(file))
}

type Pattern string

func (pattern Pattern) Load(t *template.Template) (*template.Template, error) {
	return t.ParseGlob(string(pattern))
}

type CachedTemplate struct {
	Loaders  []LoadTemplate
	template *template.Template
}

func NewCacheTemplate(loaders ...LoadTemplate) *CachedTemplate {
	cachedTemplate := &CachedTemplate{
		Loaders:  loaders,
		template: nil,
	}
	template.Must(cachedTemplate.Load(template.New("")))
	return cachedTemplate
}

func (ct *CachedTemplate) Load(t *template.Template) (*template.Template, error) {
	if ct.template != nil {
		return ct.template, nil
	}

	ct.template = t

	for _, loader := range ct.Loaders {
		var err error
		ct.template, err = loader.Load(ct.template)
		if err != nil {
			return nil, err
		}
	}

	return ct.template, nil
}

func (ct *CachedTemplate) Reload() {
	ct.template = nil
	template.Must(ct.Load(nil))
}

func (ct *CachedTemplate) Template() *template.Template {
	return ct.template
}

// Renderer holds cached templates for rendering
type Renderer struct {
	rootPath    string
	baseLoaders []LoadTemplate
	routes      map[string]*CachedTemplate
}

// NewRenderer constructs a template renderer with a base file
func NewRenderer(rootPath string, loaders ...LoadTemplate) *Renderer {
	return &Renderer{
		rootPath,
		loaders,
		map[string]*CachedTemplate{},
	}
}

func (r *Renderer) Load(name string) *CachedTemplate {
	cachedTemplate, present := r.routes[name]

	if !present {
		loaders := []LoadTemplate{}
		loaders = append(loaders, r.baseLoaders...)
		loaders = append(loaders, File(path.Join(r.rootPath, name)))
		cachedTemplate = NewCacheTemplate(loaders...)
		r.routes[name] = cachedTemplate
	}

	return cachedTemplate
}

// Render the template, also injects "Page" and "Config" values in the template
func (r *Renderer) Render(w io.Writer, name string, data util.H) error {
	cachedTemplate := r.Load(name)

	if config.Mode == "development" {
		cachedTemplate.Reload()
	}

	newData := util.H{}
	newData.Apply(data)
	newData["Page"] = util.H{
		// Used to inject a page specific class on <body>
		"Name": strings.TrimSuffix(path.Base(name), ".html"),
	}
	newData["Config"] = config.Object()

	return cachedTemplate.Template().ExecuteTemplate(w, "base", newData)
}
