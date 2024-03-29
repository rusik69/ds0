package file

import (
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"os"
	"strconv"

	"github.com/rusik69/ds0/pkg/ns/env"
)

// Upload uploads a file to the server.
func Upload(src, dst, host, port string) error {
	// get file size first
	file, err := os.Open(src)
	if err != nil {
		return err
	}
	defer file.Close()
	info, err := file.Stat()
	if err != nil {
		return err
	}
	fileSize := info.Size()
	url := "http://" + host + ":" + port + "/api/v0/file/upload?file=" + dst + "&size=" + strconv.FormatInt(fileSize, 10)
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		bodyStr, err := io.ReadAll(resp.Body)
		if err != nil {
			return err
		}
		return errors.New("upload failed: " + string(bodyStr))
	}
	var nodes []env.NodeInfo
	if err := json.NewDecoder(resp.Body).Decode(&nodes); err != nil {
		return err
	}
	if len(nodes) == 0 {
		return errors.New("no nodes available")
	}
	// upload file to nodes
	for _, node := range nodes {
		file, err := os.Open(src)
		if err != nil {
			return err
		}
		defer file.Close()
		url := "http://" + node.Host + ":" + node.Port + "/file/upload?file=" + dst
		resp, err := http.Post(url, "application/octet-stream", file)
		if err != nil {
			return err
		}
		resp.Body.Close()
		file.Close()
		if resp.StatusCode != http.StatusOK {
			bodyStr, err := io.ReadAll(resp.Body)
			if err != nil {
				return err
			}
			return errors.New("upload failed: " + url + " " + http.StatusText(resp.StatusCode) + " " + string(bodyStr))
		}

	}
	// commit file
	url = "http://" + host + ":" + port + "/api/v0/file/commit?file=" + dst
	resp, err = http.Get(url)
	if err != nil {
		return err
	}
	if resp.StatusCode != http.StatusOK {
		return errors.New("commit failed")
	}
	return nil
}
