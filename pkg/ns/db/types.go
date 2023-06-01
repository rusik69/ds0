package db

import (
	"github.com/rusik69/ds0/pkg/ns/env"
	clientv3 "go.etcd.io/etcd/client/v3"
)

// DBFiles is the database connection.
var DBFiles, DBNodes *clientv3.Client

// FileInfo is the file information.
type FileInfo struct {
	Nodes     []env.NodeInfo
	Committed bool
}
