package env

import (
	"errors"
	"os"
	"strconv"
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
	etcdHost := os.Getenv("NS_ETCD_HOST")
	if etcdHost == "" {
		etcdHost = "localhost"
	}
	etcdPort := os.Getenv("NS_ETCD_PORT")
	if etcdPort == "" {
		etcdPort = "2379"
	}
	replicas := os.Getenv("NS_REPLICAS")
	if replicas == "" {
		replicas = "3"
	}
	repl, err := strconv.Atoi(replicas)
	if err != nil {
		repl = 3
	}
	NSEnvInstance = &NSEnv{
		Name:     name,
		Port:     port,
		ETCDHost: etcdHost,
		ETCDPort: etcdPort,
		Replicas: repl,
	}
	return nil
}
