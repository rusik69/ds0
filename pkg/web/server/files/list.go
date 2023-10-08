package files

import (
	"github.com/gin-gonic/gin"
	fileclient "github.com/rusik69/ds0/pkg/client/file"
	"github.com/rusik69/ds0/pkg/ns/server/utils"
	"github.com/rusik69/ds0/pkg/web/env"
	"github.com/rusik69/ds0/pkg/web/metrics"
	"github.com/sirupsen/logrus"
)

// ListHandler handles the list files request.
func ListHandler(c *gin.Context) {
	metrics.Counter.Inc()
	prefix := c.Query("prefix")
	if prefix == "" {
		prefix = "/"
	}
	logrus.Println("ListHandler on prefix ", prefix)
	files, err := fileclient.List(env.EnvInstance.NSHost, env.EnvInstance.NSPort, prefix)
	if err != nil {
		utils.Error(err.Error(), 500, c)
		return
	}
	c.JSON(200, files)
}
