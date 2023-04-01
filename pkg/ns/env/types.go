package env

// NSEnv is the environment of a node.
type NSEnv struct {
	Name     string
	ETCDHost string
	ETCDPort string
	Port     string
	Replicas int
}

// NodeEnvInstance is the singleton instance of NodeEnv.
var NSEnvInstance *NSEnv
