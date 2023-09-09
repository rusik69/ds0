package server

import (
	"net/http"
	"text/template"

	humanize "github.com/dustin/go-humanize"
	"github.com/gin-gonic/gin"
	clientcluster "github.com/rusik69/ds0/pkg/client/cluster"
	"github.com/rusik69/ds0/pkg/ns/server/utils"
	"github.com/rusik69/ds0/pkg/web/env"
	"github.com/rusik69/ds0/pkg/web/metrics"
	"github.com/sirupsen/logrus"
)

// RootHandler is the root handler.
func rootHandler(c *gin.Context) {
	logrus.Println("RootHandler")
	metrics.Counter.Inc()
	clusterStats, err := clientcluster.Stats(env.EnvInstance.NSHost, env.EnvInstance.NSPort)
	if err != nil {
		c.Writer.WriteHeader(500)
		c.Writer.Write([]byte(err.Error()))
		logrus.Error(err)
		return
	}
	var nodes []NodeStats
	for _, node := range clusterStats.Nodes {
		stats := NodeStats{
			Host:       node.Host,
			Port:       node.Port,
			TotalSpace: humanize.Bytes(node.TotalSpace),
			FreeSpace:  humanize.Bytes(node.FreeSpace),
			UsedSpace:  humanize.Bytes(node.UsedSpace),
		}
		nodes = append(nodes, stats)
	}
	data := gin.H{
		"Title":            "DS0",
		"Nodes":            nodes,
		"TotalSpace":       humanize.Bytes(clusterStats.TotalSpace),
		"TotalFreeSpace":   humanize.Bytes(clusterStats.TotalFreeSpace),
		"TotalUsedSpace":   humanize.Bytes(clusterStats.TotalUsedSpace),
		"TotalFiles":       clusterStats.TotalFiles,
		"TotalSize":        humanize.Bytes(clusterStats.TotalFilesSize),
		"UncommittedFiles": clusterStats.UncommittedFiles,
		"UncommittedSize":  humanize.Bytes(clusterStats.UncommittedFilesSize),
	}
	tmpl, err := template.ParseFiles("/app/html/index.html")
	if err != nil {
		c.Writer.WriteHeader(500)
		c.Writer.Write([]byte(err.Error()))
		logrus.Error(err)
		return
	}
	err = tmpl.Execute(c.Writer, data)
	if err != nil {
		utils.Error(err.Error(), 500, c)
		return
	}
	c.Writer.WriteHeader(http.StatusOK)
}
