package server

import (
	"github.com/gin-gonic/gin"
	"github.com/rusik69/ds0/pkg/ns/env"
	"github.com/rusik69/ds0/pkg/ns/server/file"
	"github.com/rusik69/ds0/pkg/ns/server/node"
)

// Serve serves the ns.
func Serve() {
	router := gin.Default()
	router.GET("/get_node", node.GetNode)
	router.GET("/add_node", node.AddNode)
	router.GET("/remove_node", node.RemoveNode)
	router.GET("/upload", file.UploadHandler)
	router.GET("/download", file.DownloadHandler)
	router.Run(":" + string(env.NSEnvInstance.Port))
}
