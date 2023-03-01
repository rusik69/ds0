package node

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	"github.com/rusik69/ds0/pkg/ns/db"
	"github.com/rusik69/ds0/pkg/ns/env"
)

// Get returns the node from the database.
func Get(nodeName string) (string, string, error) {
	var hostname, port string
	ctx, cancelfunc := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancelfunc()
	sqlStmt := fmt.Sprintf("SELECT host, port FROM %s WHERE name=$1;", env.NSEnvInstance.DBNodesTableName)
	err := db.DB.QueryRowContext(ctx, sqlStmt, nodeName).Scan(&hostname, &port)
	if err != nil {
		if err != sql.ErrNoRows {
			return "", "", err
		}
		return "", "", nil
	}
	return hostname, port, nil
}
