package utils

import (
	"errors"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

// Error writes the error to the log and returns it to client.
func Error(e string, errCode int, c *gin.Context) {
	err := errors.New(e)
	c.Writer.WriteHeader(errCode)
	c.Writer.Write([]byte(e))
	logrus.Error(err)
}
