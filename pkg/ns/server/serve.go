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
	router.GET("/node/get", node.GetNodeHandler)
	router.GET("/node/add", node.AddNodeHandler)
	router.GET("/node/remove", node.RemoveNodeHandler)
	router.GET("/file/upload", file.UploadHandler)
	router.GET("/file/commit", file.CommitHandler)
	router.GET("/file/download", file.DownloadHandler)
	router.Run(":" + string(env.NSEnvInstance.Port))
}
