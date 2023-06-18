package file

import (
	"encoding/json"
	"errors"
	"os"
	"strconv"
	"time"

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
	fileSize := c.Query("size")
	if fileSize == "" {
		c.Writer.WriteHeader(400)
		c.Writer.Write([]byte("file size is required"))
		logrus.Error(errors.New("file size is required"))
		return
	}
	size, error := strconv.ParseUint(fileSize, 10, 64)
	if error != nil {
		c.Writer.WriteHeader(400)
		c.Writer.Write([]byte("file size is invalid"))
		logrus.Error(errors.New("file size is invalid"))
		return
	}
	logrus.Println("UploadHandler: " + fileName + " size: " + fileSize)
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
		timestamp := uint64(time.Now().Unix())
		err = dbfile.Set(fileName, db.FileInfo{Nodes: nodes, TimeAdded: timestamp, Size: size})
		if err != nil {
			c.Writer.WriteHeader(500)
			c.Writer.Write([]byte(err.Error()))
			logrus.Error(err)
			return
		}
		body, err := json.Marshal(nodes)
		if err != nil {
			c.Writer.WriteHeader(500)
			c.Writer.Write([]byte(err.Error()))
			logrus.Error(err)
			return
		}
		c.Writer.Write(body)
		return
	} else if err != nil {
		c.Writer.WriteHeader(500)
		c.Writer.Write([]byte(err.Error()))
		logrus.Error(err)
		return
	}
	body, err := json.Marshal(fileInfo.Nodes)
	if err != nil {
		c.Writer.WriteHeader(500)
		c.Writer.Write([]byte(err.Error()))
		logrus.Error(err)
		return
	}
	filesInfo, err := file.GetFilesInfo()
	if err != nil {
		c.Writer.WriteHeader(500)
		c.Writer.Write([]byte(err.Error()))
		logrus.Error(err)
		return
	}
	filesInfo.UncommittedSize += fileInfo.Size
	filesInfo.UncommittedFiles++
	err = file.SetFilesInfo(filesInfo)
	if err != nil {
		c.Writer.WriteHeader(500)
		c.Writer.Write([]byte(err.Error()))
		logrus.Error(err)
		return
	}
	c.Writer.Write(body)
}
