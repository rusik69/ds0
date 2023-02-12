package node

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/rusik69/ds0/pkg/ns/db"
)

// DeleteNode deletes node.
func DeleteNode(c *gin.Context) {
	nodeName := c.Query("name")
	if nodeName == "" {
		c.Writer.WriteHeader(400)
		c.Writer.Write([]byte("Node name is required"))
		return
	}
	err := db.DeleteNode(nodeName)
	if err != nil {
		c.Writer.WriteHeader(500)
		c.Writer.Write([]byte(err.Error()))
		return
	}
	c.Writer.WriteHeader(http.StatusOK)
	c.Writer.Write([]byte("OK"))
}
