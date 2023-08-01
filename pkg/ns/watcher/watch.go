package watcher

import (
	"fmt"
	"time"

	nodeclient "github.com/rusik69/ds0/pkg/client/node"
	dbfile "github.com/rusik69/ds0/pkg/ns/db/file"
	"github.com/rusik69/ds0/pkg/ns/db/node"
	"github.com/sirupsen/logrus"
)

// Watch for nodes
func Watch() {
	for {
		time.Sleep(60 * time.Second)
		logrus.Println("Watch nodes")
		nodes, err := node.List()
		if err != nil {
			logrus.Error(err)
			continue
		}
		for _, node := range nodes {
			logrus.Println("Watch node: " + node.Host + ":" + node.Port)
			nodeStats, err := nodeclient.Stats(node.Host, node.Port)
			if err != nil {
				logrus.Error(err)
				continue
			}
			logrus.Println("Total space: " + fmt.Sprintf("%d", nodeStats.TotalSpace))
			logrus.Println("Free space: " + fmt.Sprintf("%d", nodeStats.FreeSpace))
			logrus.Println("Used space: " + fmt.Sprintf("%d", nodeStats.UsedSpace))
		}
		uncommittedFiles, err := dbfile.ListUncommitted()
		if err != nil {
			logrus.Error(err)
			continue
		}
		for fileName, fileInfo := range uncommittedFiles {
			if fileInfo.Committed == false && time.Since(fileInfo.TimeAdded) >= time.Hour {

			}
		}
	}
}
