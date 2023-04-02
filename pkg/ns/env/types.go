package env

// NSEnv is the environment of a node.
type NSEnv struct {
	Name     string
	ETCDHost string
	ETCDPort string
	ETCDUser string
	ETCDPass string
	Port     string
	Replicas int
	Nodes    map[string]NodeInfo
}

// NodeEnvInstance is the singleton instance of NodeEnv.
var NSEnvInstance *NSEnv

// NodeInfo is the information of a node.
type NodeInfo struct {
	HostName string
	Port     string
}
