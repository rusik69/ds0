package web

import (
	"net/http"
	"text/template"

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
	tmpl, err := template.ParseFiles("/app/html/index.html")
	if err != nil {
		c.Writer.WriteHeader(500)
		c.Writer.Write([]byte(err.Error()))
		logrus.Error(err)
		return
	}
	err = tmpl.Execute(c.Writer, data)
	if err != nil {
		c.Writer.WriteHeader(500)
		c.Writer.Write([]byte(err.Error()))
		logrus.Error(err)
		return
	}
	c.Writer.WriteHeader(http.StatusOK)
}
