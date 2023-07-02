package file

import (
	"net/http"

	"github.com/gin-gonic/gin"
	dbfile "github.com/rusik69/ds0/pkg/ns/db/file"
	"github.com/rusik69/ds0/pkg/ns/server/utils"
	"github.com/sirupsen/logrus"
)

// DeleteHandler handles the delete request.
func DeleteHandler(c *gin.Context) {
	fileName := c.Query("file")
	if fileName == "" {
		utils.Error("file name is required", 400, c)
		return
	}
	logrus.Println("DeleteHandler: " + fileName)
	fileInfo, err := dbfile.Get(fileName)
	if err != nil {
		utils.Error(err.Error(), 500, c)
		return
	}
	for _, node := range fileInfo.Nodes {
		url := "http://" + node.Host + ":" + node.Port + "/file/delete?file=" + fileName
		logrus.Println("delete: " + url)
		resp, err := http.Get(url)
		if err != nil {
			logrus.Println(err)
			continue
		}
		if resp.StatusCode != http.StatusOK {
			logrus.Println("status code: " + resp.Status)
			continue
		}
	}
	err = dbfile.Delete(fileName)
	if err != nil {
		utils.Error(err.Error(), 500, c)
		return
	}
	err = decrementStats(fileInfo.Size)
	if err != nil {
		utils.Error(err.Error(), 500, c)
		return
	}
	c.Writer.WriteHeader(http.StatusOK)
}
