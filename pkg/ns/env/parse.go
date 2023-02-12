package env

import (
	"errors"
	"os"
)

// Parse parses the environment of a node.
func Parse() error {
	// parse the environment
	name := os.Getenv("NS_NAME")
	if name == "" {
		return errors.New("NS_NAME is not set")
	}
	port := os.Getenv("NS_PORT")
	if port == "" {
		port = "6969"
	}
	dbHost := os.Getenv("NS_DB_HOST")
	if dbHost == "" {
		dbHost = "localhost"
	}
	dbPort := os.Getenv("NS_DB_PORT")
	if dbPort == "" {
		dbPort = "5432"
	}
	dbUser := os.Getenv("NS_DB_USER")
	if dbUser == "" {
		dbUser = "postgres"
	}
	dbPassword := os.Getenv("NS_DB_PASSWORD")
	if dbPassword == "" {
		dbPassword = "postgres"
	}
	dbName := os.Getenv("NS_DB_NAME")
	if dbName == "" {
		dbName = "ns"
	}
	dbSSLMode := os.Getenv("NS_DB_SSL_MODE")
	if dbSSLMode == "" {
		dbSSLMode = "disable"
	}
	dbTableName := os.Getenv("NS_DB_TABLE_NAME")
	if dbTableName == "" {
		dbTableName = "nodes"
	}
	NSEnvInstance = &NodeEnv{
		Name:        name,
		Port:        port,
		DBHost:      dbHost,
		DBPort:      dbPort,
		DBUser:      dbUser,
		DBPassword:  dbPassword,
		DBName:      dbName,
		DBSSLMode:   dbSSLMode,
		DBTableName: dbTableName,
	}
	return nil
}
