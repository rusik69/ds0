package files

import (
	"github.com/gin-gonic/gin"
	"github.com/rusik69/ds0/pkg/web/metrics"
	"github.com/sirupsen/logrus"
)

// ListHandler handles the list files request.
func ListHandler(c *gin.Context) {
	metrics.Counter.Inc()
	logrus.Println("ListHandler")
}
