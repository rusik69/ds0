package server

import (
	"github.com/gin-gonic/gin"
	"github.com/rusik69/ds0/pkg/web/env"
	"github.com/rusik69/ds0/pkg/web/metrics"
	"github.com/rusik69/ds0/pkg/web/server/files"
)

// Serve serves the web server.
func Serve() {
	router := gin.New()
	router.LoadHTMLGlob("/app/html/*.html")
	router.GET("/ping", func(c *gin.Context) {
		metrics.Counter.Inc()
		c.String(200, "pong")
	})
	router.Static("/static", "/app/html/static")
	router.GET("/", rootHandler)
	router.GET("/api/v0/file/list", files.ListHandler)
	router.GET("/api/v0/file/upload", files.UploadHandler)
	router.GET("/metrics", metrics.PrometheusHandler())
	router.Run(":" + string(env.EnvInstance.ListenPort))
}
