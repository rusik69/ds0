package server

import (
	"errors"
	"net/http"
	"os"
	"path/filepath"

	"github.com/gin-gonic/gin"
	"github.com/rusik69/ds0/pkg/node/env"
	"github.com/rusik69/ds0/pkg/node/metrics"
	"github.com/sirupsen/logrus"
)

// deleteHandler is the handler for deleting a file.
func deleteHandler(c *gin.Context) {
	metrics.Counter.Inc()
	fileName := c.Query("file")
	if fileName == "" {
		c.Writer.WriteHeader(400)
		c.Writer.Write([]byte("file name is required"))
		logrus.Error(errors.New("file name is required"))
		return
	}
	fileName = filepath.Join(env.NodeEnvInstance.Dir, fileName)
	err := os.Remove(fileName)
	if err != nil {
		c.String(http.StatusInternalServerError, err.Error())
		logrus.Error(err.Error())
		return
	}
	c.Writer.WriteHeader(http.StatusOK)
}
