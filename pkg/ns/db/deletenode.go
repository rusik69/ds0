package db

import (
	"context"
	"fmt"
	"time"

	"github.com/rusik69/urlshortener/pkg/env"
)

// DeleteNode returns the node from the database.
func DeleteNode(nodeName string) error {
	ctx, cancelfunc := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancelfunc()
	sqlStmt := fmt.Sprintf("DELETE FROM %s WHERE name=$1;", env.ConfigInstance.DBTableName)
	stmt, err := DB.PrepareContext(ctx, sqlStmt)
	if err != nil {
		return err
	}
	result, err := stmt.ExecContext(ctx, nodeName)
	if err != nil {
		return err
	}
	count, err := result.RowsAffected()
	if err != nil {
		return err
	}
	if count == 0 {
		return fmt.Errorf("no rows affected")
	}
	return nil
}
