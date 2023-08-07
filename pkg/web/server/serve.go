package server

import (
	"github.com/gin-gonic/gin"
	"github.com/rusik69/ds0/pkg/web/env"
)

// Serve serves the web server.
func Serve() {
	router := gin.New()
	router.LoadHTMLGlob("/app/html/*.html")
	router.GET("/ping", func(c *gin.Context) {
		c.String(200, "pong")
	})
	router.Static("/static", "/app/html/static")
	router.GET("/", rootHandler)
	router.Run(":" + string(env.EnvInstance.ListenPort))
}
