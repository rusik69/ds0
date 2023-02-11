package env

// NodeEnv is the environment of a node.
type NodeEnv struct {
	Name string
	Dir  string
	Port string
}

// NodeEnvInstance is the singleton instance of NodeEnv.
var NodeEnvInstance *NodeEnv
