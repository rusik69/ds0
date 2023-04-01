package node

import (
	"context"
	"encoding/json"
	"time"

	"github.com/rusik69/ds0/pkg/ns/db"
	clientv3 "go.etcd.io/etcd/client/v3"
)

// List returns the list of nodes in the database.
func List() ([]db.HostInfo, error) {
	var nodes []db.HostInfo
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	resp, err := db.DB.Get(ctx, "/nodes/", clientv3.WithPrefix())
	cancel()
	if err != nil {
		return nodes, err
	}
	for _, ev := range resp.Kvs {
		var s db.HostInfo
		json.Unmarshal(ev.Value, &s)
		nodes = append(nodes, s)
	}
	return nodes, nil
}