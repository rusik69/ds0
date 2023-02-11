package env

// NodeEnv is the environment of a node.
type NodeEnv struct {
	Name        string
	DBHost      string
	DBPort      string
	DBUser      string
	DBPassword  string
	DBName      string
	DBSSLMode   string
	DBTableName string
	Port        string
}

// NodeEnvInstance is the singleton instance of NodeEnv.
var NodeEnvInstance *NodeEnv
