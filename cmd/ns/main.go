package main

import (
	"net"

	"github.com/gin-gonic/gin"
	"github.com/rusik69/ds0/pkg/ns/db"
	"github.com/rusik69/ds0/pkg/ns/db/node"
	"github.com/rusik69/ds0/pkg/ns/env"
	"github.com/rusik69/ds0/pkg/ns/server"
	"github.com/rusik69/ds0/pkg/ns/watcher"
	"github.com/sirupsen/logrus"
)

var version string

func main() {
	logger := logrus.New()
	logger.SetReportCaller(true)
	logrus.SetFormatter(&logrus.JSONFormatter{})
	gin.DefaultWriter = logger.Writer()
	gin.DefaultErrorWriter = logger.Writer()
	logrus.Println("Version: ", version)
	err := env.Parse()
	if err != nil {
		panic(err)
	}
	db.DB, err = db.Connect(
		env.NSEnvInstance.ETCDHost,
		env.NSEnvInstance.ETCDPort,
		env.NSEnvInstance.ETCDUser,
		env.NSEnvInstance.ETCDPass)
	if err != nil {
		panic(err)
	}
	defer db.DB.Close()
	logrus.Println("Connected to ETCD: " + env.NSEnvInstance.ETCDHost + ":" + env.NSEnvInstance.ETCDPort)
	if env.NSEnvInstance.NodesStatefulSetName != "" {
		ips, err := net.LookupIP(env.NSEnvInstance.NodesStatefulSetName)
		if err != nil {
			panic(err)
		}
		if len(ips) == 0 {
			panic("No IP found in stateful set")
		}
		for _, ip := range ips {
			err = node.Add(ip.String(), ip.String(), env.NSEnvInstance.NodesStatefulSetPort)
			if err != nil {
				panic(err)
			}
		}
	}
	go watcher.Watch()
	server.Serve()
}
