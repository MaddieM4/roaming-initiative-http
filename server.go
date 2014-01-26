package main

import (
	"encoding/json"
	"github.com/codegangsta/martini"
	"github.com/codegangsta/martini-contrib/render"
	"io/ioutil"
	"log"
	"net/http"
)

type Template struct {
	Handlers []string
	Tmpl     string
	Title    string
}

type ServerList map[string]string

type Config struct {
	Servers   ServerList
	Templates []Template
}

func GetConfig() Config {
	bytes_data, err := ioutil.ReadFile("./config.json")
	if err != nil {
		log.Fatal(err)
	}

	var conf Config
	err = json.Unmarshal(bytes_data, &conf)
	if err != nil {
		log.Fatal(err)
	}
	return conf
}

type ServerInfo struct {
	Ip   string
	Name string
}

func getServerInfo(req *http.Request, s ServerList) ServerInfo {
	info := ServerInfo{Ip: "0.0.0.0", Name: "dev"}

	if hlist, ok := req.Header["Server-Ip"]; ok {
		info.Ip = hlist[0]
	}
	if server_name, ok := s[info.Ip]; ok {
		info.Name = server_name
	}

	return info
}

func handleTemplate(r render.Render, req *http.Request, t Template, s ServerList) {
	tmpl_data := make(map[string]interface{})
	tmpl_data["title"] = t.Title
	tmpl_data["server"] = getServerInfo(req, s)

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

func NewMartiniServer() *martini.ClassicMartini {
	conf := GetConfig()

	m := martini.Classic()
	m.Use(render.Renderer(render.Options{
		Layout: "layout",
	}))
	m.Use(martini.Static("root"))
	m.Map(conf.Servers)

	for _, t := range conf.Templates {
		setupTemplate(m, t)
	}

	return m
}
