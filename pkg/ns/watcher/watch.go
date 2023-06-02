package watcher

import (
	"time"

	nodeclient "github.com/rusik69/ds0/pkg/client/node"
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
			logrus.Println("Node: " + node.Host + ":" + node.Port)
			logrus.Println("Total space: " + string(nodeStats.TotalSpace))
			logrus.Println("Free space: " + string(nodeStats.FreeSpace))
			logrus.Println("Used space: " + string(nodeStats.UsedSpace))
		}
	}
}
