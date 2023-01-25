package main

import (
	"embed"
	"html/template"
	"net/http"

	"github.com/luebken/helloweb/pkg/web"

	"github.com/gin-gonic/gin"
)

//go:embed templates/*
var resources embed.FS
var templ = template.Must(template.ParseFS(resources, "templates/*"))

//go:embed static/*
var staticFiles embed.FS
var staticFS = http.FS(staticFiles)

func main() {
	r := gin.Default()
	r.SetHTMLTemplate(templ)
	r.StaticFS("/static/", staticFS)
	r.GET("/", web.IndexHandler())
	r.Run() // listen and serve on 0.0.0.0:8080
}
