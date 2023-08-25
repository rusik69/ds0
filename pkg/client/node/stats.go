package node

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"

	server "github.com/rusik69/ds0/pkg/node/server"
)

// Stats gets the info of a node.
func Stats(host, port string) (server.NodeStats, error) {
	url := fmt.Sprintf("http://%s:%s/stats", host, port)
	resp, err := http.Get(url)
	if err != nil {
		return server.NodeStats{}, err
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return server.NodeStats{}, errors.New("node stats failed")
	}
	bodyBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		return server.NodeStats{}, err
	}
	var nodeStats server.NodeStats
	if err := json.Unmarshal(bodyBytes, &nodeStats); err != nil {
		return server.NodeStats{}, err
	}
	nodeStats.Host = host
	nodeStats.Port = port
	return nodeStats, nil
}
