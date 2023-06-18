package db

import (
	"github.com/rusik69/ds0/pkg/ns/env"
	clientv3 "go.etcd.io/etcd/client/v3"
)

// DB is the database connection.
var DB *clientv3.Client

// FileInfo is the file information.
type FileInfo struct {
	Nodes     []env.NodeInfo
	Committed bool
	TimeAdded uint64
	Size      uint64
}

// FilesInfo is the files information.
type FilesInfo struct {
	TotalFiles       uint64
	TotalSize        uint64
	UncommittedFiles uint64
	UncommittedSize  uint64
}
