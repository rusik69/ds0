package cluster

import (
	"github.com/rusik69/ds0/pkg/node/server"
)

// ClusterStats is the info of a cluster.
type ClusterStats struct {
	TotalSpace           uint64
	TotalFreeSpace       uint64
	TotalUsedSpace       uint64
	NodesCount           int
	Nodes                []server.NodeStats
	Replicas             int
	TotalFiles           uint64
	TotalFilesSize       uint64
	UncommittedFiles     uint64
	UncommittedFilesSize uint64
}
