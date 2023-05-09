package file

import (
	"encoding/json"
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	dbfile "github.com/rusik69/ds0/pkg/ns/db/file"
	"github.com/sirupsen/logrus"
)

// DownloadHandler is the download handler.
func DownloadHandler(c *gin.Context) {
	fileName := c.Query("file")
	if fileName == "" {
		c.Writer.WriteHeader(400)
		c.Writer.Write([]byte("file name is required"))
		logrus.Error(errors.New("file name is required"))
		return
	}
	logrus.Println("DownloadHandler: " + fileName)
	nodes, err := dbfile.Get(fileName)
	if err != nil {
		c.Writer.WriteHeader(500)
		c.Writer.Write([]byte(err.Error()))
		logrus.Error(err)
		return
	}
	s, err := json.Marshal(nodes)
	if err != nil {
		c.Writer.WriteHeader(500)
		c.Writer.Write([]byte(err.Error()))
		logrus.Error(err)
		return
	}
	c.Writer.WriteHeader(http.StatusOK)
	c.Writer.Write(s)
}
