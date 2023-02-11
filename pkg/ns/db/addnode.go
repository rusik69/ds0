package db

import (
	"context"
	"fmt"
	"time"

	"github.com/rusik69/urlshortener/pkg/env"
)

// AddNode adds the node to the database.
func AddNode(name, hostname string, port int) error {
	ctx, cancelfunc := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancelfunc()
	sqlStmt := fmt.Sprintf("INSERT INTO %s(name, hostname, port) VALUES($1, $2, $3);", env.ConfigInstance.DBTableName)
	_, err := DB.ExecContext(ctx, sqlStmt, name, hostname, port)
	if err != nil {
		return err
	}
	return nil
}
