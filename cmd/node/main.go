package main

import (
	"github.com/rusik69/ds0/pkg/node/env"
	"github.com/rusik69/ds0/pkg/node/server"
	"github.com/sirupsen/logrus"
)

var version string

func main() {
	logrus.SetReportCaller(true)
	logrus.Println("Version: ", version)
	nodeEnvInstance, err := env.Parse()
	if err != nil {
		panic(err)
	}
	env.NodeEnvInstance = nodeEnvInstance
	logrus.Infof("parsed node environment: %+v", env.NodeEnvInstance)
	server.Serve()
}
