package node

import (
	"net/http"

	"github.com/gin-gonic/gin"
	dbnode "github.com/rusik69/ds0/pkg/ns/db/node"
	"github.com/sirupsen/logrus"
)

// GetNodeHandler gets a node info.
func GetNodeHandler(c *gin.Context) {
	nodeName := c.Query("name")
	if nodeName == "" {
		c.Writer.WriteHeader(400)
		c.Writer.Write([]byte("Node name is required"))
		return
	}
	logrus.Println("GetNode: " + nodeName)
	host, port, err := dbnode.Get(nodeName)
	if err != nil {
		c.Writer.WriteHeader(400)
		c.Writer.Write([]byte(err.Error()))
		return
	}
	c.Writer.WriteHeader(http.StatusOK)
	c.Writer.Write([]byte(host + ":" + port))
}
