package server

import (
	"io"
	"net/http"
	"os"
	"path/filepath"

	"github.com/gin-gonic/gin"
	"github.com/rusik69/ds0/pkg/node/env"
)

// downloadHandler is the handler for downloading a file.
func downloadHandler(c *gin.Context) {
	fileName := c.Request.URL.Path
	if fileName == "" {
		c.String(http.StatusBadRequest, "no filename provided")
		return
	}
	fileName = filepath.Join(env.NodeEnvInstance.Dir, fileName)
	file, err := os.Open(fileName)
	if err != nil {
		c.String(http.StatusNotFound, "File not found")
		return
	}
	defer file.Close()
	c.Writer.WriteHeader(http.StatusOK)
	_, err = io.Copy(c.Writer, file)
	if err != nil {
		c.String(http.StatusInternalServerError, "Server error")
		return
	}
}
