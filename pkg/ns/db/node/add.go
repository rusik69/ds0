package node

import (
	"context"
	"fmt"
	"time"

	"github.com/rusik69/ds0/pkg/ns/db"
	"github.com/rusik69/ds0/pkg/ns/env"
)

// Add adds the node to the database.
func Add(name, hostname string, port int) error {
	ctx, cancelfunc := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancelfunc()
	sqlStmt := fmt.Sprintf("INSERT INTO %s(name, hostname, port) VALUES($1, $2, $3);", env.NSEnvInstance.DBNodesTableName)
	_, err := db.DB.ExecContext(ctx, sqlStmt, name, hostname, port)
	if err != nil {
		return err
	}
	return nil
}
