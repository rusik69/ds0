package main

import (
	"github.com/rusik69/ds0/pkg/ns/db"
	"github.com/rusik69/ds0/pkg/ns/db/node"
	"github.com/rusik69/ds0/pkg/ns/env"
	"github.com/rusik69/ds0/pkg/ns/server"
	"github.com/sirupsen/logrus"
)

func main() {
	logrus.SetReportCaller(true)
	err := env.Parse()
	if err != nil {
		panic(err)
	}
	db.DBFiles, err = db.Connect(
		env.NSEnvInstance.ETCDHostFiles,
		env.NSEnvInstance.ETCDPortFiles,
		env.NSEnvInstance.ETCDUserFiles,
		env.NSEnvInstance.ETCDPassFiles)
	if err != nil {
		panic(err)
	}
	defer db.DBFiles.Close()
	db.DBNodes, err = db.Connect(
		env.NSEnvInstance.ETCDHostNodes,
		env.NSEnvInstance.ETCDPortNodes,
		env.NSEnvInstance.ETCDUserNodes,
		env.NSEnvInstance.ETCDPassNodes)
	if err != nil {
		panic(err)
	}
	defer db.DBNodes.Close()
	for name, hostInfo := range env.NSEnvInstance.Nodes {
		err = node.Add(name, hostInfo.HostName, hostInfo.Port)
		if err != nil {
			panic(err)
		}
	}
	server.Serve()
}
