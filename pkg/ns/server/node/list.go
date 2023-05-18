package node

import (
	"encoding/json"
	"net/http"

	"github.com/gin-gonic/gin"
	dbnode "github.com/rusik69/ds0/pkg/ns/db/node"
	"github.com/sirupsen/logrus"
)

// ListNodesHandler lists nodes.
func ListNodesHandler(c *gin.Context) {
	nodes, err := dbnode.List()
	if err != nil {
		c.Writer.WriteHeader(500)
		c.Writer.Write([]byte(err.Error()))
		logrus.Error(err)
		return
	}
	body, err := json.Marshal(nodes)
	if err != nil {
		c.Writer.WriteHeader(500)
		c.Writer.Write([]byte(err.Error()))
		logrus.Error(err)
		return
	}
	c.Writer.WriteHeader(http.StatusOK)
	c.Writer.Write(body)
}
