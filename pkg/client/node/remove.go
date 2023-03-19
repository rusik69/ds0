package node

import (
	"errors"
	"fmt"
	"net/http"
)

// Add adds the node to the nameserver
func Remove(nodeName, host, port string) error {
	url := fmt.Sprintf("http://%s:%s/node/remove?name=%s", host, port, nodeName)
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return errors.New("remove failed")
	}
	return nil
}
