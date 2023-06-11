package node

import (
	"context"
	"encoding/json"
	"time"

	"github.com/rusik69/ds0/pkg/ns/db"
	"github.com/rusik69/ds0/pkg/ns/env"
	"github.com/sirupsen/logrus"
	clientv3 "go.etcd.io/etcd/client/v3"
)

// List returns the list of nodes in the database.
func List() ([]env.NodeInfo, error) {
	logrus.Println("List nodes")
	var nodes []env.NodeInfo
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	resp, err := db.DB.Get(ctx, "/nodes/", clientv3.WithPrefix())
	cancel()
	if err != nil {
		return nodes, err
	}
	for _, ev := range resp.Kvs {
		var s env.NodeInfo
		json.Unmarshal(ev.Value, &s)
		nodes = append(nodes, s)
	}
	return nodes, nil
}
