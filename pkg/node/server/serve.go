package server

import (
	"github.com/gin-gonic/gin"
	"github.com/rusik69/ds0/pkg/node/env"
)

// Serve serves the node.
func Serve() {
	r := gin.Default()
	r.POST("/", uploadHandler)
	r.GET("/", downloadHandler)
	r.Run(":" + string(env.NodeEnvInstance.Port))
}
