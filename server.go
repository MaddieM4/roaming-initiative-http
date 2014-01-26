package main

import (
	"github.com/codegangsta/martini"
	"github.com/codegangsta/martini-contrib/render"
)

type Template struct {
	Handlers []string
	Tmpl     string
	Title    string
}

func setupTemplate(m *martini.ClassicMartini, t Template) {
	handler := func(r render.Render) {
		r.HTML(200, t.Tmpl, t.Title)
	}
	for _, name := range t.Handlers {
		m.Get("/"+name, handler)
	}
}

func newTemplate(tmpl string, title string) Template {
	return Template{
		Handlers: []string{tmpl},
		Tmpl:     tmpl,
		Title:    title,
	}
}

func NewMartiniServer() *martini.ClassicMartini {
	templates := []Template{
		Template{
			Handlers: []string{"", "index", "index.htm", "index.html"},
			Tmpl:     "index",
			Title:    "Home",
		},
		newTemplate("availability", "Goodies"),
		newTemplate("contact", "Goodies"),
		newTemplate("faq", "Goodies"),
		newTemplate("goodies", "Goodies"),
		newTemplate("hardware", "Goodies"),
		newTemplate("pricing", "Goodies"),
	}

	m := martini.Classic()
	m.Use(render.Renderer(render.Options{
		Layout: "layout",
	}))
	m.Use(martini.Static("root"))

	for _, t := range templates {
		setupTemplate(m, t)
	}

	return m
}
