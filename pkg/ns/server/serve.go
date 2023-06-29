package server

import (
	"github.com/gin-gonic/gin"
	"github.com/rusik69/ds0/pkg/ns/env"
	"github.com/rusik69/ds0/pkg/ns/server/cluster"
	"github.com/rusik69/ds0/pkg/ns/server/file"
	"github.com/rusik69/ds0/pkg/ns/server/node"
	"github.com/rusik69/ds0/pkg/ns/web"
	"github.com/sirupsen/logrus"
)

// Serve serves the ns.
func Serve() {
	router := gin.Default()
	router.LoadHTMLGlob("/app/html/*.html")
	router.GET("/node/get", node.GetNodeHandler)
	router.GET("/node/add", node.AddNodeHandler)
	router.GET("/node/remove", node.RemoveNodeHandler)
	router.GET("/node/list", node.ListNodesHandler)
	router.GET("/cluster/stats", cluster.StatsHandler)
	router.GET("/file/upload", file.UploadHandler)
	router.GET("/file/commit", file.CommitHandler)
	router.GET("/file/download", file.DownloadHandler)
	router.GET("/file/delete", file.DeleteHandler)
	router.GET("/ping", func(c *gin.Context) {
		c.String(200, "pong")
	})
	router.Static("/static", "/app/html/static")
	router.GET("/", web.RootHandler)
	logrus.Println("NS is listening on port " + string(env.NSEnvInstance.Port))
	router.Run(":" + string(env.NSEnvInstance.Port))
}
