package server

import (
	"errors"
	"io"
	"net/http"
	"os"
	"path/filepath"

	"github.com/gin-gonic/gin"
	"github.com/rusik69/ds0/pkg/node/env"
	"github.com/rusik69/ds0/pkg/node/metrics"
	"github.com/sirupsen/logrus"
)

// downloadHandler is the handler for downloading a file.
func downloadHandler(c *gin.Context) {
	metrics.Counter.Inc()
	fileName := c.Query("file")
	if fileName == "" {
		c.Writer.WriteHeader(400)
		c.Writer.Write([]byte("file name is required"))
		logrus.Error(errors.New("file name is required"))
		return
	}
	fileName = filepath.Join(env.NodeEnvInstance.Dir, fileName)
	file, err := os.Open(fileName)
	if err != nil {
		c.String(http.StatusNotFound, err.Error())
		logrus.Error(err.Error())
		return
	}
	defer file.Close()
	c.Writer.WriteHeader(http.StatusOK)
	_, err = io.Copy(c.Writer, file)
	if err != nil {
		c.String(http.StatusInternalServerError, err.Error())
		logrus.Error(err.Error())
		return
	}
}
