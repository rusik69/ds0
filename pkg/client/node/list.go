package node

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"

	"github.com/rusik69/ds0/pkg/ns/db"
)

// List returns the list of nodes in the database.
func List(host, port string) ([]db.HostInfo, error) {
	url := fmt.Sprintf("http://%s:%s/node/list", host, port)
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return nil, errors.New("list failed")
	}
	bodyBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	var nodes []db.HostInfo
	if err := json.Unmarshal(bodyBytes, &nodes); err != nil {
		return nil, err
	}
	return nodes, nil
}
