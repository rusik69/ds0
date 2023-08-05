package node

import (
	"errors"
	"fmt"
	"net/http"
)

// Add adds the node to the nameserver
func Add(nodeName, nodeHostname, nodePort, host, port string) error {
	url := fmt.Sprintf("http://%s:%s/api/v0/node/add?name=%s&hostname=%s&port=%s", host, port, nodeName, nodeHostname, nodePort)
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return errors.New("node add failed")
	}
	return nil
}
