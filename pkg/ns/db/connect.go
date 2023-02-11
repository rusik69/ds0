package db

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	"github.com/rusik69/ds0/pkg/ns/env"
	"github.com/sirupsen/logrus"
)

// Connect connects to the database.
func Connect() error {
	var psqlInfo string
	if env.NodeEnvInstance.DBPassword == "" {
		psqlInfo = fmt.Sprintf("host=%s port=%s user=%s "+
			"dbname=%s sslmode=%s connect_timeout=60",
			env.NodeEnvInstance.DBHost, env.NodeEnvInstance.DBPort, env.NodeEnvInstance.DBUser,
			env.NodeEnvInstance.DBName, env.NodeEnvInstance.DBSSLMode)
	} else {
		psqlInfo = fmt.Sprintf("host=%s port=%s user=%s "+
			"password=%s dbname=%s sslmode=%s connect_timeout=60",
			env.NodeEnvInstance.DBHost, env.NodeEnvInstance.DBPort, env.NodeEnvInstance.DBUser,
			env.NodeEnvInstance.DBPassword, env.NodeEnvInstance.DBName, env.NodeEnvInstance.DBSSLMode)
	}
	logrus.Println(psqlInfo)
	success := false
	var db *sql.DB
	var err error
	for i := 1; i <= 10; i++ {
		db, err = sql.Open("postgres", psqlInfo)
		if err != nil {
			logrus.Println("Error opening database connection, retrying")
			time.Sleep(5 * time.Second)
			continue
		}
		ctx, cancelfunc := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancelfunc()
		err = db.PingContext(ctx)
		if err != nil {
			logrus.Println("Error pinging database, retrying")
			time.Sleep(1 * time.Second)
			continue
		}
		success = true
	}
	if !success {
		fmt.Println("Reached max retries, exiting")
		return err
	}
	DB = db
	return nil
}
