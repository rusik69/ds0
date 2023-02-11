package env

import (
	"errors"
	"os"
)

// Parse parses the environment of a node.
func Parse() (*NodeEnv, error) {
	// parse the environment
	name := os.Getenv("NODE_NAME")
	if name == "" {
		return nil, errors.New("NODE_NAME is not set")
	}
	dir := os.Getenv("NODE_DIR")
	if dir == "" {
		dir = "/mnt/data"
	}
	port := os.Getenv("NODE_PORT")
	if port == "" {
		port = "6969"
	}
	return &NodeEnv{
		Name: name,
		Dir:  dir,
		Port: port,
	}, nil
}
