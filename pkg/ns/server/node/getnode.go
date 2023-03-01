package node

import (
	"net/http"

	"github.com/gin-gonic/gin"
	dbnode "github.com/rusik69/ds0/pkg/ns/db/node"
)

// GetNode gets a node info.
func GetNode(c *gin.Context) {
	nodeName := c.Query("name")
	if nodeName == "" {
		c.Writer.WriteHeader(400)
		c.Writer.Write([]byte("Node name is required"))
		return
	}
	host, port, err := dbnode.Get(nodeName)
	if err != nil {
		c.Writer.WriteHeader(400)
		c.Writer.Write([]byte(err.Error()))
		return
	}
	c.Writer.WriteHeader(http.StatusOK)
	c.Writer.Write([]byte(host + ":" + port))
}
