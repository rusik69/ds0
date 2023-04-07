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
	err = db.Connect()
	if err != nil {
		panic(err)
	}
	defer db.DB.Close()
	for name, hostInfo := range env.NSEnvInstance.Nodes {
		err = node.Add(name, hostInfo.HostName, hostInfo.Port)
		if err != nil {
			panic(err)
		}
	}
	server.Serve()
}
