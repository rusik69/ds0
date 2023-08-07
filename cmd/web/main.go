package main

import (
	"github.com/gin-gonic/gin"
	"github.com/rusik69/ds0/pkg/web/env"
	"github.com/sirupsen/logrus"
)

var version string

func main() {
	logrus.SetReportCaller(true)
	logrus.SetFormatter(&logrus.JSONFormatter{})
	gin.DefaultWriter = logrus.StandardLogger().Writer()
	gin.DefaultErrorWriter = logrus.StandardLogger().Writer()
	logrus.Println("Version: ", version)
	env.EnvInstance = env.Parse()
	logrus.Println("Starting web server on " + env.EnvInstance.ListenHost + ":" + env.EnvInstance.ListenPort)
	logrus.Println("Connected to NS on " + env.EnvInstance.NSHost + ":" + env.EnvInstance.NSPort)

}
