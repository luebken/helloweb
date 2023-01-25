package main

import (
	"embed"
	"html/template"

	"github.com/luebken/helloweb/pkg/web"

	"github.com/gin-gonic/gin"
)

//go:embed templates/*
var resources embed.FS

var templ = template.Must(template.ParseFS(resources, "templates/*"))

func main() {
	r := gin.Default()
	r.SetHTMLTemplate(templ)
	r.GET("/", web.IndexHandler())
	r.Run() // listen and serve on 0.0.0.0:8080
}
