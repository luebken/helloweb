package main

import (
	"embed"
	"fmt"
	"html/template"
	"net/http"

	"github.com/luebken/helloweb/pkg/web"
	"github.com/luebken/helloweb/pkg/web/db"

	"github.com/gin-gonic/gin"
)

//go:embed templates/*
var resources embed.FS
var templ = template.Must(template.ParseFS(resources, "templates/*"))

//go:embed static/*
var staticFiles embed.FS
var staticFS = http.FS(staticFiles)

func main() {

	db := db.Database{}
	err := db.Open()
	if err != nil {
		fmt.Printf("Err %v\n", err)
	}
	defer db.Close()

	r := gin.Default()
	r.SetHTMLTemplate(templ)
	r.StaticFS("/static/", staticFS)
	r.GET("/", web.IndexHandler(db))
	r.Run() // listen and serve on 0.0.0.0:8080
}
