package node

import (
	"context"
	"encoding/json"
	"log"
	"time"

	"github.com/rusik69/ds0/pkg/ns/db"
	"github.com/sirupsen/logrus"
)

// Add adds the node to the database.
func Add(name, hostname, port string) error {
	logrus.Println("Add node: " + hostname + ":" + port)
	info := db.HostInfo{
		Host: hostname,
		Port: port,
	}
	bytes, err := json.Marshal(info)
	if err != nil {
		log.Fatal(err)
		return err
	}
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	_, err = db.DBNodes.Put(ctx, "/nodes/"+name, string(bytes))
	cancel()
	if err != nil {
		logrus.Error(err)
		return err
	}
	return nil
}
