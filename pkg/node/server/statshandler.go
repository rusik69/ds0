package server

import (
	"fmt"
	"syscall"

	"github.com/gin-gonic/gin"
	"github.com/rusik69/ds0/pkg/node/env"
	"github.com/sirupsen/logrus"
)

// statsHandler handles info requests.
func statsHandler(c *gin.Context) {
	var stat syscall.Statfs_t
	dir := env.NodeEnvInstance.Dir
	err := syscall.Statfs(dir, &stat)
	if err != nil {
		logrus.Println(err)
		c.String(500, fmt.Sprintf("Error: %v", err))
		return
	}
	// Total blocks * size per block = total space in bytes
	totalSpace := stat.Blocks * uint64(stat.Bsize)
	// Free blocks * size per block = free space in bytes
	freeSpace := stat.Bavail * uint64(stat.Bsize)
	// Used space = total space - free space
	usedSpace := totalSpace - freeSpace
	info := NodeStats{TotalSpace: totalSpace, FreeSpace: freeSpace, UsedSpace: usedSpace}
	c.JSON(200, info)
}
