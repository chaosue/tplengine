package tplengine

import (
	"github.com/labstack/echo"
	"html/template"
	"io"
)

type Renderer struct {
	*template.Template
}

type DebugRenderer struct {
	name string
	*template.Template
	globFilePatern string
	funcMap        template.FuncMap
}

func NewDebugRenderer(name string) *DebugRenderer {
	return &DebugRenderer{
		Template: template.New(name).Funcs(plugins),
		name:     name,
	}
}

func NewRenderer(name string) *Renderer {
	return &Renderer{
		Template: template.New(name).Funcs(plugins),
	}
}

func (r *Renderer) Funcs(funcMap template.FuncMap) *Renderer {
	for n, v := range plugins {
		funcMap[n] = v
	}
	r.Template.Funcs(funcMap)
	return r
}

func (r *Renderer) Render(w io.Writer, name string, data interface{}, _ echo.Context) error {
	return r.ExecuteTemplate(w, name, data)
}

func (r *Renderer) ParseGlob(patern string) error {
	_, err := r.Template.ParseGlob(patern)
	return err
}

func (r *DebugRenderer) Funcs(funcMap template.FuncMap) *DebugRenderer {
	r.funcMap = funcMap
	for n, v := range plugins {
		funcMap[n] = v
	}
	r.Template.Funcs(funcMap)
	return r
}

func (r *DebugRenderer) Render(w io.Writer, name string, data interface{}, _ echo.Context) error {
	r.Template = template.New(r.name)
	allPlugins := plugins
	if r.funcMap != nil {
		for n, v := range r.funcMap {
			allPlugins[n] = v
		}
	}
	r.Template.Funcs(allPlugins)
	r.Template.ParseGlob(r.globFilePatern)
	return r.ExecuteTemplate(w, name, data)
}

func (r *DebugRenderer) ParseGlob(patern string) error {
	r.globFilePatern = patern
	return nil
}
