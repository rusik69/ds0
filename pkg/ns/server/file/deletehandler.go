package file

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	dbfile "github.com/rusik69/ds0/pkg/ns/db/file"
	"github.com/sirupsen/logrus"
)

// DeleteHandler handles the delete request.
func DeleteHandler(c *gin.Context) {
	fileName := c.Query("file")
	if fileName == "" {
		c.Writer.WriteHeader(400)
		c.Writer.Write([]byte("file name is required"))
		logrus.Error(errors.New("file name is required"))
		return
	}
	logrus.Println("DeleteHandler: " + fileName)
	fileInfo, err := dbfile.Get(fileName)
	if err != nil {
		c.Writer.WriteHeader(500)
		c.Writer.Write([]byte(err.Error()))
		logrus.Error(err)
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
		c.Writer.WriteHeader(500)
		c.Writer.Write([]byte(err.Error()))
		logrus.Error(err)
		return
	}
}
