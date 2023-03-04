package node

import (
	"context"
	"encoding/json"
	"log"
	"time"

	"github.com/rusik69/ds0/pkg/ns/db"
)

// Add adds the node to the database.
func Add(name, hostname, port string) error {
	info :=  db.HostInfo{
		Host: hostname,
		Port:     port,
	}
	str, err := json.Marshal(info)
	if err != nil {
		log.Fatal(err)
		return err
	}
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	_, err = db.DB.Put(ctx, name, string(str))
	cancel()
	if err != nil {
		log.Fatal(err)
		return err
	}
	return nil
}
