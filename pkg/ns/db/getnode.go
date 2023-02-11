package db

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	"github.com/rusik69/urlshortener/pkg/env"
)

// GetNode returns the node from the database.
func GetNode(nodeName string) (string, string, error) {
	var hostname, port string
	ctx, cancelfunc := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancelfunc()
	sqlStmt := fmt.Sprintf("SELECT hostname, port FROM %s WHERE name=$1;", env.ConfigInstance.DBTableName)
	err := DB.QueryRowContext(ctx, sqlStmt, nodeName).Scan(&hostname, &port)
	if err != nil {
		if err != sql.ErrNoRows {
			return "", "", err
		}
		return "", "", nil
	}
	return hostname, port, nil
}
