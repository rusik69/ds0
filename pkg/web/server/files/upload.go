package files

import (
	"io"
	"os"

	"github.com/gin-gonic/gin"
	fileclient "github.com/rusik69/ds0/pkg/client/file"
	"github.com/rusik69/ds0/pkg/ns/server/utils"
	"github.com/rusik69/ds0/pkg/web/env"
	"github.com/rusik69/ds0/pkg/web/metrics"
	"github.com/sirupsen/logrus"
)

// UploadHandler handles the upload file request.
func UploadHandler(c *gin.Context) {
	metrics.Counter.Inc()
	file, err := c.FormFile("file")
	if err != nil {
		c.String(400, "Bad request")
		return
	}
	fileName := file.Filename
	logrus.Println("Upload file: " + fileName)
	src, err := file.Open()
	if err != nil {
		utils.Error(err.Error(), 500, c)
		return
	}
	defer src.Close()
	tempFile, err := os.CreateTemp("", fileName)
	if err != nil {
		utils.Error(err.Error(), 500, c)
		return
	}
	defer tempFile.Close()
	_, err = io.Copy(tempFile, src)
	if err != nil {
		utils.Error(err.Error(), 500, c)
		return
	}
	err = fileclient.Upload(tempFile.Name(), fileName, env.EnvInstance.NSHost, env.EnvInstance.NSPort)
	if err != nil {
		utils.Error(err.Error(), 500, c)
		return
	}
	c.String(200, "OK")
}
