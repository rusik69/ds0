package cluster

import (
	"encoding/json"
	"errors"
	"io"
	"net/http"

	clusterstats "github.com/rusik69/ds0/pkg/ns/server/cluster"
)

// Stats returns cluster stats.
func Stats(host, port string) (clusterstats.ClusterStats, error) {
	url := "http://" + host + ":" + port + "/api/v0/cluster/stats"
	var clusterStats clusterstats.ClusterStats
	resp, err := http.Get(url)
	if err != nil {
		return clusterStats, err
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return clusterStats, errors.New("stats failed")
	}
	bodyBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		return clusterStats, err
	}
	err = json.Unmarshal(bodyBytes, &clusterStats)
	if err != nil {
		return clusterStats, err
	}
	return clusterStats, nil
}
