package watcher

import (
	"time"

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
		}
	}
}
