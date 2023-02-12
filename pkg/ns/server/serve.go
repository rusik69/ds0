package server

import (
	"github.com/gin-gonic/gin"
	"github.com/rusik69/ds0/pkg/ns/env"
	"github.com/rusik69/ds0/pkg/ns/server/node"
)

// Serve serves the ns.
func Serve() {
	router := gin.Default()
	router.GET("/node", node.GetNode)
	router.Run(":" + string(env.NSEnvInstance.Port))
}
