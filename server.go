package main

import (
	"github.com/codegangsta/martini"
	"github.com/codegangsta/martini-contrib/render"
	"net/http"
)

type Template struct {
	Handlers []string
	Tmpl     string
	Title    string
}

type ServerList map[string]string

func NewServerList() ServerList {
	servers := make(ServerList)
	servers["173.255.210.202"] = "clearnet4"
	servers["2600:3c01::f03c:91ff:feae:1082"] = "clearnet6"
	servers["fcd5:7d07:2146:f18f:f937:d46e:77c9:80e7"] = "hyperboria"

	return servers
}

func handleTemplate(r render.Render, req *http.Request, t Template, s ServerList) {
	var server_ip string
	if hlist, ok := req.Header["Server-Ip"]; ok {
		server_ip = hlist[0]
	} else {
		server_ip = "0.0.0.0"
	}

	tmpl_data := make(map[string]interface{})
	tmpl_data["title"] = t.Title
	tmpl_data["server_ip"] = server_ip

	if server_name, ok := s[server_ip]; ok {
		tmpl_data["server_name"] = server_name
	} else {
		tmpl_data["server_name"] = "dev"
	}

	r.HTML(200, t.Tmpl, tmpl_data)
}

func setupTemplate(m *martini.ClassicMartini, t Template) {
	handler := func(r render.Render, req *http.Request, s ServerList) {
		handleTemplate(r, req, t, s)
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
	m.Map(NewServerList())

	for _, t := range templates {
		setupTemplate(m, t)
	}

	return m
}
