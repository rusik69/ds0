package env

// NodeEnv is the environment of a node.
type NodeEnv struct {
	Name     string
	ETCDHost string
	ETCDPort string
	Port     string
	Replicas int
}

// NodeEnvInstance is the singleton instance of NodeEnv.
var NSEnvInstance *NodeEnv
