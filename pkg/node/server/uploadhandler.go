package server

import (
	"io"
	"net/http"
	"os"
	"path/filepath"

	"github.com/gin-gonic/gin"
	"github.com/rusik69/ds0/pkg/node/env"
)

// uploadHandler handles file upload.
func uploadHandler(c *gin.Context) {
	fileName := c.Request.URL.Path
	if fileName == "" {
		c.String(http.StatusBadRequest, "no filename provided")
		return
	}
	file, _, err := c.Request.FormFile("file")
	if err != nil {
		c.String(http.StatusBadRequest, "Bad request")
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
