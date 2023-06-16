package web

import (
	"net/http"
	"text/template"

	"github.com/gin-gonic/gin"
	clientnode "github.com/rusik69/ds0/pkg/client/node"
	"github.com/rusik69/ds0/pkg/ns/db/node"
	"github.com/sirupsen/logrus"
)

// RootHandler is the root handler.
func RootHandler(c *gin.Context) {
	logrus.Println("RootHandler")
	nodes, err := node.List()
	if err != nil {
		c.Writer.WriteHeader(500)
		c.Writer.Write([]byte(err.Error()))
		logrus.Error(err)
		return
	}
	var totalSpace, totalFreeSpace, totalUsedSpace uint64
	for _, node := range nodes {
		stats, err := clientnode.Stats(node.Host, node.Port)
		if err != nil {
			logrus.Error(err)
			continue
		}
		node.Stats.TotalSpace = stats.TotalSpace
		node.Stats.FreeSpace = stats.FreeSpace
		node.Stats.UsedSpace = stats.UsedSpace
		totalSpace += stats.TotalSpace
		totalFreeSpace += stats.FreeSpace
		totalUsedSpace += stats.UsedSpace
	}
	data := gin.H{
		"Title":          "DS0",
		"Nodes":          nodes,
		"TotalSpace":     totalSpace,
		"TotalFreeSpace": totalFreeSpace,
		"TotalUsedSpace": totalUsedSpace,
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
		c.Writer.WriteHeader(500)
		c.Writer.Write([]byte(err.Error()))
		logrus.Error(err)
		return
	}
	c.Writer.WriteHeader(http.StatusOK)
}
