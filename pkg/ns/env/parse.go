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
	etcdHost := os.Getenv("NS_ETCD_HOST")
	if etcdHost == "" {
		etcdHost = "localhost"
	}
	etcdPort := os.Getenv("NS_ETCD_PORT")
	if etcdPort == "" {
		etcdPort = "2379"
	}
	NSEnvInstance = &NodeEnv{
		Name:   name,
		Port:   port,
		ETCDHost: etcdHost,
		ETCDPort: etcdPort,
	}
	return nil
}
