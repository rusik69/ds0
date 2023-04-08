package env

import (
	"errors"
	"os"
	"strconv"
	"strings"

	"github.com/sirupsen/logrus"
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
	etcdHost := os.Getenv("NS_ETCD_HOST_FILES")
	if etcdHost == "" {
		etcdHost = "localhost"
	}
	etcdPort := os.Getenv("NS_ETCD_PORT_FILES")
	if etcdPort == "" {
		etcdPort = "2379"
	}
	etcdUser := os.Getenv("NS_ETCD_USER_FILES")
	if etcdUser == "" {
		etcdUser = ""
	}
	etcdPass := os.Getenv("NS_ETCD_PASS_FILES")
	if etcdPass == "" {
		etcdPass = ""
	}
	etcdHostNodes := os.Getenv("NS_ETCD_HOST_NODES")
	if etcdHost == "" {
		etcdHostNodes = "localhost"
	}
	etcdPortNodes := os.Getenv("NS_ETCD_PORT_NODES")
	if etcdPortNodes == "" {
		etcdPortNodes = "2379"
	}
	etcdUserNodes := os.Getenv("NS_ETCD_USER_NODES")
	if etcdUserNodes == "" {
		etcdUserNodes = ""
	}
	etcdPassNodes := os.Getenv("NS_ETCD_PASS_NODES")
	if etcdPassNodes == "" {
		etcdPassNodes = ""
	}
	replicas := os.Getenv("NS_REPLICAS")
	if replicas == "" {
		replicas = "3"
	}

	nodes := os.Getenv("NS_NODES")
	nodesMap := map[string]NodeInfo{}
	// name:hostname:port
	if nodes != "" {
		nodeList := strings.Split(nodes, ",")
		for _, node := range nodeList {
			nodeInfo := strings.Split(node, ":")
			if len(nodeInfo) != 3 {
				return errors.New("NS_NODES is not set correctly")
			}
			nodesMap[nodeInfo[0]] = NodeInfo{HostName: nodeInfo[1], Port: nodeInfo[2]}
		}
	}
	repl, err := strconv.Atoi(replicas)
	if err != nil {
		repl = 3
	}

	NSEnvInstance = &NSEnv{
		Name:          name,
		Port:          port,
		ETCDHost:      etcdHost,
		ETCDPort:      etcdPort,
		ETCDUser:      etcdUser,
		ETCDPass:      etcdPass,
		ETCDHostNodes: etcdHostNodes,
		ETCDPortNodes: etcdPortNodes,
		ETCDUserNodes: etcdUserNodes,
		ETCDPassNodes: etcdPassNodes,
		Replicas:      repl,
		Nodes:         nodesMap,
	}

	logrus.Println("node name: ", name)
	logrus.Println("node port: ", port)
	logrus.Println("etcd host: ", etcdHost)
	logrus.Println("etcd port: ", etcdPort)
	logrus.Println("etcd user: ", etcdUser)
	logrus.Println("etcd pass: ", etcdPass)
	logrus.Println("replicas: ", repl)
	logrus.Println("nodes: ", nodesMap)

	return nil
}
