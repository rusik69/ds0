package web

import (
	"net/http"
	"text/template"

	humanize "github.com/dustin/go-humanize"
	"github.com/gin-gonic/gin"
	clientnode "github.com/rusik69/ds0/pkg/client/node"
	dbfile "github.com/rusik69/ds0/pkg/ns/db/file"
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
	for i := 0; i < len(nodes); i++ {
		stats, err := clientnode.Stats(nodes[i].Host, nodes[i].Port)
		if err != nil {
			logrus.Error(err)
			continue
		}
		nodes[i].Stats.TotalSpace = humanize.Bytes(stats.TotalSpace)
		nodes[i].Stats.FreeSpace = humanize.Bytes(stats.FreeSpace)
		nodes[i].Stats.UsedSpace = humanize.Bytes(stats.UsedSpace)
		totalSpace += stats.TotalSpace
		totalFreeSpace += stats.FreeSpace
		totalUsedSpace += stats.UsedSpace
	}
	filesInfo, err := dbfile.GetFilesInfo()
	if err != nil {
		c.Writer.WriteHeader(500)
		c.Writer.Write([]byte(err.Error()))
		logrus.Error(err)
		return
	}
	logrus.Println("Total space: " + humanize.Bytes(totalSpace))
	logrus.Println("Total free space: " + humanize.Bytes(totalFreeSpace))
	logrus.Println("Total used space: " + humanize.Bytes(totalUsedSpace))
	logrus.Println("Total files: " + humanize.Comma(int64(filesInfo.TotalFiles)))
	logrus.Println("Total size: " + humanize.Bytes(filesInfo.TotalSize))
	logrus.Println("Uncommitted files: " + humanize.Comma(int64(filesInfo.UncommittedFiles)))
	logrus.Println("Uncommitted size: " + humanize.Bytes(filesInfo.UncommittedSize))
	data := gin.H{
		"Title":            "DS0",
		"Nodes":            nodes,
		"TotalSpace":       humanize.Bytes(totalSpace),
		"TotalFreeSpace":   humanize.Bytes(totalFreeSpace),
		"TotalUsedSpace":   humanize.Bytes(totalUsedSpace),
		"TotalFiles":       filesInfo.TotalFiles,
		"TotalSize":        humanize.Bytes(filesInfo.TotalSize),
		"UncommittedFiles": filesInfo.UncommittedFiles,
		"UncommittedSize":  humanize.Bytes(filesInfo.UncommittedSize),
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
