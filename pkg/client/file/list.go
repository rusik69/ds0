package file

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"

	"github.com/rusik69/ds0/pkg/ns/db"
)

// List returns the list of files by prefix.
func List(host, port, prefix string) (map[string]db.FileInfo, error) {
	url := fmt.Sprintf("http://%s:%s/api/v0/file/list?prefix=%s", host, port, prefix)
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
	var files map[string]db.FileInfo
	if err := json.Unmarshal(bodyBytes, &files); err != nil {
		return nil, err
	}
	return files, nil
}
