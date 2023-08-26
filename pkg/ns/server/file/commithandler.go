package file

import (
	"net/http"

	"github.com/gin-gonic/gin"
	dbfile "github.com/rusik69/ds0/pkg/ns/db/file"
	"github.com/rusik69/ds0/pkg/ns/metrics"
	"github.com/rusik69/ds0/pkg/ns/server/utils"
	"github.com/sirupsen/logrus"
)

// CommitHandler is the handler for commit.
func CommitHandler(c *gin.Context) {
	metrics.Counter.Inc()
	fileName := c.Query("file")
	if fileName == "" {
		utils.Error("file name is required", 400, c)
		return
	}
	logrus.Println("CommitHandler: " + fileName)
	file, err := dbfile.GetUncommitted(fileName)
	if err != nil {
		utils.Error(err.Error(), 500, c)
		return
	}
	err = commitStats(file.Size)
	if err != nil {
		utils.Error(err.Error(), 500, c)
		return
	}
	err = dbfile.Commit(fileName)
	if err != nil {
		utils.Error(err.Error(), 500, c)
		return
	}
	c.Writer.WriteHeader(http.StatusOK)
}
