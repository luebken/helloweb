package web

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/luebken/helloweb/pkg/web/db"
)

func IndexHandler(db db.Database) gin.HandlerFunc {
	return func(c *gin.Context) {
		email := db.GetUser().Email
		c.HTML(http.StatusOK, "index.tmpl", gin.H{
			"title": "Hello Web " + email,
		})
	}
}
