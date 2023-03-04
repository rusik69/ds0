package db

import (
	clientv3 "go.etcd.io/etcd/client/v3"
)

// DB is the database connection.
var DB *clientv3.Client

// HostInfo is the host information.
type HostInfo struct {
	Host string
	Port string
}
