package node

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"

	node "github.com/rusik69/ds0/pkg/node/server"
)

// Stats gets the info of a node.
func Stats(host, port string) (node.NodeStats, error) {
	url := fmt.Sprintf("http://%s:%s/stats", host, port)
	resp, err := http.Get(url)
	if err != nil {
		return node.NodeStats{}, err
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return node.NodeStats{}, errors.New("info failed")
	}
	bodyBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		return node.NodeStats{}, err
	}
	var nodeStats node.NodeStats
	if err := json.Unmarshal(bodyBytes, &nodeStats); err != nil {
		return node.NodeStats{}, err
	}
	return nodeStats, nil
}
