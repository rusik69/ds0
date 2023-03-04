package node

import (
	"context"
	"time"

	"github.com/rusik69/ds0/pkg/ns/db"
)

// Delete removes the node from the database.
func Delete(nodeName string) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	_, err := db.DB.Delete(ctx, nodeName)
	cancel()
	if err != nil {
		return err
	}
	return nil
}
