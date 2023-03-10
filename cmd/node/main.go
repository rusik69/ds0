package node

import (
	"github.com/rusik69/ds0/pkg/node/env"
	"github.com/rusik69/ds0/pkg/node/server"
	"github.com/sirupsen/logrus"
)

func main() {
	logrus.SetReportCaller(true)
	nodeEnvInstance, err := env.Parse()
	if err != nil {
		panic(err)
	}
	env.NodeEnvInstance = nodeEnvInstance
	logrus.Infof("parsed node environment: %+v", env.NodeEnvInstance)
	server.Serve()
}
