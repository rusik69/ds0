package file

import (
	"errors"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/rusik69/ds0/pkg/ns/db"
	"github.com/rusik69/ds0/pkg/ns/db/file"
	dbfile "github.com/rusik69/ds0/pkg/ns/db/file"
	dbnode "github.com/rusik69/ds0/pkg/ns/db/node"
	"github.com/rusik69/ds0/pkg/ns/env"
	"github.com/sirupsen/logrus"
)

// UploadHandler is the handler for uploading a file.
func UploadHandler(c *gin.Context) {
	fileName := c.Query("file")
	if fileName == "" {
		c.Writer.WriteHeader(400)
		c.Writer.Write([]byte("file name is required"))
		logrus.Error(errors.New("file name is required"))
		return
	}
	logrus.Println("UploadHandler: " + fileName)
	fileInfo, err := dbfile.Get(fileName)
	if err == os.ErrNotExist {
		newNodes, err := dbnode.List()
		if err != nil {
			c.Writer.WriteHeader(500)
			c.Writer.Write([]byte(err.Error()))
			logrus.Error(err)
			return
		}
		if len(newNodes) == 0 {
			c.Writer.WriteHeader(500)
			c.Writer.Write([]byte("no nodes available"))
			logrus.Error(errors.New("no nodes available"))
			logrus.Error(env.NSEnvInstance.Nodes)
			return
		} else if len(newNodes) < env.NSEnvInstance.Replicas {
			c.Writer.WriteHeader(500)
			c.Writer.Write([]byte("not enough nodes available"))
			logrus.Error(errors.New("not enough nodes available"))
			return
		}
		nodes := file.ChooseNodes(newNodes)
		err = dbfile.Set(fileName, db.FileInfo{Nodes: nodes})
		if err != nil {
			c.Writer.WriteHeader(500)
			c.Writer.Write([]byte(err.Error()))
			logrus.Error(err)
			return
		}
		c.JSON(http.StatusOK, nodes)
	} else if err != nil {
		c.Writer.WriteHeader(500)
		c.Writer.Write([]byte(err.Error()))
		logrus.Error(err)
		return
	}
	c.JSON(http.StatusOK, fileInfo.Nodes)
}
