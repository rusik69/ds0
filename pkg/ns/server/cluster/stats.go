package cluster

import (
	"encoding/json"
	"net/http"

	"github.com/gin-gonic/gin"
	clientnode "github.com/rusik69/ds0/pkg/client/node"
	"github.com/rusik69/ds0/pkg/node/server"
	dbnode "github.com/rusik69/ds0/pkg/ns/db/node"
	"github.com/rusik69/ds0/pkg/ns/env"
	"github.com/rusik69/ds0/pkg/ns/metrics"
	"github.com/sirupsen/logrus"
)

// StatsHandler handles the stats request.
func StatsHandler(c *gin.Context) {
	metrics.Counter.Inc()
	nodesList, err := dbnode.List()
	if err != nil {
		c.Writer.WriteHeader(500)
		c.Writer.Write([]byte(err.Error()))
		logrus.Error(err)
		return
	}
	var nodes []server.NodeStats
	var totalSpace, totalFreeSpace, totalUsedSpace uint64
	for _, node := range nodesList {
		nodeStats, err := clientnode.Stats(node.Host, node.Port)
		if err != nil {
			c.Writer.WriteHeader(500)
			c.Writer.Write([]byte(err.Error()))
			logrus.Error(err)
			return
		}
		nodes = append(nodes, nodeStats)
		totalSpace += nodeStats.TotalSpace
		totalFreeSpace += nodeStats.FreeSpace
		totalUsedSpace += nodeStats.UsedSpace
	}
	stats := ClusterStats{
		Nodes:          nodes,
		NodesCount:     len(nodesList),
		TotalSpace:     totalSpace,
		TotalFreeSpace: totalFreeSpace,
		TotalUsedSpace: totalUsedSpace,
		Replicas:       env.NSEnvInstance.Replicas,
	}
	body, err := json.Marshal(stats)
	if err != nil {
		c.Writer.WriteHeader(500)
		c.Writer.Write([]byte(err.Error()))
		logrus.Error(err)
		return
	}
	c.Writer.WriteHeader(http.StatusOK)
	c.Writer.Write(body)
}
