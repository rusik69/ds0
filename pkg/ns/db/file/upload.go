package file

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	"github.com/rusik69/ds0/pkg/ns/db"
	"github.com/rusik69/ds0/pkg/ns/env"
)

// Upload writes file info to the database.
func Upload(fileName string) (string, error) {
	ctx, cancelfunc := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancelfunc()
	var nodes string
	nodesList := [][]string{}
	var found bool
	sqlStmt := fmt.Sprintf("SELECT nodes FROM %s WHERE name=$1;", env.NSEnvInstance.DBFilesTableName)
	err := db.DB.QueryRowContext(ctx, sqlStmt, fileName).Scan(&nodes)
	if err != nil {
		if err != sql.ErrNoRows {
			return "", err
		} else {
			found = true
		}
	}
	if !found {
		sqlStmt = fmt.Sprintf("SELECT name, host, port FROM %s;", env.NSEnvInstance.DBNodesTableName)
		rows, err := db.DB.QueryContext(ctx, sqlStmt)
		if err != nil {
			return "", err
		}
		for rows.Next() {
			var name, host, port string
			if err := rows.Scan(&name, &host, &port); err != nil {
				panic(err)
			}
			nodesList = append(nodesList, []string{name, host, port})
		}
	}
	return nil
}
