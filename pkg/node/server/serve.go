package server

import (
	"github.com/gin-gonic/gin"
	"github.com/rusik69/ds0/pkg/node/env"
	"github.com/rusik69/ds0/pkg/node/metrics"
)

// Serve serves the node.
func Serve() {
	r := gin.New()
	r.POST("/file/upload", uploadHandler)
	r.GET("/file/download", downloadHandler)
	r.GET("/file/delete", deleteHandler)
	r.GET("/ping", func(c *gin.Context) {
		metrics.Counter.Inc()
		c.String(200, "pong")
	})
	r.GET("/stats", statsHandler)
	r.GET("/metrics", metrics.PrometheusHandler())
	r.Run(":" + string(env.NodeEnvInstance.Port))
}
