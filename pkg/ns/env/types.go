package env

import "github.com/rusik69/ds0/pkg/node/server"

// NSEnv is the environment of a node.
type NSEnv struct {
	Name                 string
	ETCDHost             string
	ETCDPort             string
	ETCDUser             string
	ETCDPass             string
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

// NodeInfo is the host information.
type NodeInfo struct {
	Host  string
	Port  string
	Stats server.NodeStats
}
