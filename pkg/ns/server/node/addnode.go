package node

import (
	"net/http"

	"github.com/gin-gonic/gin"
	dbnode "github.com/rusik69/ds0/pkg/ns/db/node"
	"github.com/sirupsen/logrus"
)

// AddNodeHandler adds a node.
func AddNodeHandler(c *gin.Context) {
	hostName := c.Query("hostname")
	if hostName == "" {
		c.Writer.WriteHeader(400)
		c.Writer.Write([]byte("Node name is required"))
		return
	}
	port := c.Query("port")
	if port == "" {
		c.Writer.WriteHeader(400)
		c.Writer.Write([]byte("Node port is required"))
		return
	}
	logrus.Println("AddNode: " + hostName + ":" + port)
	err := dbnode.Add(hostName, hostName, port)
	if err != nil {
		c.Writer.WriteHeader(500)
		c.Writer.Write([]byte(err.Error()))
		return
	}
	c.Writer.WriteHeader(http.StatusOK)
	c.Writer.Write([]byte("OK"))
}
