package ns

import (
	"github.com/rusik69/ds0/pkg/ns/db"
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
	logrus.Infof("parsed node environment: %+v", env.NodeEnvInstance)
	err = db.Connect()
	if err != nil {
		panic(err)
	}
	server.Serve()
}
