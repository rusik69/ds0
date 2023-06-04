package web

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/rusik69/ds0/pkg/ns/db/node"
	"github.com/sirupsen/logrus"
)

// RootHandler is the root handler.
func RootHandler(c *gin.Context) {
	logrus.Println("RootHandler")
	nodes, err := node.List()
	if err != nil {
		c.Writer.WriteHeader(500)
		c.Writer.Write([]byte(err.Error()))
		logrus.Error(err)
		return
	}
	data := gin.H{
		"Tittle": "DS0",
		"Nodes":  nodes,
	}
	c.HTML(http.StatusOK, "html/index.html", data)
}
