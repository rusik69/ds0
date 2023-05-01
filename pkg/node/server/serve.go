package server

import (
	"github.com/gin-gonic/gin"
	"github.com/rusik69/ds0/pkg/node/env"
)

// Serve serves the node.
func Serve() {
	r := gin.Default()
	r.POST("/file/upload", uploadHandler)
	r.GET("/file/download*", downloadHandler)
	r.GET("/ping", func(c *gin.Context) {
		c.String(200, "pong")
	})
	r.Run(":" + string(env.NodeEnvInstance.Port))
}
