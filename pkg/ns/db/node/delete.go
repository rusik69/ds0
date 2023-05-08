package node

import (
	"context"
	"time"

	"github.com/rusik69/ds0/pkg/ns/db"
	"github.com/sirupsen/logrus"
)

// Delete removes the node from the database.
func Delete(nodeName string) error {
	logrus.Println("Delete node: " + nodeName)
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	_, err := db.DBNodes.Delete(ctx, nodeName)
	cancel()
	if err != nil {
		return err
	}
	return nil
}
