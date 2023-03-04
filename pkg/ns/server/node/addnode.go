package node

import (
	"net/http"

	"github.com/gin-gonic/gin"
	dbnode "github.com/rusik69/ds0/pkg/ns/db/node"
)

// AddNode adds a node.
func AddNode(c *gin.Context) {
	nodeName := c.Query("name")
	if nodeName == "" {
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
	err := dbnode.Add(nodeName, nodeName, port)
	if err != nil {
		c.Writer.WriteHeader(500)
		c.Writer.Write([]byte(err.Error()))
		return
	}
	c.Writer.WriteHeader(http.StatusOK)
	c.Writer.Write([]byte("OK"))
}
