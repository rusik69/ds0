package file

import (
	"encoding/json"
	"net/http"

	"github.com/gin-gonic/gin"
	dbfile "github.com/rusik69/ds0/pkg/ns/db/file"
	"github.com/rusik69/ds0/pkg/ns/metrics"
	"github.com/rusik69/ds0/pkg/ns/server/utils"
	"github.com/sirupsen/logrus"
)

// DownloadHandler is the download handler.
func DownloadHandler(c *gin.Context) {
	metrics.Counter.Inc()
	fileName := c.Query("file")
	if fileName == "" {
		utils.Error("file name is required", 400, c)
		return
	}
	logrus.Println("DownloadHandler: " + fileName)
	nodes, err := dbfile.Get(fileName)
	if err != nil {
		utils.Error(err.Error(), 500, c)
		return
	}
	s, err := json.Marshal(nodes)
	if err != nil {
		utils.Error(err.Error(), 500, c)
		return
	}
	c.Writer.WriteHeader(http.StatusOK)
	c.Writer.Write(s)
}
