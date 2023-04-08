package db

import (
	clientv3 "go.etcd.io/etcd/client/v3"
)

// DBFiles is the database connection.
var DBFiles, DBNodes *clientv3.Client

// HostInfo is the host information.
type HostInfo struct {
	Host string
	Port string
}

// FileInfo is the file information.
type FileInfo struct {
	Nodes     []HostInfo
	Committed bool
}
