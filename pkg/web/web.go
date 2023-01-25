package web

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func IndexHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.tmpl", gin.H{
			"title": "Hello Web",
		})
	}
}
