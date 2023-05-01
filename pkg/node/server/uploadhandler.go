package server

import (
	"errors"
	"io"
	"net/http"
	"os"
	"path/filepath"

	"github.com/gin-gonic/gin"
	"github.com/rusik69/ds0/pkg/node/env"
	"github.com/sirupsen/logrus"
)

// uploadHandler handles file upload.
func uploadHandler(c *gin.Context) {
	fileName := c.Query("file")
	if fileName == "" {
		c.Writer.WriteHeader(400)
		c.Writer.Write([]byte("file name is required"))
		logrus.Error(errors.New("file name is required"))
		return
	}
	file, _, err := c.Request.FormFile("file")
	if err != nil {
		c.String(http.StatusBadRequest, "bad request")
		logrus.Error(err)
		return
	}
	fileName = filepath.Join(env.NodeEnvInstance.Dir, fileName)
	out, err := os.Create(fileName)
	if err != nil {
		c.String(http.StatusInternalServerError, "Server error")
		return
	}
	defer out.Close()
	_, err = io.Copy(out, file)
	if err != nil {
		c.String(http.StatusInternalServerError, "Server error")
		return
	}
	c.String(http.StatusOK, "OK")
}
