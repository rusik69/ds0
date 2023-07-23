package file

import (
	"encoding/json"
	"os"
	"strconv"
	"time"

	"github.com/rusik69/ds0/pkg/ns/server/utils"

	"github.com/gin-gonic/gin"
	"github.com/rusik69/ds0/pkg/ns/db"
	dbfile "github.com/rusik69/ds0/pkg/ns/db/file"
	dbnode "github.com/rusik69/ds0/pkg/ns/db/node"
	"github.com/rusik69/ds0/pkg/ns/env"
	"github.com/sirupsen/logrus"
)

// UploadHandler is the handler for uploading a file.
func UploadHandler(c *gin.Context) {
	fileName := c.Query("file")
	if fileName == "" {
		utils.Error("file name is required", 400, c)
		return
	}
	fileSize := c.Query("size")
	if fileSize == "" {
		utils.Error("file size is required", 400, c)
		return
	}
	size, error := strconv.ParseUint(fileSize, 10, 64)
	if error != nil {
		utils.Error("file size is invalid", 400, c)
		return
	}
	logrus.Println("UploadHandler: " + fileName + " size: " + fileSize)
	fileInfo, err := dbfile.Get(fileName)
	if err == os.ErrNotExist {
		newNodes, err := dbnode.List()
		if err != nil {
			utils.Error(err.Error(), 500, c)
			return
		}
		if len(newNodes) == 0 {
			utils.Error("no nodes available", 500, c)
			logrus.Error(env.NSEnvInstance.Nodes)
			return
		} else if len(newNodes) < env.NSEnvInstance.Replicas {
			utils.Error("not enough nodes available", 500, c)
			return
		}
		nodes := dbfile.ChooseNodes(newNodes)
		timestamp := uint64(time.Now().Unix())
		err = dbfile.Set(fileName, db.FileInfo{Nodes: nodes, TimeAdded: timestamp, Size: size})
		if err != nil {
			utils.Error(err.Error(), 500, c)
			return
		}
		body, err := json.Marshal(nodes)
		if err != nil {
			utils.Error(err.Error(), 500, c)
			return
		}
		err = incrementStats(fileInfo.Size)
		if err != nil {
			utils.Error(err.Error(), 500, c)
			return
		}
		c.Writer.Write(body)
		return
	} else if err != nil {
		utils.Error(err.Error(), 500, c)
		return
	}
	body, err := json.Marshal(fileInfo.Nodes)
	if err != nil {
		utils.Error(err.Error(), 500, c)
		return
	}
	err = incrementStats(fileInfo.Size)
	if err != nil {
		utils.Error(err.Error(), 500, c)
		return
	}
	c.Writer.Write(body)
}
