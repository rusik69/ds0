package cluster

import (
	"github.com/rusik69/ds0/pkg/ns/env"
)

// ClusterStats is the info of a cluster.
type ClusterStats struct {
	TotalSpace     uint64
	TotalFreeSpace uint64
	TotalUsedSpace uint64
	NodesCount     int
	Nodes          []env.NodeInfo
	Replicas       int
}
