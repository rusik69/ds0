package file

import (
	"math/rand"
	"time"

	"github.com/rusik69/ds0/pkg/ns/db"
	"github.com/rusik69/ds0/pkg/ns/env"
	"github.com/sirupsen/logrus"
)

// ChooseNodes chooses nodes to upload a file.
func ChooseNodes(nodes []db.HostInfo) []db.HostInfo {
	logrus.Println("ChooseNodes")
	rand.Seed(time.Now().UnixNano())
	rand.Shuffle(len(nodes), func(i, j int) {
		nodes[i], nodes[j] = nodes[j], nodes[i]
	})
	return nodes[:env.NSEnvInstance.Replicas]
}
