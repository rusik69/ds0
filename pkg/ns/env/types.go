package env

// NSEnv is the environment of a node.
type NSEnv struct {
	Name                 string
	ETCDHostFiles        string
	ETCDPortFiles        string
	ETCDUserFiles        string
	ETCDPassFiles        string
	ETCDHostNodes        string
	ETCDPortNodes        string
	ETCDUserNodes        string
	ETCDPassNodes        string
	Port                 string
	Replicas             int
	Nodes                map[string]NodeInfo
	NodesStatefulSetName string
	NodesStatefulSetPort string
}

// NodeEnvInstance is the singleton instance of NodeEnv.
var NSEnvInstance *NSEnv

// NodeInfo is the information of a node.
type NodeInfo struct {
	HostName string
	Port     string
}
