package db

import (
	"context"
	"fmt"
	"time"

	"github.com/rusik69/ds0/pkg/ns/env"
	"github.com/sirupsen/logrus"
)

// Init initializes the database
func Init() error {
	sqlStmtNodes := fmt.Sprintf("CREATE TABLE IF NOT EXISTS %s (name TEXT PRIMARY KEY, host TEXT, port INTEGER)", env.NSEnvInstance.DBNodesTableName)
	sqlStmtFiles := fmt.Sprintf("CREATE TABLE IF NOT EXISTS %s (name TEXT PRIMARY KEY, nodes TEXT)", env.NSEnvInstance.DBFilesTableName)
	logrus.Println(sqlStmtNodes)
	logrus.Println(sqlStmtFiles)
	ctx, cancelfunc := context.WithTimeout(context.Background(), 60*time.Second)
	defer cancelfunc()
	_, err := DB.ExecContext(ctx, sqlStmtNodes)
	if err != nil {
		return err
	}
	_, err = DB.ExecContext(ctx, sqlStmtFiles)
	if err != nil {
		return err
	}
	return nil
}
