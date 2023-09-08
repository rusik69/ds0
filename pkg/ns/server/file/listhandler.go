package file

import (
	"log"

	"github.com/gin-gonic/gin"
	dbfile "github.com/rusik69/ds0/pkg/ns/db/file"
	"github.com/rusik69/ds0/pkg/ns/metrics"
	"github.com/rusik69/ds0/pkg/ns/server/utils"
)

// ListHandler handles the list files request.
func ListHandler(c *gin.Context) {
	metrics.Counter.Inc()
	prefix := c.Query("prefix")
	if prefix == "" {
		prefix = "/"
	}
	log.Println("ListHandler: " + prefix)
	files, err := dbfile.List(prefix)
	if err != nil {
		utils.Error(err.Error(), 500, c)
	}
	c.JSON(200, files)
}
