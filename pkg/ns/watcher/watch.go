package watcher

import (
	"time"

	fileclient "github.com/rusik69/ds0/pkg/client/file"
	nodeclient "github.com/rusik69/ds0/pkg/client/node"
	dbfile "github.com/rusik69/ds0/pkg/ns/db/file"
	"github.com/rusik69/ds0/pkg/ns/db/node"
	"github.com/sirupsen/logrus"
)

// Watch for nodes
func Watch() {
	for {
		time.Sleep(60 * time.Second)
		//logrus.Println("Watch nodes")
		nodes, err := node.List()
		if err != nil {
			logrus.Error(err)
			continue
		}
		for _, node := range nodes {
			//logrus.Println("Watch node: " + node.Host + ":" + node.Port)
			_, err := nodeclient.Stats(node.Host, node.Port)
			if err != nil {
				logrus.Error(err)
				continue
			}
		}
		uncommittedFiles, err := dbfile.ListUncommitted()
		if len(uncommittedFiles) == 0 {
			continue
		}
		if err != nil {
			logrus.Error(err)
			continue
		}
		for fileName, fileInfo := range uncommittedFiles {
			if !fileInfo.Committed && time.Since(time.Unix(int64(fileInfo.TimeAdded), 0)) >= time.Hour {
				logrus.Println("File " + fileName + " is uncommitted for more than 1 hour")
				for _, node := range fileInfo.Nodes {
					logrus.Println("Delete file " + fileName + " from node " + node.Host + ":" + node.Port)
					err = fileclient.Delete(node.Host, node.Port, fileName)
					if err != nil {
						logrus.Error(err)
						continue
					}
				}
				err = dbfile.Delete(fileName)
				if err != nil {
					logrus.Error(err)
					continue
				}
			}
		}
	}
}
