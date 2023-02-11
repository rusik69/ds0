package server

import (
	"net/http"

	"github.com/rusik69/ds0/pkg/node/env"
	"github.com/sirupsen/logrus"
)

// Serve serves the node.
func Serve() {
	http.Handle("/", http.FileServer(http.Dir(env.NodeEnvInstance.Dir)))
	logrus.Infof("serving node %s at port %s", env.NodeEnvInstance.Name, env.NodeEnvInstance.Port)
	http.ListenAndServe(":"+env.NodeEnvInstance.Port, nil)
}
