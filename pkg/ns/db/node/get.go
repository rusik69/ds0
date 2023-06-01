package node

import (
	"context"
	"encoding/json"
	"errors"
	"time"

	"github.com/rusik69/ds0/pkg/ns/db"
	"github.com/rusik69/ds0/pkg/ns/env"
	"github.com/sirupsen/logrus"
)

// Get returns the node from the database.
func Get(nodeName string) (string, string, error) {
	logrus.Println("Get node: " + nodeName)
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	resp, err := db.DBNodes.Get(ctx, "/nodes/"+nodeName)
	cancel()
	if err != nil {
		logrus.Error(err)
	}
	if len(resp.Kvs) == 0 {
		return "", "", errors.New("node not found")
	}
	var s env.NodeInfo
	json.Unmarshal(resp.Kvs[0].Value, &s)
	return s.Host, s.Port, nil
}
