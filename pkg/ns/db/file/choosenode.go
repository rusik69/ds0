package file

import (
	"math/rand"
	"time"

	"github.com/rusik69/ds0/pkg/ns/db"
	"github.com/rusik69/ds0/pkg/ns/env"
)

// ChooseNodes chooses nodes to upload a file.
func ChooseNodes(nodes []db.HostInfo) []db.HostInfo {
	rand.Seed(time.Now().UnixNano())
	rand.Shuffle(len(nodes), func(i, j int) {
		nodes[i], nodes[j] = nodes[j], nodes[i]
	})
	return nodes[:env.NSEnvInstance.Replicas]
}