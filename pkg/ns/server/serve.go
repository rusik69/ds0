package server

import (
	"github.com/gin-gonic/gin"
	"github.com/rusik69/ds0/pkg/ns/env"
	"github.com/rusik69/ds0/pkg/ns/metrics"
	"github.com/rusik69/ds0/pkg/ns/server/cluster"
	"github.com/rusik69/ds0/pkg/ns/server/file"
	"github.com/rusik69/ds0/pkg/ns/server/node"
	"github.com/sirupsen/logrus"
)

// Serve serves the ns.
func Serve() {
	router := gin.New()
	router.GET("/api/v0/node/get", node.GetNodeHandler)
	router.GET("/api/v0/node/add", node.AddNodeHandler)
	router.GET("/api/v0/node/remove", node.RemoveNodeHandler)
	router.GET("/api/v0/node/list", node.ListNodesHandler)
	router.GET("/api/v0/cluster/stats", cluster.StatsHandler)
	router.GET("/api/v0/file/upload", file.UploadHandler)
	router.GET("/api/v0/file/commit", file.CommitHandler)
	router.GET("/api/v0/file/download", file.DownloadHandler)
	router.GET("/api/v0/file/delete", file.DeleteHandler)
	router.GET("/ping", func(c *gin.Context) {
		metrics.Counter.Inc()
		c.String(200, "pong")
	})
	router.GET("/metrics", metrics.PrometheusHandler())
	logrus.Println("NS is listening on port " + string(env.NSEnvInstance.Port))
	router.Run(":" + string(env.NSEnvInstance.Port))
}
